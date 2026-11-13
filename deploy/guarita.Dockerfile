# Imagem de produção da Guarita (Módulo 22). As imagens de base vêm
# pelo digest, como no Dockerfile do laboratório.
ARG GO_IMAGE=golang:1.27.4@sha256:5e8afd493a7124a8586e134d9b532731161d4b9becf9dee5a5956b107426f93d
ARG RUN_IMAGE=gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab

FROM ${GO_IMAGE} AS build
ARG VERSION=dev
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY policies/ ./policies/
ENV CGO_ENABLED=0
RUN go build -trimpath -buildvcs=false \
      -ldflags="-s -w -X main.version=${VERSION}" \
      -o /out/guarita ./cmd/guarita

# Sem shell, sem gerenciador de pacotes, usuário 65532: só o binário
# estático, os certificados raiz e as políticas.
FROM ${RUN_IMAGE}
COPY --from=build /out/guarita /guarita
COPY --from=build /src/policies /policies
USER 65532:65532
EXPOSE 8443 9090
ENTRYPOINT ["/guarita"]
