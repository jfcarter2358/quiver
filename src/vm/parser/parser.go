package parser

import (
	"vm/enums"
	"vm/memstore"
	"vm/utils"
)

type Instruction struct {
	OpCode    byte
	Args      [][]byte
	DataTypes []byte
	ArgLength int
}

func advanceByteCode(byteCode []byte, byteCounter, advance int) ([]byte, int) {
	byteCode = byteCode[advance:]
	byteCounter = byteCounter + advance

	return byteCode, byteCounter
}

func advanceInstruction(instructions []byte, advance int) []byte {
	instructions = instructions[advance:]

	return instructions
}

func ParseBlockData(byteCode []byte) ([]byte, error) {
	dataLength := utils.CoerceByteInt(byteCode[:8])
	byteCode = byteCode[8:]

	byteCounter := 0
	for byteCounter < dataLength {
		dotCode := byteCode[0]
		byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, 1)

		dataType := enums.DATATYPE_BYTE_NULL

		switch dotCode {
		case enums.DOT_CODE_BYTE_LITERAL:
			dataType = byteCode[0]
			byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, 1)

			sourceLength := utils.CoerceByteInt(byteCode[:1])
			byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, 1)

			source := utils.CoerceString(byteCode[:sourceLength])
			byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, sourceLength)

			dataLength := utils.CoerceByteInt(byteCode[:8])
			byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, 8)

			switch dataType {
			case enums.DATATYPE_BYTE_BOOL:
				data, err := utils.CoerceBool(byteCode[:dataLength])
				if err != nil {
					return nil, err
				}
				memstore.BoolData[source] = data
				byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, dataLength)
			case enums.DATATYPE_BYTE_INT:
				data, err := utils.CoerceInt(byteCode[:dataLength])
				if err != nil {
					return nil, err
				}
				memstore.IntData[source] = data
				byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, dataLength)
			case enums.DATATYPE_BYTE_FLOAT:
				data, err := utils.CoerceFloat(byteCode[:dataLength])
				if err != nil {
					return nil, err
				}
				memstore.FloatData[source] = data
				byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, dataLength)
			case enums.DATATYPE_BYTE_STRING:
				data := utils.CoerceString(byteCode[:dataLength])
				memstore.StringData[source] = data
				byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, dataLength)
			}

		case enums.DOT_CODE_BYTE_LABEL:
			labelNameLength := utils.CoerceByteInt(byteCode[:1])
			byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, 1)

			labelName := utils.CoerceString(byteCode[:labelNameLength])
			byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, labelNameLength)

			programLine := utils.CoerceByteInt(byteCode[:8])

			memstore.LabelData[labelName] = programLine - 1
			byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, 8)
		}
	}
	return byteCode, nil
}

func ParseInstructions(byteCode []byte) []Instruction {
	instructions := []Instruction{}

	for len(byteCode) > 0 {
		var instruction Instruction

		opCode := byteCode[0]
		byteCode = advanceInstruction(byteCode, 1)

		switch opCode {
		case enums.OP_CODE_BYTE_ADD:
			byteCode, instruction = parseAdd(byteCode)
		case enums.OP_CODE_BYTE_SUBTRACT:
			byteCode, instruction = parseSubtract(byteCode)
		case enums.OP_CODE_BYTE_MULTIPLY:
			byteCode, instruction = parseMultiply(byteCode)
		case enums.OP_CODE_BYTE_DIVIDE:
			byteCode, instruction = parseDivide(byteCode)
		case enums.OP_CODE_BYTE_MODULO:
			byteCode, instruction = parseModulo(byteCode)
		case enums.OP_CODE_BYTE_POWER:
			byteCode, instruction = parsePower(byteCode)
		case enums.OP_CODE_BYTE_BINARY_ADD:
			byteCode, instruction = parseBinAdd(byteCode)
		case enums.OP_CODE_BYTE_BINARY_SUBTRACT:
			byteCode, instruction = parseBinSubtract(byteCode)
		case enums.OP_CODE_BYTE_GREATER:
			byteCode, instruction = parseGreater(byteCode)
		case enums.OP_CODE_BYTE_GREATER_EQUAL:
			byteCode, instruction = parseGreaterEqual(byteCode)
		case enums.OP_CODE_BYTE_EQUAL:
			byteCode, instruction = parseEqual(byteCode)
		case enums.OP_CODE_BYTE_LESS_EQUAL:
			byteCode, instruction = parseLessEqual(byteCode)
		case enums.OP_CODE_BYTE_LESS:
			byteCode, instruction = parseLess(byteCode)
		case enums.OP_CODE_BYTE_INPUT_BLOCK:
			byteCode, instruction = parseInputBlock(byteCode)
		case enums.OP_CODE_BYTE_INPUT_NON_BLOCK:
			byteCode, instruction = parseInputNonBlock(byteCode)
		case enums.OP_CODE_BYTE_OUTPUT:
			byteCode, instruction = parseOutput(byteCode)
		case enums.OP_CODE_BYTE_BRANCH_POSITIVE:
			byteCode, instruction = parseBranchPositive(byteCode)
		case enums.OP_CODE_BYTE_BRANCH_NOT_POSITIVE:
			byteCode, instruction = parseBranchNotPositive(byteCode)
		case enums.OP_CODE_BYTE_BRANCH_ZERO:
			byteCode, instruction = parseBranchZero(byteCode)
		case enums.OP_CODE_BYTE_BRANCH_NOT_ZERO:
			byteCode, instruction = parseBranchNotZero(byteCode)
		case enums.OP_CODE_BYTE_BRANCH_NEGATIVE:
			byteCode, instruction = parseBranchNegative(byteCode)
		case enums.OP_CODE_BYTE_BRANCH_NOT_NEGATIVE:
			byteCode, instruction = parseBranchNotNegative(byteCode)
		case enums.OP_CODE_BYTE_GOTO:
			byteCode, instruction = parseGoto(byteCode)
		case enums.OP_CODE_BYTE_STOP:
			byteCode, instruction = parseStop(byteCode)
		}
		instructions = append(instructions, instruction)
	}
	return instructions

}

