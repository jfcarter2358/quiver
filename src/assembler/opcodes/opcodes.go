package opcodes

import (
	"assembler/enums"
	"assembler/utils"
	"strconv"
)

func ProcessAdd(parts []string) []byte {
	var returnBytes []byte

	opCode := enums.OP_CODE_BYTE_ADD

	source1 := parts[1]
	sourceLength1 := utils.IntToByteArray(len(source1), 1)

	source2 := parts[2]
	sourceLength2 := utils.IntToByteArray(len(source2), 1)

	dest := parts[3]
	destLength := utils.IntToByteArray(len(dest), 1)

	returnBytes = []byte{opCode}
	returnBytes = append(returnBytes, sourceLength1...)
	returnBytes = append(returnBytes, []byte(source1)...)
	returnBytes = append(returnBytes, sourceLength2...)
	returnBytes = append(returnBytes, []byte(source2)...)
	returnBytes = append(returnBytes, destLength...)
	returnBytes = append(returnBytes, []byte(dest)...)

	return returnBytes
}

func ProcessStop(parts []string) []byte {
	var returnBytes []byte

	opCode := enums.OP_CODE_BYTE_STOP

	returnCodeInt, _ := strconv.Atoi(parts[1])
	returnCode := utils.IntToByteArray(returnCodeInt, 1)

	returnBytes = []byte{opCode}
	returnBytes = append(returnBytes, returnCode...)

	return returnBytes
}

func ProcessOutput(parts []string) []byte {
	var returnBytes []byte

	opCode := enums.OP_CODE_BYTE_OUTPUT

	source := parts[1]
	sourceLength := utils.IntToByteArray(len(source), 1)

	returnBytes = []byte{opCode}
	returnBytes = append(returnBytes, sourceLength...)
	returnBytes = append(returnBytes, []byte(source)...)

	return returnBytes
}
