#! /usr/bin/env bash

# Ensure dist directory exists
mkdir -p dist

# Build quiver
pushd src/quiver
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o quiver
popd

# Build quiverc
pushd src/quiverc
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o quiverc
popd

# Move binaries to dist directory
mv src/quiver/quiver dist/quiver
mv src/quiverc/quiverc dist/quiverc