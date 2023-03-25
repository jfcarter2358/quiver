package parser

import (
	"quiverc/assembler/dotcodes"
	"quiverc/assembler/opcodes"
	"quiverc/enums"
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
			opData := opcodes.ProcessAdd(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_SUBTRACT:
			opData := opcodes.ProcessSubtract(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_MULTIPLY:
			opData := opcodes.ProcessMultiply(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_DIVIDE:
			opData := opcodes.ProcessDivide(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_MODULO:
			opData := opcodes.ProcessModulo(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_POWER:
			opData := opcodes.ProcessPower(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_BINARY_ADD:
			opData := opcodes.ProcessBinaryAdd(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_BINARY_SUBTRACT:
			opData := opcodes.ProcessBinarySubtract(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_GREATER:
			opData := opcodes.ProcessGreater(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_GREATER_EQUAL:
			opData := opcodes.ProcessGreaterEqual(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_EQUAL:
			opData := opcodes.ProcessEqual(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_LESS_EQUAL:
			opData := opcodes.ProcessLessEqual(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_LESS:
			opData := opcodes.ProcessLess(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_INPUT_BLOCK:
			opData := opcodes.ProcessInputBlock(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_INPUT_NON_BLOCK:
			opData := opcodes.ProcessInputNonBlock(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_OUTPUT:
			opData := opcodes.ProcessOutput(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_BRANCH_POSITIVE:
			opData := opcodes.ProcessBranchPositive(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_BRANCH_NOT_POSITIVE:
			opData := opcodes.ProcessBranchNotPositive(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_BRANCH_ZERO:
			opData := opcodes.ProcessBranchZero(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_BRANCH_NOT_ZERO:
			opData := opcodes.ProcessBranchNotZero(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_BRANCH_NEGATIVE:
			opData := opcodes.ProcessBranchNegative(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_BRANCH_NOT_NEGATIVE:
			opData := opcodes.ProcessBranchNotNegative(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_GOTO:
			opData := opcodes.ProcessGoto(parts)
			blockData = append(blockData, opData...)
		case enums.OP_CODE_NAME_STOP:
			opData := opcodes.ProcessStop(parts)
			blockData = append(blockData, opData...)
		default:
			programCounter += 1
		}
	}

	return blockData, nil
}
