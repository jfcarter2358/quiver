package main

import (
	"fmt"
	"os"
	"quiver/utils"
	"quiver/vm"
)

func main() {
	usage := `
usage: quiver <path> [-h|--help]

arguments:
	--help      Show this help message and exit
`

	args := os.Args[1:]

	if utils.Contains(args, "-h") || utils.Contains(args, "--help") {
		fmt.Println(usage)
		os.Exit(0)
	}

	if len(args) < 1 {
		fmt.Println("ERROR: Insufficient number of arguments")
		fmt.Println(usage)
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("ERROR: Too many arguments")
		fmt.Println(usage)
		os.Exit(1)
	}

	path := args[0]

	vm.Run(path)
}
