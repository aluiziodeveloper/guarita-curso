package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aluiziodeveloper/guarita-curso/internal/adapters/sandbox"
)

var (
	errSandboxMount = errors.New(
		"-sandbox-mount must be source:target or source:target:rw")
	errSandboxMode = errors.New(
		"-sandbox is for a stdio server, after --; HTTP servers are remote")
	errSandboxHome = errors.New("-sandbox-mount needs the home " +
		"directory to check the credential folders")
)

// defenseFlags declara as opções das defesas dos Módulos 13 a 22: a
// aprovação, os inspetores, a sandbox, os limites, a auditoria, o
// digest do servidor, os padrões seguros e a escuta de operação.
func defenseFlags(flags *flag.FlagSet, c *config) {
	approvalFlags(flags, &c.approval)
	inspectFlags(flags, &c.inspect)
	sandboxFlags(flags, &c.sandbox)
	limitsFlags(flags, &c.limits)
	auditFlags(flags, &c.audit)
	artifactFlags(flags, c)
	secureFlags(flags, c)
	opsFlags(flags, &c.ops)
}

// sandboxConfig põe o servidor stdio num contêiner.
type sandboxConfig struct {
	image  string // vazio: o servidor roda no host, como até aqui
	mounts []sandbox.Mount
}

// sandboxFlags declara as opções da sandbox. Cada montagem é só de
// leitura, a menos que termine em :rw.
func sandboxFlags(flags *flag.FlagSet, c *sandboxConfig) {
	flags.StringVar(&c.image, "sandbox", "",
		"imagem em que o servidor stdio roda isolado")
	flags.Func("sandbox-mount", "diretório para o servidor: "+
		"origem:destino[:rw]", func(value string) error {
		parts := strings.Split(value, ":")
		if len(parts) < 2 || len(parts) > 3 ||
			(len(parts) == 3 && parts[2] != "rw") {
			return errSandboxMount
		}
		c.mounts = append(c.mounts, sandbox.Mount{
			Source: parts[0], Target: parts[1],
			Writable: len(parts) == 3,
		})
		return nil
	})
}

// profile monta o perfil da sandbox, com cada montagem conferida pelo
// caminho real. Sem -sandbox, não há perfil, e o servidor roda no host.
// Sem a home, uma montagem não sobe: as pastas de credenciais dela
// ficariam fora da conferência.
func (c sandboxConfig) profile() (*sandbox.Profile, error) {
	if c.image == "" {
		return nil, nil
	}
	p := sandbox.Default(c.image)
	if len(c.mounts) == 0 {
		return &p, p.Validate()
	}
	home, err := userHome()
	if err != nil {
		return nil, err
	}
	for _, m := range c.mounts {
		checked, err := sandbox.CheckMount(m, home)
		if err != nil {
			return nil, err
		}
		p.Mounts = append(p.Mounts, checked)
	}
	return &p, p.Validate()
}

// userHome é a home do usuário, por um caminho absoluto. Vazia ou
// relativa, não serve para conferir as montagens.
func userHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("%w: %w", errSandboxHome, err)
	}
	if !filepath.IsAbs(home) {
		return "", fmt.Errorf("%w: %q is not absolute", errSandboxHome,
			home)
	}
	return home, nil
}
