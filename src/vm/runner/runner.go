package runner

import (
	"fmt"
	"os"
	"vm/enums"
	"vm/memstore"
	"vm/utils"
)

func advanceInstruction(instructions []byte, advance int) []byte {
	instructions = instructions[advance:]

	return instructions
}

func Run(instructions []byte) error {
	var err error

	for len(instructions) > 0 {
		opCode := instructions[0]
		instructions = advanceInstruction(instructions, 1)

		switch opCode {
		case enums.OP_CODE_BYTE_ADD:
			instructions, err = doAdd(instructions)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_OUTPUT:
			instructions = doOutput(instructions)
		case enums.OP_CODE_BYTE_STOP:
			doStop(instructions)
		}
	}
	return nil
}

func doAdd(instructions []byte) ([]byte, error) {
	source1Length := utils.CoerceByteInt(instructions[:1])
	instructions = advanceInstruction(instructions, 1)

	source1 := utils.CoerceString(instructions[:source1Length])
	instructions = advanceInstruction(instructions, source1Length)

	source2Length := utils.CoerceByteInt(instructions[:1])
	instructions = advanceInstruction(instructions, 1)

	source2 := utils.CoerceString(instructions[:source2Length])
	instructions = advanceInstruction(instructions, source2Length)

	destLength := utils.CoerceByteInt(instructions[:1])
	instructions = advanceInstruction(instructions, 1)

	dest := utils.CoerceString(instructions[:destLength])
	instructions = advanceInstruction(instructions, destLength)

	dataTypeSource1 := utils.GetVariableDataType(source1)
	dataTypeSource2 := utils.GetVariableDataType(source2)
	dataTypeDest := utils.GetVariableDataType(dest)

	if dataTypeSource1 != dataTypeSource2 {
		return nil, fmt.Errorf("data types %s and %s do not match in add operation", enums.ByteToDataType(dataTypeSource1), enums.ByteToDataType(dataTypeSource2))
	}
	if dataTypeDest != enums.DATATYPE_BYTE_NULL && dataTypeDest != dataTypeSource1 {
		return nil, fmt.Errorf("destination data type %s does not match operand type %s", enums.ByteToDataType(dataTypeDest), enums.ByteToDataType(dataTypeSource1))
	}

	switch dataTypeSource1 {
	case enums.DATATYPE_BYTE_BOOL:
		return nil, fmt.Errorf("add operation is not implemented for data type %s", enums.ByteToDataType(dataTypeSource1))
	case enums.DATATYPE_BYTE_INT:
		memstore.IntData[dest] = memstore.IntData[source1] + memstore.IntData[source2]
	case enums.DATATYPE_BYTE_FLOAT:
		memstore.FloatData[dest] = memstore.FloatData[source1] + memstore.FloatData[source2]
	case enums.DATATYPE_BYTE_STRING:
		memstore.StringData[dest] = memstore.StringData[source1] + memstore.StringData[source2]
	case enums.DATATYPE_BYTE_LIST:
		return nil, fmt.Errorf("add operation is not implemented for data type %s", enums.ByteToDataType(dataTypeSource1))
	case enums.DATATYPE_BYTE_DICT:
		return nil, fmt.Errorf("add operation is not implemented for data type %s", enums.ByteToDataType(dataTypeSource1))
	}

	return instructions, nil
}

func doOutput(instructions []byte) []byte {
	sourceLength := utils.CoerceByteInt(instructions[:1])
	instructions = advanceInstruction(instructions, 1)

	source := utils.CoerceString(instructions[:sourceLength])
	instructions = advanceInstruction(instructions, sourceLength)

	dataType := utils.GetVariableDataType(source)

	switch dataType {
	case enums.DATATYPE_BYTE_BOOL:
		fmt.Printf("%v", memstore.BoolData[source])
	case enums.DATATYPE_BYTE_INT:
		fmt.Printf("%v", memstore.IntData[source])
	case enums.DATATYPE_BYTE_FLOAT:
		fmt.Printf("%v", memstore.FloatData[source])
	case enums.DATATYPE_BYTE_STRING:
		fmt.Printf("%v", memstore.StringData[source])
	case enums.DATATYPE_BYTE_LIST:
		fmt.Printf("%v", memstore.ListData[source])
	case enums.DATATYPE_BYTE_DICT:
		fmt.Printf("%v", memstore.DictData[source])
	}

	return instructions
}

func doStop(instructions []byte) {
	returnCode := utils.CoerceByteInt(instructions[:1])
	fmt.Println("Bailing!")
	os.Exit(returnCode)
}
