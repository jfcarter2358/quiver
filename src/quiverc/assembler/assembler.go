package assembler

import (
	"fmt"
	"quiverc/assembler/fileio"
	"quiverc/assembler/parser"
	"quiverc/assembler/utils"
	"strings"
)

func Assemble(filePath string) {
	// args := os.Args[1:]
	parts := strings.Split(filePath, ".")
	fileName := strings.Join(parts[:len(parts)-1], ".")
	fileExtension := parts[len(parts)-1]

	lines, err := fileio.ReadFile(fmt.Sprintf("%s.%s", fileName, fileExtension))
	if err != nil {
		panic(err)
	}

	// data, names, err := firstPass(lines)
	blockData, err := parser.FirstPass(lines)
	if err != nil {
		panic(err)
	}

	opData, err := parser.SecondPass(lines)

	if err != nil {
		panic(err)
	}

	blockDataLength := len(blockData)
	byteData := append(utils.IntToByteArray(blockDataLength, 8), blockData...)
	byteData = append(byteData, opData...)

	err = fileio.WriteFile(fmt.Sprintf("%s.qvc", fileName), byteData)
	if err != nil {
		panic(err)
	}
}
