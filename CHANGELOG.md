# Changelog

Este arquivo segue o formato de [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/), e a Guarita segue o [Versionamento Semântico](https://semver.org/lang/pt-BR/). Até a 1.0.0, cada módulo do curso é uma *tag* (`aula-02-11` a `aula-23-11`), e a API pública ainda não é estável.

## [Não publicado] — 1.0.0

A 1.0.0 sai como *release* assinada no dia da edição 1.0 do curso, depois da revisão técnica externa. O conteúdo é o da *tag* `aula-23-11`.

### Adicionado

- **Proxy MCP** (Caps. 1 a 7): modo stdio com repasse concorrente e *parser* estrito (um valor por mensagem, sem chave repetida, sem lote, tamanho limitado, *fuzzing*); modo HTTP com *streaming*, guarda de Host contra *DNS rebinding* e prazos.
- **Identidade** (Módulo 9): token do agente verificado (Ed25519, emissor, audiência por rota), credencial própria da Guarita para cada servidor, nada de repasse do token do usuário; revisão `2026-07-28` do MCP.
- **Catálogo** (Caps. 9 e 10): *pin* pela impressão digital SHA-256 da definição, em SQLite com migrações; varredor estrutural de caracteres invisíveis e de controle; um namespace por servidor e recusa da silhueta repetida.
- **Política** (Módulo 12): Rego com o OPA embutido, negação por padrão, concessões por agente e servidor, `opa test` na CI.
- **Aprovação humana** (Módulo 13): fila com prazo, limite de pendentes e recusa da repetição; interface do operador separada, pela linha de comando e por uma página.
- **Inspeção** (Caps. 13 a 16): DLP de CPF, chaves e entropia, lista de destinos e JSON Schema; revisão e marcação dos resultados; classificador local com Ollama, limiar medido e plano B; linha de inspetores por rota com camada exigida.
- **Contenção** (Caps. 17 a 19): sandbox do servidor stdio num contêiner sem rede e sem privilégios; cota por balde de fichas, orçamento de custo e disjuntor; trilha de auditoria encadeada por *hashes*, com HMAC dos argumentos, verificador e trace W3C.
- **Produção** (Caps. 20 e 21): servidor fixado pelo digest do executável; *release* reproduzível com SBOM e assinatura; identidade e TLS 1.3 por padrão; escuta de operação com métricas, saúde e perfis só no loopback; SLO de 5 ms no p99 na CI; imagem distroless e manifestos de Compose e Kubernetes.
- **Red team** (Módulo 23): `attacks/redteam`, a cadeia combinada contra a Guarita inteira, nos dois modos; [relatório de segurança](docs/security-report.md) e [*roadmap*](docs/roadmap.md).

### Corrigido

- O modo stdio aceitava e ignorava `-dlp`, `-quota`, `-require-stages` e as opções da volta; agora aplica a linha de ida e recusa o que não aplica (Módulo 23).
- O erro JSON-RPC da resposta de `tools/call` passava sem revisão (Módulo 23).

### Segurança

- go1.27.2, que corrige 12 vulnerabilidades alcançáveis do go1.27.1 apontadas pelo `govulncheck` (Módulo 21).
- OPA 1.21.2, `x/sync` 0.24.0, `x/text` 0.43.0, `x/crypto` 0.57.0 e `x/net` 0.60.0 (Módulo 23).
- Nenhuma ameaça do [modelo](docs/threats/model.md) está mitigada por completo; o risco residual está no relatório de segurança.
