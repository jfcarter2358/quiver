#! /usr/bin/env bash

mkdir -p out

file_path="${1%.*}"
file="${file_path##*/}"

./dist/quiver assemble $@

mv "${file_path}.qvc" "out/${file}.qvc"
