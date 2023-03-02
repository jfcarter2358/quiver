package main

import (
	"assembler/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"encoding/binary"
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

func readFile(path string) ([]string, error) {
	lines := []string{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeFile(path string, data []byte) error {
	outFile, err := os.Create(path)
	if err != nil {
		return err
	}
	for _, datum := range data {
		err := binary.Write(outFile, binary.BigEndian, datum)
		if err != nil {
			return err
		}
	}
	outFile.Close()
	return nil
}

func main() {
	args := os.Args[1:]
	parts := strings.Split(args[0], ".")
	fileName := strings.Join(parts[:len(parts)-1], ".")
	fileExtension := parts[len(parts)-1]

	lines, err := readFile(fmt.Sprintf("%s.%s", fileName, fileExtension))
	if err != nil {
		panic(err)
	}

	data, names, err := firstPass(lines)

	if err != nil {
		panic(err)
	}

	dataLength := len(data)
	byteData := append(utils.IntToByteArray(dataLength, 8), data...)

	fmt.Println(byteData)
	fmt.Println(data)
	fmt.Println(names)

	err = writeFile(fmt.Sprintf("%s.qvc", fileName), byteData)
	if err != nil {
		panic(err)
	}
	// writeFile(args)
}
