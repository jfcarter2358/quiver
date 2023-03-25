package dotcodes

import (
	"quiverc/assembler/utils"
	"quiverc/enums"
)

func ProcessLiteral(parts []string) ([]byte, error) {
	var returnBytes []byte

	dotCode := enums.DOT_CODE_BYTE_LITERAL

	name := parts[1]
	nameLength := utils.IntToByteArray(len(name), 1)

	data := []byte(parts[2])
	dataLength := utils.IntToByteArray(len(data), 8)

	// names = append(names, name)
	dataType, err := utils.GetDataType(parts[2])
	if err != nil {
		// return nil, nil, err
		return nil, err
	}

	returnBytes = append(returnBytes, dotCode)
	returnBytes = append(returnBytes, dataType)
	returnBytes = append(returnBytes, nameLength...)
	returnBytes = append(returnBytes, []byte(name)...)
	returnBytes = append(returnBytes, dataLength...)
	returnBytes = append(returnBytes, data...)

	return returnBytes, nil
}

func ProcessLabel(parts []string, programCounter int) []byte {
	var returnBytes []byte

	dotCode := enums.DOT_CODE_BYTE_LABEL

	name := parts[1]
	nameLength := utils.IntToByteArray(len(name), 1)

	line := utils.IntToByteArray(programCounter+1, 8)

	returnBytes = append(returnBytes, dotCode)
	returnBytes = append(returnBytes, nameLength...)
	returnBytes = append(returnBytes, []byte(name)...)
	returnBytes = append(returnBytes, line...)

	return returnBytes
}
