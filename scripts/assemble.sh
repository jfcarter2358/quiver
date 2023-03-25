#! /usr/bin/env bash

mkdir -p out

file_path="${1%.*}"
file="${file_path##*/}"

./dist/quiverc assemble $@

mv "${file_path}.qvc" "out/${file}.qvc"
