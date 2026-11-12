// Command lab-sandbox-probe é a sonda do laboratório do Módulo 18.
// Ela tenta, de dentro de onde roda, o que um servidor comprometido
// tentaria: escrever no sistema de arquivos, abrir uma conexão, ler um
// arquivo do host, usar o chaveiro do kernel e virar root. Cada tentativa é
// local e inofensiva: o arquivo escrito é um marcador, a conexão é com
// um ouvinte do próprio teste, e nada é executado.
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// marker é tudo o que a sonda escreve.
var marker = []byte("[LAB]")

// checks são as tentativas da sonda, na ordem em que aparecem.
func checks(connect, secret string) []struct {
	name string
	try  func() error
} {
	return []struct {
		name string
		try  func() error
	}{
		{"escrever na raiz", func() error {
			return os.WriteFile("/sonda-lab.txt", marker, 0o600)
		}},
		{"conectar ao host", func() error {
			c, err := net.DialTimeout("tcp", connect, 2*time.Second)
			if err == nil {
				_ = c.Close()
			}
			return err
		}},
		{"ler arquivo do host", func() error {
			//nolint:gosec // G304: ler o arquivo do host é a tentativa
			_, err := os.ReadFile(secret)
			return err
		}},
		{"chaveiro do kernel", keyring},
		{"virar root", setuidRoot},
	}
}


func main() {
	connect := flag.String("connect", "", "endereço do ouvinte do teste")
	secret := flag.String("secret", "", "arquivo do host a ler")
	flag.Parse()
	fmt.Printf("%-20s %d\n", "usuário", os.Getuid())
	for _, c := range checks(*connect, *secret) {
		result := "conseguiu"
		if err := c.try(); err != nil {
			result = "negado: " + err.Error()
		}
		fmt.Printf("%-20s %s\n", c.name, result)
	}
}
