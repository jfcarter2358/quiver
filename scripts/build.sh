#! /usr/bin/env bash

mkdir -p dist

# Build assembler
pushd src/assembler
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o quiver-assembler
popd

# Move assembler binary
mv src/assembler/quiver-assembler dist/quiver-assembler

# Build VM
pushd src/vm
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o quiver-vm
popd

# Move VM binary
mv src/vm/quiver-vm dist/quiver-vm
