#!/usr/bin/env bash

set -x

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "skycoin binary dir:" "$DIR"
pushd "$DIR" >/dev/null

COMMIT=$(git rev-parse HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
GOLDFLAGS="-X main.Commit=${COMMIT} -X main.Branch=${BRANCH}"

go run -ldflags "${GOLDFLAGS}" cmd/haicoin/haicoin.go \
    -gui-dir="${DIR}/src/gui/static/" \
    -launch-browser=false \
    -enable-wallet-api=true \
    -enable-gui=false \
    -rpc-interface=true \
    -log-level=debug \
    -enable-seed-api=true \
    -web-interface=true \
    -enable-unversioned-api=true \
    -disable-csrf=true \
    $@

popd >/dev/null
