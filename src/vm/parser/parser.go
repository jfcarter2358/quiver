package parser

import (
	"vm/enums"
	"vm/memstore"
	"vm/utils"
)

func advanceByteCode(byteCode []byte, byteCounter, advance int) ([]byte, int) {
	byteCode = byteCode[advance:]
	byteCounter = byteCounter + advance

	return byteCode, byteCounter
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

			memstore.LabelData[labelName] = programLine
			byteCode, byteCounter = advanceByteCode(byteCode, byteCounter, 8)
		}
	}
	return byteCode, nil
}
