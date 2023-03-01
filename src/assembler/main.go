package main

import (
	"assembler/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getParts(line string) []string {
	var parts []string
	var part string
	quoted := false

	for _, char := range line {
		if char == ' ' && !quoted {
			parts = append(parts, part)
			part = ""
			continue
		}
		if char == '"' {
			quoted = !quoted
		}
		part += string(char)
	}
	parts = append(parts, part)

	return parts
}

func firstPass(lines []string) ([]byte, []string, error) {
	var blockData []byte
	var names []string
	programCounter := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := getParts(line)
		code := parts[0]

		switch code {
		case ".LITERAL":
			opCode := []byte{byte(0)}
			name := parts[1]
			data := []byte(parts[2])
			length := utils.IntToByteArray(len(data), 8)

			names = append(names, name)
			blockData = append(blockData, opCode...)
			blockData = append(blockData, length...)
			blockData = append(blockData, data...)
		case ".RESERVE":
			opCode := []byte{byte(1)}
			name := parts[1]
			length, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, nil, err
			}
			byteLength := utils.IntToByteArray(length, 8)
			data := []byte{}

			names = append(names, name)
			for i := 0; i < length; i++ {
				data = append(data, byte(0))
			}
			blockData = append(blockData, opCode...)
			blockData = append(blockData, byteLength...)
			blockData = append(blockData, data...)
		case ".LABEL":
			opCode := []byte{byte(2)}
			name := parts[1]
			line := utils.IntToByteArray(programCounter+1, 8)

			names = append(names, name)
			blockData = append(blockData, opCode...)
			blockData = append(blockData, line...)
		default:
			programCounter += 1
		}
	}
	return blockData, names, nil
}

func main() {
	args := os.Args[1:]

	lines := []string{}

	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	file.Close()

	data, names, err := firstPass(lines)

	if err != nil {
		panic(err)
	}

	fmt.Println(names)
	fmt.Println(data)
	fmt.Println(len(data))

}
