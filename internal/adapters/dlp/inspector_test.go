package dlp

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/aluiziodeveloper/guarita-curso/internal/domain"
)

// withChecks completa nove dígitos com os verificadores: nenhum CPF
// pronto fica gravado no repositório.
func withChecks(base string) string {
	digits := []byte(base)
	for _, size := range []int{9, 10} {
		sum := 0
		for i := range size {
			sum += int(digits[i]-'0') * (size + 1 - i)
		}
		digits = append(digits, byte('0'+sum*10%11%10))
	}
	return string(digits)
}

// Formatos de chave montados em tempo de execução: o texto do teste
// não casa com os padrões que varredores de segredo procuram.
var (
	awsKey  = "AKIA" + strings.Repeat("LAB0", 4)
	ghToken = "ghp" + "_" + strings.Repeat("Lab0", 9)
	pemHead = "-----BEGIN " + "RSA PRIVATE KEY-----"
	random  = "q8Zf3+Lw9/Xk2Ta7Bm4Rc1Ny6Hp0Jv5Gd8Es3Ui2Oo=" // sem sentido
)

func inspect(t *testing.T, arguments string) domain.Inspection {
	t.Helper()
	inspection, err := New(DefaultActions(), 4.3).Inspect(
		context.Background(), domain.PolicyInput{Call: domain.Call{
			ToolName: "http_request", Arguments: json.RawMessage(arguments),
		}})
	if err != nil {
		t.Fatal(err)
	}
	return inspection
}

// detectorCases são os argumentos do teste e a decisão de cada um.
func detectorCases(cpf string) []struct {
	name, arguments string
	want            domain.Decision
} {
	formatted := cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
	encoded := base64.StdEncoding.EncodeToString([]byte(cpf))
	approve, deny := domain.DecisionRequireApproval, domain.DecisionDeny
	allow := domain.DecisionAllow
	return []struct {
		name, arguments string
		want            domain.Decision
	}{
		{"CPF formatado", `{"body":"cliente ` + formatted + `"}`, approve},
		{"CPF sem pontuação", `{"body":"` + cpf + `"}`, approve},
		{"CPF em base64", `{"body":"` + encoded + `"}`, approve},
		{"verificador errado", `{"order":"123.456.789-00"}`, allow},
		{"onze iguais", `{"order":"11111111111"}`, allow},
		{"chave AWS em note", `{"note":"` + awsKey + `"}`, deny},
		{"token GitHub", `{"headers":{"auth":"` + ghToken + `"}}`, deny},
		{"chave privada", `{"body":"` + pemHead + `"}`, deny},
		{"base64 aleatório", `{"blob":"` + random + `"}`, approve},
		{"texto comum", `{"body":"relatorio_2026_q3_final"}`, allow},
	}
}

// Cada caso, um argumento e o que o inspetor acha nele. O relatório
// mostra o campo, o sinal e o detalhe, e nunca o valor.
func TestInspect_detectores(t *testing.T) {
	cpf := withChecks("123456789")
	for _, c := range detectorCases(cpf) {
		got := inspect(t, c.arguments)
		t.Logf("%-18s %-16s %v", c.name, got.Decision, got.Findings)
		if got.Decision != c.want {
			t.Errorf("%s: %s; want %s", c.name, got.Decision, c.want)
		}
		for _, f := range got.Findings {
			if strings.Contains(f.Detail, cpf) || f.Detail == awsKey {
				t.Errorf("%s: o detalhe vazou o dado", c.name)
			}
		}
	}
}

// fullwidth troca cada caractere ASCII visível pela forma de largura
// total (U+FF01 a U+FF5E): para quem lê, e para o servidor que
// normaliza, é o mesmo texto; para \d e para os prefixos, não.
func fullwidth(text string) string {
	return strings.Map(func(r rune) rune {
		if '!' <= r && r <= '~' {
			return r + 0xFEE0
		}
		return r
	}, text)
}

// O dado escrito com dígitos e letras de largura total é achado como
// o escrito em ASCII: o inspetor normaliza com NFKC e examina as duas
// formas. A forma original continua valendo: um dígito de largura
// total colado no CPF não o esconde.
func TestInspect_larguraTotal(t *testing.T) {
	cpf := withChecks("123456789")
	formatted := cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
	approve, deny := domain.DecisionRequireApproval, domain.DecisionDeny
	for _, c := range []struct {
		name, text string
		want       domain.Decision
	}{
		{"CPF de largura total", fullwidth(cpf), approve},
		{"CPF formatado", fullwidth(formatted), approve},
		{"só os dígitos", "cliente " + fullwidth(cpf[:6]) + cpf[6:], approve},
		{"dígito colado", fullwidth("1") + cpf, approve},
		{"chave AWS", fullwidth(awsKey), deny},
		{
			"verificador errado", fullwidth("12345678900"),
			domain.DecisionAllow,
		},
	} {
		arguments, _ := json.Marshal(map[string]string{"body": c.text})
		got := inspect(t, string(arguments))
		t.Logf("%-22s %-16s %v", c.name, got.Decision, got.Findings)
		if got.Decision != c.want {
			t.Errorf("%s: %s; want %s", c.name, got.Decision, c.want)
		}
		if len(got.Findings) > 1 {
			t.Errorf("%s: achado repetido: %v", c.name, got.Findings)
		}
	}
}
