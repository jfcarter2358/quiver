package parser

import (
	"assembler/dotcodes"
	"assembler/enums"
	"assembler/opcodes"
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

func FirstPass(lines []string) ([]byte, error) {
	var blockData []byte
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

		switch strings.ToUpper(code) {
		case enums.DOT_CODE_NAME_LITERAL:
			dotData, err := dotcodes.ProcessLiteral(parts)
			if err != nil {
				return nil, err
			}
			blockData = append(blockData, dotData...)
		case enums.DOT_CODE_NAME_LABEL:
			dotData := dotcodes.ProcessLabel(parts, programCounter)
			blockData = append(blockData, dotData...)
		default:
			programCounter += 1
		}
	}
	return blockData, nil
}

func SecondPass(lines []string) ([]byte, error) {
	var blockData []byte
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

		switch strings.ToUpper(code) {
		case enums.OP_CODE_NAME_ADD:
			dotData := opcodes.ProcessAdd(parts)
			blockData = append(blockData, dotData...)
		case enums.OP_CODE_NAME_OUTPUT:
			dotData := opcodes.ProcessOutput(parts)
			blockData = append(blockData, dotData...)
		case enums.OP_CODE_NAME_STOP:
			dotData := opcodes.ProcessStop(parts)
			blockData = append(blockData, dotData...)
		default:
			programCounter += 1
		}
	}

	return blockData, nil
}
