package main

import (
	"fmt"
	"os"
	"quiver/assembler"
	"quiver/utils"
	"quiver/vm"
)

func main() {
	usage := `
usage: quiver <command> <path> [-h|--help]

commands:
	assemble    Assemble a .qasm file into a bytecode
	compile     Compile quiver source down to bytecode
	run         Run an assembled quiver application

arguments:
	--help      Show this help message and exit
`

	args := os.Args[1:]

	if utils.Contains(args, "-h") || utils.Contains(args, "--help") {
		fmt.Println(usage)
		os.Exit(0)
	}

	if len(args) < 2 {
		fmt.Println("ERROR: Insufficient number of arguments")
		fmt.Println(usage)
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("ERROR: Too many arguments")
		fmt.Println(usage)
		os.Exit(1)
	}

	command := args[0]
	path := args[1]

	switch command {
	case "assemble":
		assembler.Assemble(path)
	case "compile":
		fmt.Println("Not implemented")
	case "run":
		vm.Run(path)
	default:
		fmt.Printf("Invalid command: %s\n", command)
		fmt.Println(usage)
		os.Exit(1)
	}

}
