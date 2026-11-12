package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aluiziodeveloper/guarita-curso/internal/adapters/sandbox"
)

// fakeHome aponta HOME para um diretório temporário com .ssh, para o
// teste valer em qualquer máquina.
func fakeHome(t *testing.T) string {
	home := t.TempDir()
	t.Setenv("HOME", home)
	_ = os.Mkdir(filepath.Join(home, ".ssh"), 0o700)
	return home
}

// sandboxProfile lê as opções e monta o perfil, como na partida.
func sandboxProfile(args ...string) (*sandbox.Profile, error) {
	cfg, err := parseConfig(append([]string{"-allow", "x"}, args...),
		&strings.Builder{})
	if err != nil {
		return nil, err
	}
	return cfg.sandbox.profile()
}

// A sandbox não sobe com a montagem mal escrita, no modo HTTP nem com
// a home montada; a home é falsa, com .ssh.
func TestSandbox_recusas(t *testing.T) {
	home := fakeHome(t)
	for _, c := range []struct {
		name, args string
		want       error
	}{
		{"montagem sem destino", "-sandbox-mount /a", errUsage},
		{"modo HTTP", "-upstream http://x/mcp", errSandboxMode},
		{
			"home montada", "-sandbox-mount " + home + ":/data -- /srv",
			sandbox.ErrMount,
		},
	} {
		args := append([]string{"-sandbox", "img"},
			strings.Fields(c.args)...)
		_, err := sandboxProfile(args...)
		t.Logf("%-20s recusada=%v", c.name, errors.Is(err, c.want))
		if !errors.Is(err, c.want) {
			t.Errorf("%s: %v; quer %v", c.name, err, c.want)
		}
	}
}

func TestSandbox_perfil(t *testing.T) {
	data := t.TempDir()
	p, err := sandboxProfile("-sandbox", "img", "-sandbox-mount",
		data+":/data", "--", "/srv")
	if err != nil || len(p.Mounts) != 1 || p.Mounts[0].Writable {
		t.Errorf("perfil = %+v, %v", p, err)
	}
	if p, _ := sandboxProfile("--", "/srv"); p != nil {
		t.Error("sem -sandbox, não há perfil")
	}
}

// Sem a home, as pastas de credenciais dela não têm como ser
// conferidas: com HOME vazia ou relativa, a montagem não sobe. Sem
// montagens, não há o que conferir, e o perfil sai.
func TestSandbox_homeIndeterminada(t *testing.T) {
	data := t.TempDir()
	for _, home := range []string{"", "operador"} {
		t.Setenv("HOME", home)
		_, err := sandboxProfile("-sandbox", "img", "-sandbox-mount",
			data+":/data", "--", "/srv")
		t.Logf("HOME=%q: %v", home, err)
		if !errors.Is(err, errSandboxHome) {
			t.Errorf("HOME=%q: %v; quer errSandboxHome", home, err)
		}
		if _, err := sandboxProfile("-sandbox", "img", "--",
			"/srv"); err != nil {
			t.Errorf("HOME=%q, sem montagens: %v", home, err)
		}
	}
}
