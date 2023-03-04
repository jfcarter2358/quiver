package opcodes

import (
	"quiver/assembler/utils"
	"quiver/enums"
)

func process1(opCode byte, parts []string) []byte {
	var returnBytes []byte

	a := parts[1]
	aLength := utils.IntToByteArray(len(a), 1)

	returnBytes = []byte{opCode}
	returnBytes = append(returnBytes, aLength...)
	returnBytes = append(returnBytes, []byte(a)...)

	return returnBytes
}

func process2(opCode byte, parts []string) []byte {
	var returnBytes []byte

	a := parts[1]
	aLength := utils.IntToByteArray(len(a), 1)

	b := parts[2]
	bLength := utils.IntToByteArray(len(b), 1)

	returnBytes = []byte{opCode}
	returnBytes = append(returnBytes, aLength...)
	returnBytes = append(returnBytes, []byte(a)...)
	returnBytes = append(returnBytes, bLength...)
	returnBytes = append(returnBytes, []byte(b)...)

	return returnBytes
}

func process3(opCode byte, parts []string) []byte {
	var returnBytes []byte

	a := parts[1]
	aLength := utils.IntToByteArray(len(a), 1)

	b := parts[2]
	bLength := utils.IntToByteArray(len(b), 1)

	c := parts[3]
	cLength := utils.IntToByteArray(len(c), 1)

	returnBytes = []byte{opCode}
	returnBytes = append(returnBytes, aLength...)
	returnBytes = append(returnBytes, []byte(a)...)
	returnBytes = append(returnBytes, bLength...)
	returnBytes = append(returnBytes, []byte(b)...)
	returnBytes = append(returnBytes, cLength...)
	returnBytes = append(returnBytes, []byte(c)...)

	return returnBytes
}

func ProcessAdd(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_ADD, parts)

	return returnBytes
}

func ProcessSubtract(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_SUBTRACT, parts)

	return returnBytes
}

func ProcessMultiply(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_MULTIPLY, parts)

	return returnBytes
}

func ProcessDivide(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_DIVIDE, parts)

	return returnBytes
}

func ProcessModulo(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_MODULO, parts)

	return returnBytes
}

func ProcessPower(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_POWER, parts)

	return returnBytes
}

func ProcessBinaryAdd(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_BINARY_ADD, parts)

	return returnBytes
}

func ProcessBinarySubtract(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_BINARY_SUBTRACT, parts)

	return returnBytes
}

func ProcessGreater(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_GREATER, parts)

	return returnBytes
}

func ProcessGreaterEqual(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_GREATER_EQUAL, parts)

	return returnBytes
}

func ProcessEqual(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_EQUAL, parts)

	return returnBytes
}

func ProcessLessEqual(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_LESS_EQUAL, parts)

	return returnBytes
}

func ProcessLess(parts []string) []byte {
	returnBytes := process3(enums.OP_CODE_BYTE_LESS, parts)

	return returnBytes
}

func ProcessInputBlock(parts []string) []byte {
	returnBytes := process1(enums.OP_CODE_BYTE_INPUT_BLOCK, parts)

	return returnBytes
}

func ProcessInputNonBlock(parts []string) []byte {
	returnBytes := process2(enums.OP_CODE_BYTE_INPUT_NON_BLOCK, parts)

	return returnBytes
}

func ProcessOutput(parts []string) []byte {
	returnBytes := process1(enums.OP_CODE_BYTE_OUTPUT, parts)

	return returnBytes
}

func ProcessBranchPositive(parts []string) []byte {
	returnBytes := process2(enums.OP_CODE_BYTE_BRANCH_POSITIVE, parts)

	return returnBytes
}

func ProcessBranchNotPositive(parts []string) []byte {
	returnBytes := process2(enums.OP_CODE_BYTE_BRANCH_NOT_POSITIVE, parts)

	return returnBytes
}

func ProcessBranchZero(parts []string) []byte {
	returnBytes := process2(enums.OP_CODE_BYTE_BRANCH_ZERO, parts)

	return returnBytes
}

func ProcessBranchNotZero(parts []string) []byte {
	returnBytes := process2(enums.OP_CODE_BYTE_BRANCH_NOT_ZERO, parts)

	return returnBytes
}

func ProcessBranchNegative(parts []string) []byte {
	returnBytes := process2(enums.OP_CODE_BYTE_BRANCH_NEGATIVE, parts)

	return returnBytes
}

func ProcessBranchNotNegative(parts []string) []byte {
	returnBytes := process2(enums.OP_CODE_BYTE_BRANCH_NOT_NEGATIVE, parts)

	return returnBytes
}

func ProcessGoto(parts []string) []byte {
	returnBytes := process1(enums.OP_CODE_BYTE_GOTO, parts)

	return returnBytes
}

func ProcessStop(parts []string) []byte {
	returnBytes := process1(enums.OP_CODE_BYTE_STOP, parts)

	return returnBytes
}
