package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/aluiziodeveloper/guarita-curso/internal/adapters/artifact"
)

// guarita digest cat dá o digest do cat desta máquina. Com ele, a
// Guarita inicia o servidor e repassa; com outro, ela recusa antes de
// iniciar, e nada chega ao servidor.
func TestRun_digestDoServidor(t *testing.T) {
	var out bytes.Buffer
	if err := run(t.Context(), []string{"digest", "cat"}, nil,
		nopWriteCloser{&out}, io.Discard); err != nil {
		t.Fatal(err)
	}
	approved, _, _ := strings.Cut(out.String(), " ")
	t.Logf("guarita digest cat: %.23s…", approved)
	other := "sha256:" + strings.Repeat("0", 64)
	for _, d := range []string{approved, other} {
		var stdout, stderr bytes.Buffer
		ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
		err := run(ctx, []string{"-server-digest", d, "--", "cat"},
			io.NopCloser(strings.NewReader(ping(1))),
			nopWriteCloser{&stdout}, &stderr)
		cancel()
		t.Logf("%.15s…: eco %q; erro: %.60v", d,
			strings.TrimSpace(stdout.String()), err)
		if (d == approved) != (err == nil) ||
			(err != nil && !errors.Is(err, artifact.ErrMismatch)) {
			t.Errorf("%s: %v", d, err)
		}
	}
}

// Com -sandbox, o digest é o da referência da imagem: uma etiqueta é
// recusada antes do docker run.
func TestRun_digestDaImagem(t *testing.T) {
	d := "sha256:" + strings.Repeat("ab", 32)
	err := run(t.Context(), []string{
		"-server-digest", d,
		"-sandbox", "guarita-lab:m21", "--", "/lab-files",
	},
		io.NopCloser(strings.NewReader("")),
		nopWriteCloser{&bytes.Buffer{}}, io.Discard)
	t.Logf("etiqueta: %v", err)
	if !errors.Is(err, artifact.ErrUnpinned) {
		t.Errorf("a etiqueta passou: %v", err)
	}
}
