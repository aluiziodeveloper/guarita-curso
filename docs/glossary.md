# Glossário

A linguagem da Guarita: os termos que o código, os testes, as ADRs e o modelo de
ameaças usam com o mesmo sentido. A coluna "Módulo" diz em que módulo do curso
o termo foi apresentado. O glossário cresce a cada módulo.

## Termos

| Termo | Significado | Módulo |
|---|---|---|
| `internal` | Pasta cujos pacotes só podem ser importados dentro do mesmo prefixo de caminho | 01 |
| ADR | Registro de decisão de arquitetura: uma decisão, com o contexto e as consequências | 01 |
| Agente | Aplicação que coordena o modelo e as conexões com servidores MCP | 01 |
| Build em múltiplos estágios | Dockerfile que compila num estágio e copia só o resultado para outro | 01 |
| Chamada de ferramentas | Ciclo em que o modelo pede uma ferramenta e recebe o resultado como mensagem | 01 |
| CI | Integração contínua: compilação e testes automáticos a cada push, no GitHub Actions | 01 |
| Conjunto de métodos | Métodos que pertencem a `T` ou a `*T`; decide a satisfação de interfaces | 03 |
| Embrulhar um erro | Criar um erro novo que guarda o anterior, com `%w` | 03 |
| Envenenamento de ferramenta (*tool poisoning*) | Instrução hostil escondida nos metadados de uma ferramenta, que o modelo lê | 01 |
| Erro sentinela | Valor de erro declarado no pacote e comparado por identidade | 03 |
| Exemplo testável | Função `Example` cuja saída o `go test` confere | 03 |
| Fechamento | Função que captura variáveis do escopo em que foi criada | 01 |
| Ferramenta | Capacidade anunciada por um servidor MCP, com nome, descrição e esquema | 01 |
| Fronteira de confiança (*trust boundary*) | Ponto em que dados de uma parte passam a influenciar decisões de outra | 01 |
| JSON-RPC | Formato de chamada remota em JSON usado por todas as mensagens MCP | 01 |
| Marcador sintético | Texto fixo do laboratório, sem correspondência com um segredo real | 01 |
| Modelo de domínio anêmico | Tipos com nomes de negócio e pouco ou nenhum comportamento próprio | 03 |
| Módulo | Unidade Go que declara caminho, versão mínima e dependências | 01 |
| Pacote de teste externo | Pacote `x_test` que só enxerga a API exportada de `x` | 03 |
| Receptor | Parâmetro que associa um método a um tipo, por valor ou por ponteiro | 03 |
| Teste orientado a tabelas (*table-driven tests*) | Casos nomeados numa fatia, executados pelo mesmo laço com subtestes | 03 |
| Tipo definido | Tipo novo, com identidade própria, declarado a partir de um tipo subjacente | 03 |
| Toolchain | Compilador e ferramentas acessados pelo comando `go` | 01 |
| Validação na fronteira | Conferência dos dados antes de entregá-los ao próximo componente | 03 |
| Valor zero | Valor que Go atribui a uma variável sem inicialização explícita | 03 |

## Siglas

| Sigla | Por extenso | Módulo |
|---|---|---|
| ADR | Registro de decisão de arquitetura (*architecture decision record*) | 01 |
| API | Interface de programação de aplicações | 01 |
| CI | Integração contínua (*continuous integration*) | 01 |
| MCP | Protocolo de contexto de modelo (*Model Context Protocol*) | 01 |
| OWASP | *Open Worldwide Application Security Project* | 01 |
| SDK | Kit de desenvolvimento de software | 01 |
