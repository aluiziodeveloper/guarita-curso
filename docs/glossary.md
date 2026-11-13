# Glossário

A linguagem da Guarita: os termos que o código, os testes, as ADRs e o modelo de
ameaças usam com o mesmo sentido. A coluna "Módulo" diz em que módulo do curso
o termo foi apresentado. O glossário cresce a cada módulo.

## Termos

| Termo | Significado | Módulo |
|---|---|---|
| *Confused deputy* | Componente com privilégio que age em nome de quem não o tem | 04 |
| *Slowloris* | Ataque que prende conexões enviando o pedido devagar | 08 |
| *Token passthrough* | Repassar ao destino um token que não foi emitido para ele | 09 |
| `errgroup` | Grupo de goroutines com erro e cancelamento compartilhados | 06 |
| `fmt.Stringer` | Interface com `String() string`, usada pela formatação | 04 |
| `internal` | Pasta cujos pacotes só podem ser importados dentro do mesmo prefixo de caminho | 01 |
| `iota` | Contador de linhas de um bloco `const`, começando em 0 | 04 |
| `opa test` | O executor de testes escritos em Rego | 12 |
| `select` | Espera o primeiro de vários canais | 06 |
| `slog.LogValuer` | Interface com que um valor diz como quer ser registrado | 09 |
| `sync.Mutex` | Trava de exclusão mútua: uma goroutine por vez entre `Lock` e `Unlock` | 07 |
| `user_version` | Inteiro no cabeçalho do SQLite, livre para a aplicação | 10 |
| Adaptador | Peça de tecnologia que encaixa numa porta | 05 |
| ADR | Registro de decisão de arquitetura: uma decisão, com o contexto e as consequências | 01 |
| Agente | Aplicação que coordena o modelo e as conexões com servidores MCP | 01 |
| Agência excessiva | O agente tem ferramentas, permissões ou autonomia além do que a tarefa exige | 12 |
| Aprovação forjada | Uma aprovação registrada sem a decisão de um operador | 13 |
| Arquitetura hexagonal (*hexagonal architecture*) | Núcleo isolado da infraestrutura, conversando por portas | 05 |
| Arquivo *golden* | Arquivo com a saída esperada de um teste, regravado de propósito | 07 |
| Ativo | O que a Guarita precisa proteger | 04 |
| Audiência (`aud`) | Recurso para o qual o token foi emitido | 09 |
| Bolha | O grupo de goroutines de `synctest.Test`, com relógio falso | 13 |
| Build em múltiplos estágios | Dockerfile que compila num estágio e copia só o resultado para outro | 01 |
| Cabeçalho de salto a salto | Cabeçalho que vale só entre dois nós, como `Connection` | 08 |
| Cabeçalho espelhado | Cópia, num cabeçalho HTTP, de um campo do corpo | 09 |
| Camada anticorrupção | Camada que traduz entre o protocolo externo e o domínio | 07 |
| Canal | Cano tipado entre goroutines, que também sincroniza | 06 |
| Capacidade (de acesso) | Um valor cuja posse já dá acesso, como um identificador que não se adivinha | 13 |
| Caractere de formatação | Ponto de código `Cf`: não desenha nada, ou muda o que está em volta | 11 |
| Caso de uso | Operação da aplicação, como `InterceptCall` | 05 |
| Categoria geral | Classificação Unicode de cada ponto de código, como `Cf` ou `Mn` | 11 |
| Chamada de ferramentas | Ciclo em que o modelo pede uma ferramenta e recebe o resultado como mensagem | 01 |
| Chave duplicada | Mesmo nome duas vezes num objeto JSON, com leitura imprevisível | 07 |
| CI | Integração contínua: compilação e testes automáticos a cada push, no GitHub Actions | 01 |
| Comparação em tempo constante | Comparação cujo tempo não depende de onde os valores diferem | 13 |
| Concessão | A decisão dada a um agente, num servidor, para uma ferramenta | 12 |
| Confiança no primeiro uso | Aceitar e fixar o que se vê da primeira vez, e comparar depois | 10 |
| Confusão de algoritmo | Aceitar o algoritmo que o atacante escreveu no cabeçalho | 09 |
| Conjunto de métodos | Métodos que pertencem a `T` ou a `*T`; decide a satisfação de interfaces | 03 |
| Consulta parametrizada | Texto da consulta e valores enviados separados | 10 |
| Contrabando de mensagem | Segunda mensagem colada depois da primeira na mesma linha | 07 |
| Contrapressão | Espera que volta até quem envia quando o destino não acompanha | 06 |
| Corpus de ataques | Conjunto de entradas em arquivos, cada uma com o resultado esperado | 11 |
| Corpus de sementes | Entradas iniciais do fuzzing, de `f.Add` e de `testdata/fuzz` | 07 |
| Corrida de dados | Acesso concorrente sem ordem a um valor, com ao menos uma escrita | 06 |
| CSRF | Pedido disparado por outra página, com a credencial que o navegador manda sozinho | 08 |
| DLP | Prevenção contra vazamento de dados: regras que acham dados sensíveis no que sai | 14 |
| DNS rebinding | Troca do endereço por trás de um nome para alcançar um serviço local na mesma origem | 08 |
| Dono de goroutine | Código que a inicia, a faz parar e espera o fim dela | 06 |
| Dublê de teste (*test double*) | Substituto de uma dependência durante o teste | 05 |
| Dígito verificador | Dígito calculado a partir dos outros, que denuncia um número inventado | 14 |
| Ed25519 | Algoritmo de assinatura por curva elíptica; `EdDSA` no JWT | 09 |
| Embrulhar um erro | Criar um erro novo que guarda o anterior, com `%w` | 03 |
| Endpoint MCP | Endereço único do transporte HTTP com streaming, como `/mcp`, que recebe os pedidos | 08 |
| Entrada da política | O que a política recebe: agente, servidor e chamada | 12 |
| Entropia de Shannon | Imprevisibilidade de um texto, em bits por caractere | 14 |
| Enumeração completa | Teste que percorre todos os valores de um tipo enumerado | 04 |
| Envenenamento de ferramenta (*tool poisoning*) | Instrução hostil escondida nos metadados de uma ferramenta, que o modelo lê | 01 |
| Erro sentinela | Valor de erro declarado no pacote e comparado por identidade | 03 |
| Espião (*spy*) | Dublê de teste que registra as chamadas que recebe, para o teste conferi-las | 05 |
| Estado terminal | Estado sem saída: o pedido já foi decidido | 13 |
| Exemplo testável | Função `Example` cuja saída o `go test` confere | 03 |
| Exfiltração | Levar para fora, pelos argumentos de uma chamada, um dado que não deveria sair | 14 |
| Fadiga de aprovação (*approval fatigue*) | O operador aprova por cansaço, depois de pedidos repetidos | 13 |
| Falha fechada (*fail closed*) | Negar em vez de permitir diante do imprevisto ou quando a defesa não consegue decidir | 05 |
| Falso positivo | Valor legítimo acusado por um detector | 14 |
| Fechamento | Função que captura variáveis do escopo em que foi criada | 01 |
| Ferramenta | Capacidade anunciada por um servidor MCP, com nome, descrição e esquema | 01 |
| Fluxo de eventos (SSE) | Resposta `text/event-stream`, com eventos enviados um a um | 08 |
| Forma canônica | Uma única maneira de escrever um valor, para que o hash seja sempre o mesmo | 10 |
| Fronteira de confiança (*trust boundary*) | Ponto em que dados de uma parte passam a influenciar decisões de outra | 01 |
| Fuzzing | Teste que gera entradas guiado pela cobertura de código | 07 |
| Goroutine | Função executando de forma independente, gerida pelo runtime de Go | 06 |
| Impressão digital | SHA-256 da forma canônica da definição | 10 |
| Indefinido | O estado de uma regra cujo corpo não vale | 12 |
| Inspetor | Etapa que examina o conteúdo de uma chamada depois da política | 14 |
| Inversão de dependência | O núcleo declara a interface; a infraestrutura a satisfaz | 05 |
| JSON-RPC | Formato de chamada remota em JSON usado por todas as mensagens MCP | 01 |
| JWT | Token com afirmações assinadas, em três partes | 09 |
| Lista de destinos | Os hosts para onde uma URL nos argumentos pode apontar | 14 |
| Lista de permissão | Regra que só aceita o que foi declarado; o resto fica de fora até ser ligado | 05 |
| Lote | Array de mensagens JSON-RPC, removido do MCP | 07 |
| Marcador sintético | Texto fixo do laboratório, sem correspondência com um segredo real | 01 |
| Menor privilégio (*least privilege*) | Cada agente opera com o mínimo necessário para a tarefa | 12 |
| Middleware | `http.Handler` que embrulha outro, como um decorador | 08 |
| Migração | Passo numerado de mudança no esquema do banco | 10 |
| Modelo de ameaças | Documento versionado com ativos, zonas, fronteiras e ameaças, mantido junto do código | 04 |
| Modelo de domínio anêmico | Tipos com nomes de negócio e pouco ou nenhum comportamento próprio | 03 |
| Máquina de estados | Conjunto fechado de estados e das transições permitidas entre eles | 13 |
| Módulo | Unidade Go que declara caminho, versão mínima e dependências | 01 |
| Namespace (de servidor) | Nome que o operador dá a um servidor, e o caminho dele na Guarita | 11 |
| Negação por padrão | Tudo é negado, salvo o que for concedido | 12 |
| Normalização | Escolher uma escrita só para textos equivalentes (NFC, NFKC) | 11 |
| Notificação | Mensagem com `method` e sem `id`, que não recebe resposta | 07 |
| Opção funcional | Função que ajusta um valor em construção | 05 |
| Origem | Esquema, nome e porta de uma página | 08 |
| Pacote de teste externo | Pacote `x_test` que só enxerga a API exportada de `x` | 03 |
| Parser estrito | Parser que recusa a mensagem inteira diante de qualquer ambiguidade | 07 |
| Parâmetro variádico | Último parâmetro `...T`, recebido como fatia | 04 |
| PDP | Ponto que decide: o OPA com as regras de `policies/` | 12 |
| Pedido | Mensagem JSON-RPC com `method` e `id`, que espera resposta | 07 |
| PEM | Texto em Base64 entre marcadores, com a chave em PKIX ou PKCS #8 | 09 |
| PEP | Ponto que intercepta a chamada e aplica a decisão: o `InterceptCall` | 12 |
| Pin (*pinning*) | Definição fixada de uma ferramenta num servidor | 10 |
| Plano do agente | O endpoint MCP e os tokens dos agentes | 13 |
| Plano do operador | A interface de aprovação e as credenciais dos operadores | 13 |
| Políticas como código (*policy as code*) | Regras de autorização versionadas, revisadas e testadas como o resto do código | 12 |
| Porta | Conversa com propósito, expressa como interface do caso de uso | 05 |
| Portas e adaptadores (*ports and adapters*) | Outro nome da arquitetura hexagonal: o núcleo fala com o mundo por portas, e cada tecnologia é um adaptador | 05 |
| Proxy reverso | Servidor que recebe pedidos em nome de outro e os repassa | 08 |
| Quórum | Número de operadores diferentes exigido para aprovar | 13 |
| Raiz de composição | O ponto, em `cmd/guarita`, que escolhe e monta os adaptadores | 05 |
| RE2 | Sintaxe de expressões regulares com tempo de execução linear | 14 |
| Receptor | Parâmetro que associa um método a um tipo, por valor ou por ponteiro | 03 |
| Rego | A linguagem declarativa de políticas do OPA | 12 |
| Regra completa | Regra que define um único valor para um nome | 12 |
| Regra de dependência | Dependências de código apontam para o núcleo | 05 |
| Repositório | Porta que esconde onde e como os dados ficam | 10 |
| Restrição de compilação | Linha `//go:build` que decide se um arquivo entra no pacote | 06 |
| Runa | Valor de um ponto de código Unicode em Go, um `int32` | 11 |
| Semente fixa | Valor inicial do gerador que torna os casos reprodutíveis | 14 |
| Servidor de autorização | Quem autentica o usuário e emite tokens | 09 |
| Servidor de recursos | Quem recebe tokens e decide o que eles permitem; a Guarita, para o agente | 09 |
| Silhueta | O nome em minúsculas ASCII, com hífen e ponto como sublinhado | 11 |
| Sombreamento de ferramenta (*tool shadowing*) | Um servidor anuncia uma ferramenta com o nome que outro já usa | 11 |
| STRIDE | Seis categorias de ameaça, cada uma ligada a uma propriedade de segurança | 04 |
| Suíte de ataques | Testes que rodam cada ataque sem e com a defesa | 10 |
| Teste de consistência | Teste que compara um documento de projeto com os dados do código | 04 |
| Teste de contrato | Teste do comportamento visível de uma interface, pedido a pedido | 08 |
| Teste de propriedade | Teste que afirma algo sobre uma família de casos e o confere em casos gerados | 14 |
| Teste orientado a tabelas (*table-driven tests*) | Casos nomeados numa fatia, executados pelo mesmo laço com subtestes | 03 |
| Tipo definido | Tipo novo, com identidade própria, declarado a partir de um tipo subjacente | 03 |
| Toolchain | Compilador e ferramentas acessados pelo comando `go` | 01 |
| Troca silenciosa (*rug pull*) | O servidor muda a definição de uma ferramenta depois de aprovada | 10 |
| Validar em vez de normalizar | Aceitar ou recusar os dados como vieram, sem consertá-los | 11 |
| Validação na fronteira | Conferência dos dados antes de entregá-los ao próximo componente | 03 |
| Valor de contexto | Dado de um pedido guardado no `context.Context`, com chave própria | 09 |
| Valor zero | Valor que Go atribui a uma variável sem inicialização explícita | 03 |
| Varredor estrutural | Procura sinais objetivos de disfarce, não intenção | 11 |
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
| DLP | Prevenção contra vazamento de dados (*data loss prevention*) | 14 |
| HMAC | Resumo com chave, que só quem tem a chave calcula (*hash-based message authentication code*) | 09 |
| JWT | Token assinado em JSON (*JSON Web Token*) | 08 |
| LGPD | Lei Geral de Proteção de Dados Pessoais | 14 |
| MCP | Protocolo de contexto de modelo (*Model Context Protocol*) | 01 |
| OPA | O motor de políticas *Open Policy Agent* | 05 |
| OWASP | *Open Worldwide Application Security Project* | 01 |
| PDP | Ponto de decisão da política (*policy decision point*) | 12 |
| PEP | Ponto de aplicação da política (*policy enforcement point*) | 12 |
| SDK | Kit de desenvolvimento de software | 01 |
| SSE | Fluxo de eventos enviados pelo servidor (*Server-Sent Events*) | 08 |
| SSRF | Falsificação de pedido pelo servidor (*server-side request forgery*) | 08 |
| STRIDE | As seis categorias de ameaça: falsificação de identidade, adulteração, repúdio, divulgação de informação, negação de serviço e elevação de privilégio | 04 |
