#! /usr/bin/env bash

mkdir -p dist

pushd src/assembler
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o quiver-assembler
popd

mv src/assembler/quiver-assembler dist/quiver-assembler
