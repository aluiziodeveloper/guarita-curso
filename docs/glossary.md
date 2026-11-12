# Glossário

A linguagem da Guarita: os termos que o código, os testes, as ADRs e o modelo de
ameaças usam com o mesmo sentido. A coluna "Módulo" diz em que módulo do curso
o termo foi apresentado. O glossário cresce a cada módulo.

## Termos

| Termo | Significado | Módulo |
|---|---|---|
| *Confused deputy* | Componente com privilégio que age em nome de quem não o tem | 04 |
| *Slowloris* | Ataque que prende conexões enviando o pedido devagar | 08 |
| `errgroup` | Grupo de goroutines com erro e cancelamento compartilhados | 06 |
| `fmt.Stringer` | Interface com `String() string`, usada pela formatação | 04 |
| `internal` | Pasta cujos pacotes só podem ser importados dentro do mesmo prefixo de caminho | 01 |
| `iota` | Contador de linhas de um bloco `const`, começando em 0 | 04 |
| `select` | Espera o primeiro de vários canais | 06 |
| `sync.Mutex` | Trava de exclusão mútua: uma goroutine por vez entre `Lock` e `Unlock` | 07 |
| Adaptador | Peça de tecnologia que encaixa numa porta | 05 |
| ADR | Registro de decisão de arquitetura: uma decisão, com o contexto e as consequências | 01 |
| Agente | Aplicação que coordena o modelo e as conexões com servidores MCP | 01 |
| Arquitetura hexagonal (*hexagonal architecture*) | Núcleo isolado da infraestrutura, conversando por portas | 05 |
| Arquivo *golden* | Arquivo com a saída esperada de um teste, regravado de propósito | 07 |
| Ativo | O que a Guarita precisa proteger | 04 |
| Build em múltiplos estágios | Dockerfile que compila num estágio e copia só o resultado para outro | 01 |
| Cabeçalho de salto a salto | Cabeçalho que vale só entre dois nós, como `Connection` | 08 |
| Camada anticorrupção | Camada que traduz entre o protocolo externo e o domínio | 07 |
| Canal | Cano tipado entre goroutines, que também sincroniza | 06 |
| Caso de uso | Operação da aplicação, como `InterceptCall` | 05 |
| Chamada de ferramentas | Ciclo em que o modelo pede uma ferramenta e recebe o resultado como mensagem | 01 |
| Chave duplicada | Mesmo nome duas vezes num objeto JSON, com leitura imprevisível | 07 |
| CI | Integração contínua: compilação e testes automáticos a cada push, no GitHub Actions | 01 |
| Conjunto de métodos | Métodos que pertencem a `T` ou a `*T`; decide a satisfação de interfaces | 03 |
| Contrabando de mensagem | Segunda mensagem colada depois da primeira na mesma linha | 07 |
| Contrapressão | Espera que volta até quem envia quando o destino não acompanha | 06 |
| Corpus de sementes | Entradas iniciais do fuzzing, de `f.Add` e de `testdata/fuzz` | 07 |
| Corrida de dados | Acesso concorrente sem ordem a um valor, com ao menos uma escrita | 06 |
| CSRF | Pedido disparado por outra página, com a credencial que o navegador manda sozinho | 08 |
| DNS rebinding | Troca do endereço por trás de um nome para alcançar um serviço local na mesma origem | 08 |
| Dono de goroutine | Código que a inicia, a faz parar e espera o fim dela | 06 |
| Dublê de teste (*test double*) | Substituto de uma dependência durante o teste | 05 |
| Embrulhar um erro | Criar um erro novo que guarda o anterior, com `%w` | 03 |
| Endpoint MCP | Endereço único do transporte HTTP com streaming, como `/mcp`, que recebe os pedidos | 08 |
| Enumeração completa | Teste que percorre todos os valores de um tipo enumerado | 04 |
| Envenenamento de ferramenta (*tool poisoning*) | Instrução hostil escondida nos metadados de uma ferramenta, que o modelo lê | 01 |
| Erro sentinela | Valor de erro declarado no pacote e comparado por identidade | 03 |
| Espião (*spy*) | Dublê de teste que registra as chamadas que recebe, para o teste conferi-las | 05 |
| Exemplo testável | Função `Example` cuja saída o `go test` confere | 03 |
| Falha fechada (*fail closed*) | Negar em vez de permitir diante do imprevisto ou quando a defesa não consegue decidir | 05 |
| Fechamento | Função que captura variáveis do escopo em que foi criada | 01 |
| Ferramenta | Capacidade anunciada por um servidor MCP, com nome, descrição e esquema | 01 |
| Fluxo de eventos (SSE) | Resposta `text/event-stream`, com eventos enviados um a um | 08 |
| Fronteira de confiança (*trust boundary*) | Ponto em que dados de uma parte passam a influenciar decisões de outra | 01 |
| Fuzzing | Teste que gera entradas guiado pela cobertura de código | 07 |
| Goroutine | Função executando de forma independente, gerida pelo runtime de Go | 06 |
| Inversão de dependência | O núcleo declara a interface; a infraestrutura a satisfaz | 05 |
| JSON-RPC | Formato de chamada remota em JSON usado por todas as mensagens MCP | 01 |
| Lista de permissão | Regra que só aceita o que foi declarado; o resto fica de fora até ser ligado | 05 |
| Lote | Array de mensagens JSON-RPC, removido do MCP | 07 |
| Marcador sintético | Texto fixo do laboratório, sem correspondência com um segredo real | 01 |
| Middleware | `http.Handler` que embrulha outro, como um decorador | 08 |
| Modelo de ameaças | Documento versionado com ativos, zonas, fronteiras e ameaças, mantido junto do código | 04 |
| Modelo de domínio anêmico | Tipos com nomes de negócio e pouco ou nenhum comportamento próprio | 03 |
| Módulo | Unidade Go que declara caminho, versão mínima e dependências | 01 |
| Notificação | Mensagem com `method` e sem `id`, que não recebe resposta | 07 |
| Opção funcional | Função que ajusta um valor em construção | 05 |
| Origem | Esquema, nome e porta de uma página | 08 |
| Pacote de teste externo | Pacote `x_test` que só enxerga a API exportada de `x` | 03 |
| Parser estrito | Parser que recusa a mensagem inteira diante de qualquer ambiguidade | 07 |
| Parâmetro variádico | Último parâmetro `...T`, recebido como fatia | 04 |
| Pedido | Mensagem JSON-RPC com `method` e `id`, que espera resposta | 07 |
| Porta | Conversa com propósito, expressa como interface do caso de uso | 05 |
| Portas e adaptadores (*ports and adapters*) | Outro nome da arquitetura hexagonal: o núcleo fala com o mundo por portas, e cada tecnologia é um adaptador | 05 |
| Proxy reverso | Servidor que recebe pedidos em nome de outro e os repassa | 08 |
| Raiz de composição | O ponto, em `cmd/guarita`, que escolhe e monta os adaptadores | 05 |
| Receptor | Parâmetro que associa um método a um tipo, por valor ou por ponteiro | 03 |
| Regra de dependência | Dependências de código apontam para o núcleo | 05 |
| Restrição de compilação | Linha `//go:build` que decide se um arquivo entra no pacote | 06 |
| STRIDE | Seis categorias de ameaça, cada uma ligada a uma propriedade de segurança | 04 |
| Teste de consistência | Teste que compara um documento de projeto com os dados do código | 04 |
| Teste de contrato | Teste do comportamento visível de uma interface, pedido a pedido | 08 |
| Teste orientado a tabelas (*table-driven tests*) | Casos nomeados numa fatia, executados pelo mesmo laço com subtestes | 03 |
| Tipo definido | Tipo novo, com identidade própria, declarado a partir de um tipo subjacente | 03 |
| Toolchain | Compilador e ferramentas acessados pelo comando `go` | 01 |
| Validação na fronteira | Conferência dos dados antes de entregá-los ao próximo componente | 03 |
| Valor zero | Valor que Go atribui a uma variável sem inicialização explícita | 03 |
| Vazamento de goroutine (*goroutine leak*) | Goroutine que nunca termina e acumula recursos | 06 |
| Zona de confiança | Parte do sistema que opera sob um mesmo nível de confiança | 04 |
| Árvore de ataque (*attack tree*) | Objetivo do atacante na raiz, caminhos em ramos E e OU | 04 |

## Siglas

| Sigla | Por extenso | Módulo |
|---|---|---|
| ADR | Registro de decisão de arquitetura (*architecture decision record*) | 01 |
| API | Interface de programação de aplicações | 01 |
| CI | Integração contínua (*continuous integration*) | 01 |
| CPF | Cadastro de Pessoas Físicas | 04 |
| CSRF | Falsificação de pedido entre sites (*cross-site request forgery*) | 08 |
| JWT | Token assinado em JSON (*JSON Web Token*) | 08 |
| MCP | Protocolo de contexto de modelo (*Model Context Protocol*) | 01 |
| OPA | O motor de políticas *Open Policy Agent* | 05 |
| OWASP | *Open Worldwide Application Security Project* | 01 |
| SDK | Kit de desenvolvimento de software | 01 |
| SSE | Fluxo de eventos enviados pelo servidor (*Server-Sent Events*) | 08 |
| SSRF | Falsificação de pedido pelo servidor (*server-side request forgery*) | 08 |
| STRIDE | As seis categorias de ameaça: falsificação de identidade, adulteração, repúdio, divulgação de informação, negação de serviço e elevação de privilégio | 04 |
