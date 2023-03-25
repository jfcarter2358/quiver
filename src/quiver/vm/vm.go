package vm

import (
	"quiver/vm/fileio"
	"quiver/vm/memstore"
	"quiver/vm/parser"
	"quiver/vm/runner"
)

func Run(path string) {

	byteCode, err := fileio.ReadQVC(path)

	if err != nil {
		panic(err)
	}

	memstore.Init()

	byteCode, err = parser.ParseBlockData(byteCode)
	if err != nil {
		panic(err)
	}

	// memstore.PrintState()

	instructions := parser.ParseInstructions(byteCode)

	// fmt.Printf("%v\n", instructions)

	err = runner.Run(instructions)
	if err != nil {
		panic(err)
	}

	// memstore.PrintState()
}
