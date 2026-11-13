#!/bin/sh
# Monta a release da Guarita como a de uma versão, sem publicar nada e
# sem rede: o contêiner roda com --network none, os módulos vêm do cache
# local, só para leitura, e as ferramentas, de $TOOLS. Sem Go no host, o
# cache é o volume go-mod do contêiner de desenvolvimento.
#   TOOLS=dir-com-goreleaser-syft-e-cosign \
#   COSIGN_PRIVATE_KEY="$(cat cosign.key)" COSIGN_PASSWORD=… \
#   deploy/release/snapshot.sh [opções do goreleaser release]
set -eu
if command -v go >/dev/null 2>&1; then modcache=$(go env GOMODCACHE)
else modcache=guarita-dev_go-mod; fi
image=golang:1.27.4@sha256:5e8afd493a7124a8586e134d9b532731161d4b9becf9dee5a5956b107426f93d
docker run --rm --network none -u "$(id -u):$(id -g)" -e HOME=/tmp \
  -v "$TOOLS:/tools:ro" -v "$modcache:/go/pkg/mod:ro" \
  -v "$PWD:/src" -w /src \
  -e GOFLAGS=-mod=readonly -e GOPROXY=off -e GOTOOLCHAIN=local \
  -e SYFT_CHECK_FOR_APP_UPDATE=false \
  -e COSIGN_PRIVATE_KEY -e COSIGN_PASSWORD \
  "$image" sh -c 'git config --global --add safe.directory /src &&
    PATH=/tools:$PATH goreleaser release --snapshot --clean "$@"' \
  sh "$@"
