package dlp

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/aluiziodeveloper/guarita-curso/internal/domain"
)

// Só lab.test e os subdomínios de api.lab.test. O truque do usuário
// antes da arroba, a maiúscula e o ponto final não enganam o url.Parse.
func TestInspect_destinos(t *testing.T) {
	inspector := New(DefaultActions(), 4.3,
		WithAllowedHosts("lab.test", ".api.lab.test"))
	for _, url := range []string{
		"https://lab.test/reports",
		"https://v1.api.lab.test/upload",
		"https://collector.evil.test/c",
		"https://lab.test@collector.evil.test/c",
		"HTTPS://LAB.TEST./reports",
		"//collector.evil.test/c",
		"envie para https://collector.evil.test/c hoje",
		"lab.test/reports e só texto",
	} {
		got, _ := inspector.Inspect(context.Background(),
			domain.PolicyInput{Call: domain.Call{
				ToolName:  "http_request",
				Arguments: json.RawMessage(`{"url":"` + url + `"}`),
			}})
		t.Logf("%-40s %-6s %v", url, got.Decision, got.Findings)
		hostile := strings.Contains(url, "evil")
		if hostile != (got.Decision == domain.DecisionDeny) {
			t.Errorf("%s: %s", url, got.Decision)
		}
	}
}

// A URL escrita em largura total é conferida nas duas formas, a
// escrita e a normalizada, e basta uma apontar para fora da lista.
func TestInspect_destinosEmLarguraTotal(t *testing.T) {
	inspector := New(DefaultActions(), 4.3,
		WithAllowedHosts("lab.test", ".api.lab.test"))
	for _, c := range []struct {
		url  string
		deny bool
	}{
		{fullwidth("https") + "://collector.evil.test/c", true},
		{"https://" + fullwidth("collector.evil.test") + "/c", true},
		{"https://lab.test/" + fullwidth("reports"), false},
	} {
		arguments, _ := json.Marshal(map[string]string{"url": c.url})
		got, _ := inspector.Inspect(context.Background(),
			domain.PolicyInput{Call: domain.Call{
				ToolName: "http_request", Arguments: arguments,
			}})
		t.Logf("%-44s %-6s %v", c.url, got.Decision, got.Findings)
		if c.deny != (got.Decision == domain.DecisionDeny) {
			t.Errorf("%s: %s", c.url, got.Decision)
		}
	}
}
