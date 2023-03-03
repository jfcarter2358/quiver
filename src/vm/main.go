package main

import (
	"os"
	"vm/fileio"
	"vm/memstore"
	"vm/parser"
	"vm/runner"
)

func main() {
	args := os.Args[1:]

	byteCode, err := fileio.ReadQVC(args[0])

	if err != nil {
		panic(err)
	}

	memstore.Init()

	byteCode, err = parser.ParseBlockData(byteCode)
	if err != nil {
		panic(err)
	}

	// memstore.PrintState()

	err = runner.Run(byteCode)
	if err != nil {
		panic(err)
	}

	// memstore.PrintState()
}
