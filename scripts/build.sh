#! /usr/bin/env bash

# Ensure dist directory exists
mkdir -p dist

# Build Quiver
pushd src
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o quiver
popd

# Move binary to dist directory
mv src/quiver dist/quiver