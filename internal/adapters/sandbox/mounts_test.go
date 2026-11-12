package sandbox

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// layout monta uma home falsa, com .ssh e um projeto, os dados fora
// dela e um symlink de nome inocente para o .ssh.
func layout(t *testing.T) (home, data, inside, link string) {
	t.Helper()
	dir := t.TempDir()
	home = filepath.Join(dir, "home", "operador")
	data = filepath.Join(dir, "projeto", "dados")
	inside = filepath.Join(home, "projeto")
	for _, d := range []string{
		filepath.Join(home, ".ssh"), data,
		inside,
	} {
		if err := os.MkdirAll(d, 0o750); err != nil {
			t.Fatal(err)
		}
	}
	link = filepath.Join(dir, "projeto", "relatorios")
	if err := os.Symlink(filepath.Join(home, ".ssh"), link); err != nil {
		t.Fatal(err)
	}
	return home, data, inside, link
}

// Os dados, fora ou dentro da home, passam; a home, o pai dela, o .ssh
// pelo symlink de nome inocente, o socket do Docker e um caminho
// relativo, não.
func TestCheckMount(t *testing.T) {
	home, data, inside, link := layout(t)
	for _, c := range []struct {
		name, source string
		ok           bool
	}{
		{"dados", data, true},
		{"dentro da home", inside, true},
		{"home", home, false},
		{"pai da home", filepath.Dir(home), false},
		{"symlink", link, false},
		{"docker.sock", "/var/run", false},
		{"relativo", "projeto/dados", false},
	} {
		got, err := CheckMount(Mount{Source: c.source, Target: "/data"},
			home)
		t.Logf("%-14s %v", c.name, err)
		// A origem devolvida é a real: no macOS, o diretório
		// temporário fica em /var, um symlink para /private/var.
		want, _ := filepath.EvalSymlinks(c.source)
		if (err == nil) != c.ok || (err != nil && !errors.Is(err,
			ErrMount)) || (c.ok && got.Source != want) {
			t.Errorf("%s: %v %v", c.name, got, err)
		}
	}
}

// aliased refaz o layout atrás de um symlink, como /var no macOS, que
// aponta para /private/var: a home e as listas chegam pelo atalho, e
// a origem resolvida chega pelo caminho real.
func aliased(t *testing.T) (home, realRoot string) {
	t.Helper()
	dir, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	realRoot = filepath.Join(dir, "private")
	alias := filepath.Join(dir, "atalho")
	for _, d := range []string{
		filepath.Join(realRoot, "home", "operador", ".ssh"),
		filepath.Join(realRoot, "home", "operador", "projeto"),
		filepath.Join(realRoot, "run", "x"),
		filepath.Join(realRoot, "dados"),
	} {
		if err := os.MkdirAll(d, 0o750); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Symlink(realRoot, alias); err != nil {
		t.Fatal(err)
	}
	return filepath.Join(alias, "home", "operador"), realRoot
}

// Com a home e um diretório do sistema escritos por um symlink, a
// comparação precisa resolver os dois lados: a home, o pai dela, o
// .ssh e o diretório do sistema são recusados pelo caminho real e
// pelo atalho, e os dados continuam passando.
func TestCheckMount_listasPorSymlink(t *testing.T) {
	home, realRoot := aliased(t)
	alias := filepath.Dir(filepath.Dir(home))
	old := system
	system = []string{filepath.Join(alias, "run")}
	t.Cleanup(func() { system = old })
	link := filepath.Join(realRoot, "dados", "relatorios")
	ssh := filepath.Join(realRoot, "home", "operador", ".ssh")
	if err := os.Symlink(ssh, link); err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		name, source string
		ok           bool
	}{
		{"dados", filepath.Join(realRoot, "dados"), true},
		{"dentro da home", filepath.Join(home, "projeto"), true},
		{"home real", filepath.Join(realRoot, "home", "operador"), false},
		{"home pelo atalho", home, false},
		{"pai da home", filepath.Join(realRoot, "home"), false},
		{"pai pelo atalho", filepath.Dir(home), false},
		{".ssh real", ssh, false},
		{".ssh pelo atalho", filepath.Join(home, ".ssh"), false},
		{"symlink para .ssh", link, false},
		{"sistema real", filepath.Join(realRoot, "run"), false},
		{"sistema pelo atalho", filepath.Join(alias, "run"), false},
		{"dentro do sistema", filepath.Join(realRoot, "run", "x"), false},
	} {
		_, err := CheckMount(Mount{Source: c.source, Target: "/data"},
			home)
		t.Logf("%-20s %v", c.name, err)
		if (err == nil) != c.ok ||
			(err != nil && !errors.Is(err, ErrMount)) {
			t.Errorf("%s: %v", c.name, err)
		}
	}
}

// O diretório de credenciais que ainda não existe é resolvido até o
// último pai que existe: a home pelo atalho continua valendo.
func TestResolve(t *testing.T) {
	home, realRoot := aliased(t)
	for _, c := range []struct{ path, want string }{
		{home, filepath.Join(realRoot, "home", "operador")},
		{
			filepath.Join(home, ".aws", "credentials"),
			filepath.Join(realRoot, "home", "operador", ".aws",
				"credentials"),
		},
		{"/nao-existe/x", "/nao-existe/x"},
	} {
		if got := resolve(c.path); got != c.want {
			t.Errorf("resolve(%q) = %q; quer %q", c.path, got, c.want)
		}
	}
}

func TestWithin(t *testing.T) {
	for _, c := range []struct {
		path, dir string
		want      bool
	}{
		{"/etc", "/etc", true},
		{"/etc/ssh", "/etc", true},
		{"/etcetera", "/etc", false},
		{"/srv", "/", true},
		{"/", "/etc", false},
	} {
		if got := within(c.path, c.dir); got != c.want {
			t.Errorf("within(%q, %q) = %v", c.path, c.dir, got)
		}
	}
}