func parseAdd(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_ADD

	return byteCode, instruction
}

func parseSubtract(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_SUBTRACT

	return byteCode, instruction
}

func parseMultiply(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_MULTIPLY

	return byteCode, instruction
}

func parseDivide(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_DIVIDE

	return byteCode, instruction
}

func parseModulo(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_MODULO

	return byteCode, instruction
}

func parsePower(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_POWER

	return byteCode, instruction
}

func parseBinAdd(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_BINARY_ADD

	return byteCode, instruction
}

func parseBinSubtract(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_BINARY_SUBTRACT

	return byteCode, instruction
}

func parseGreater(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_GREATER

	return byteCode, instruction
}

func parseGreaterEqual(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_GREATER_EQUAL

	return byteCode, instruction
}

func parseEqual(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_EQUAL

	return byteCode, instruction
}

func parseLessEqual(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_LESS_EQUAL

	return byteCode, instruction
}

func parseLess(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg3(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_LESS

	return byteCode, instruction
}

func parseInputBlock(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg1(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_INPUT_BLOCK

	return byteCode, instruction
}

func parseInputNonBlock(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg2(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_INPUT_NON_BLOCK

	return byteCode, instruction
}

func parseOutput(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg1(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_OUTPUT

	return byteCode, instruction
}

func parseBranchPositive(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg2(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_BRANCH_POSITIVE

	return byteCode, instruction
}

func parseBranchNotPositive(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg2(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_BRANCH_NOT_POSITIVE

	return byteCode, instruction
}

func parseBranchZero(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg2(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_BRANCH_ZERO

	return byteCode, instruction
}

func parseBranchNotZero(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg2(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_BRANCH_NOT_ZERO

	return byteCode, instruction
}

func parseBranchNegative(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg2(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_BRANCH_NEGATIVE

	return byteCode, instruction
}

func parseBranchNotNegative(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg2(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_BRANCH_NOT_NEGATIVE

	return byteCode, instruction
}

func parseGoto(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg1(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_GOTO

	return byteCode, instruction
}

func parseStop(byteCode []byte) ([]byte, Instruction) {
	byteCode, instruction := getArg1(byteCode)

	instruction.OpCode = enums.OP_CODE_BYTE_STOP

	return byteCode, instruction
}

func getArg1(byteCode []byte) ([]byte, Instruction) {
	aLength := utils.CoerceByteInt(byteCode[:1])
	byteCode = advanceInstruction(byteCode, 1)

	a := byteCode[:aLength]
	byteCode = advanceInstruction(byteCode, aLength)

	args := [][]byte{}

	args = append(args, a)

	instruction := Instruction{
		Args:      args,
		ArgLength: 1,
	}

	return byteCode, instruction
}

func getArg2(byteCode []byte) ([]byte, Instruction) {
	aLength := utils.CoerceByteInt(byteCode[:1])
	byteCode = advanceInstruction(byteCode, 1)

	a := byteCode[:aLength]
	byteCode = advanceInstruction(byteCode, aLength)

	bLength := utils.CoerceByteInt(byteCode[:1])
	byteCode = advanceInstruction(byteCode, 1)

	b := byteCode[:bLength]
	byteCode = advanceInstruction(byteCode, bLength)

	args := [][]byte{}

	args = append(args, a)
	args = append(args, b)

	instruction := Instruction{
		Args:      args,
		ArgLength: 2,
	}

	return byteCode, instruction
}

func getArg3(byteCode []byte) ([]byte, Instruction) {
	aLength := utils.CoerceByteInt(byteCode[:1])
	byteCode = advanceInstruction(byteCode, 1)

	a := byteCode[:aLength]
	byteCode = advanceInstruction(byteCode, aLength)

	bLength := utils.CoerceByteInt(byteCode[:1])
	byteCode = advanceInstruction(byteCode, 1)

	b := byteCode[:bLength]
	byteCode = advanceInstruction(byteCode, bLength)

	cLength := utils.CoerceByteInt(byteCode[:1])
	byteCode = advanceInstruction(byteCode, 1)

	c := byteCode[:cLength]
	byteCode = advanceInstruction(byteCode, cLength)

	args := [][]byte{}

	args = append(args, a)
	args = append(args, b)
	args = append(args, c)

	instruction := Instruction{
		Args:      args,
		ArgLength: 3,
	}

	return byteCode, instruction
}
