# Imagem de desenvolvimento: o toolchain do projeto, sem depender do Go
# instalado na máquina de quem desenvolve.
FROM golang:1.27.4

# Um usuário comum, com o mesmo número do usuário do host, para os arquivos
# criados em /workspace não ficarem com o dono root.
ARG UID=1000
RUN useradd --create-home --uid "$UID" dev \
    && mkdir -p /home/dev/.cache/go-build \
    && chown -R dev:dev /home/dev /go \
    && git config --system --add safe.directory /workspace

USER dev
WORKDIR /workspace
