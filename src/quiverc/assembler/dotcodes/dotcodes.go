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

	// isDict := false
	// isList := false

	// dictKeyType := byte(0)
	// dictValType := byte(0)

	// listType := byte(0)

	// if dataType == enums.DATATYPE_BYTE_DICT {
	// 	isDict = true

	// 	switch parts[3] {
	// 	case enums.DATATYPE_NAME_BOOL:
	// 		dictKeyType = enums.DATATYPE_BYTE_BOOL
	// 	case enums.DATATYPE_NAME_FLOAT:
	// 		dictKeyType = enums.DATATYPE_BYTE_BOOL
	// 	case enums.DATATYPE_NAME_INT:
	// 		dictKeyType = enums.DATATYPE_BYTE_INT
	// 	case enums.DATATYPE_NAME_STRING:
	// 		dictKeyType = enums.DATATYPE_BYTE_STRING
	// 	}

	// 	switch parts[4] {
	// 	case enums.DATATYPE_NAME_BOOL:
	// 		dictValType = enums.DATATYPE_BYTE_BOOL
	// 	case enums.DATATYPE_NAME_INT:
	// 		dictValType = enums.DATATYPE_BYTE_INT
	// 	case enums.DATATYPE_NAME_FLOAT:
	// 		dictValType = enums.DATATYPE_BYTE_FLOAT
	// 	case enums.DATATYPE_NAME_STRING:
	// 		dictValType = enums.DATATYPE_BYTE_STRING
	// 	case enums.DATATYPE_NAME_DICT:
	// 		dictValType = enums.DATATYPE_BYTE_DICT
	// 	case enums.DATATYPE_NAME_LIST:
	// 		dictValType = enums.DATATYPE_BYTE_LIST
	// 	}
	// }

	// if dataType == enums.DATATYPE_BYTE_LIST {
	// 	isList = true

	// 	switch parts[3] {
	// 	case enums.DATATYPE_NAME_BOOL:
	// 		listType = enums.DATATYPE_BYTE_BOOL
	// 	case enums.DATATYPE_NAME_INT:
	// 		listType = enums.DATATYPE_BYTE_INT
	// 	case enums.DATATYPE_NAME_FLOAT:
	// 		listType = enums.DATATYPE_BYTE_FLOAT
	// 	case enums.DATATYPE_NAME_STRING:
	// 		listType = enums.DATATYPE_BYTE_STRING
	// 	case enums.DATATYPE_NAME_DICT:
	// 		listType = enums.DATATYPE_BYTE_DICT
	// 	case enums.DATATYPE_NAME_LIST:
	// 		listType = enums.DATATYPE_BYTE_LIST
	// 	}
	// }

	returnBytes = append(returnBytes, dotCode)
	returnBytes = append(returnBytes, dataType)
	returnBytes = append(returnBytes, nameLength...)
	returnBytes = append(returnBytes, []byte(name)...)
	returnBytes = append(returnBytes, dataLength...)
	returnBytes = append(returnBytes, data...)

	// if isDict {
	// 	returnBytes = append(returnBytes, dictKeyType)
	// 	returnBytes = append(returnBytes, dictValType)
	// }

	// if isList {
	// 	returnBytes = append(returnBytes, listType)
	// }

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
