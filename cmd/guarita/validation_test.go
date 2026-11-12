package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// -validate-arguments precisa do catálogo, que guarda o inputSchema
// fixado. Sem -catalog, a partida falha com o nome da opção que falta,
// em vez de subir sem a validação que o operador pediu; no stdio, que
// não tem catálogo, a opção é recusada como as outras do modo HTTP.
func TestRun_validacaoExigeCatalogo(t *testing.T) {
	http := []string{
		"-insecure-no-auth", "-listen", "127.0.0.1:0", "-allow", "x",
		"-upstream", "http://127.0.0.1:1/mcp", "-validate-arguments",
	}
	// Com prazo: se a opção for ignorada, a Guarita sobe e fica à
	// escuta, e o teste falha em vez de esperar para sempre.
	ctx, cancel := context.WithTimeout(t.Context(), 2*time.Second)
	t.Cleanup(cancel)
	err := run(ctx, http, nil, nil, &strings.Builder{})
	t.Logf("HTTP sem -catalog: %v", err)
	if !errors.Is(err, errValidationCatalog) {
		t.Errorf("HTTP sem -catalog: %v; quer errValidationCatalog", err)
	}
	for _, option := range []string{
		"-validate-arguments", "-inspect-results", "-mark-untrusted",
	} {
		stdio := []string{"-allow", "x", option, "--", "cat"}
		err := run(t.Context(), stdio, nil, nil, &strings.Builder{})
		t.Logf("stdio %s: %v", option, err)
		if !errors.Is(err, errStdioFlag) ||
			!strings.Contains(err.Error(), option) {
			t.Errorf("stdio %s: %v; quer errStdioFlag", option, err)
		}
	}
}

// pinnedServer anuncia read_file com um inputSchema estrito e conta as
// chamadas que chegam a ele.
func pinnedServer(t *testing.T) (string, *atomic.Int32) {
	t.Helper()
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			w.Header().Set("Content-Type", "application/json")
			if strings.Contains(string(body), `"tools/list"`) {
				_, _ = io.WriteString(w, `{"jsonrpc":"2.0","id":1,`+
					`"result":{"tools":[{"name":"read_file",`+
					`"description":"[LAB] lê um arquivo",`+
					`"inputSchema":{"type":"object",`+
					`"properties":{"path":{"type":"string"}},`+
					`"required":["path"],`+
					`"additionalProperties":false}}]}}`)
				return
			}
			calls.Add(1)
			_, _ = io.WriteString(w,
				`{"jsonrpc":"2.0","id":1,"result":{}}`)
		}))
	t.Cleanup(server.Close)
	return server.URL, &calls
}

// De ponta a ponta, com -catalog, -validate-arguments e -dlp: o
// esquema vem antes do DLP. A chamada dentro do contrato chega ao
// servidor; a que foge dele é negada pelo esquema, mesmo levando uma
// chave que o DLP também negaria; a que cabe no contrato e leva a
// chave é negada pelo DLP.
func TestRun_validacaoDosArgumentos(t *testing.T) {
	upstream, calls := pinnedServer(t)
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	endpoint, _ := startHTTP(t, ctx, "-listen", "127.0.0.1:0",
		"-allow", "read_file", "-upstream", upstream+"/mcp",
		"-catalog", filepath.Join(t.TempDir(), "catalog.db"),
		"-validate-arguments", "-dlp")
	listBody(t, endpoint, // a lista fixa a definição no catálogo
		`{"jsonrpc":"2.0","id":1,"method":"tools/list"}`)
	key := "AKIA" + strings.Repeat("LAB0", 4) // formato de chave da AWS
	for _, c := range []struct {
		name, arguments, want, not string
	}{
		{
			"dentro do contrato", `{"path":"relatorio.txt"}`, `"result"`,
			`"error"`,
		},
		{
			"fora do contrato, com chave",
			`{"path":"a.txt","note":"` + key + `"}`, "schema em",
			"secret em",
		},
		{
			"no contrato, com chave", `{"path":"` + key + `"}`,
			"secret em", "schema em",
		},
	} {
		reply := callWith(t, endpoint, "read_file", c.arguments)
		t.Logf("%-28s %s", c.name, reply)
		if !strings.Contains(reply, c.want) ||
			strings.Contains(reply, c.not) {
			t.Errorf("%s: %s; quer %q sem %q", c.name, reply, c.want,
				c.not)
		}
	}
	if got := calls.Load(); got != 1 {
		t.Errorf("o servidor recebeu %d chamadas; quer 1", got)
	}
}
