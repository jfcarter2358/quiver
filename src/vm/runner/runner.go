package runner

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"vm/enums"
	"vm/memstore"
	"vm/parser"
	"vm/utils"
)

func Run(instructions []parser.Instruction) error {
	var err error
	programCounter := 0

	for programCounter < len(instructions) {
		instruction := instructions[programCounter]

		switch instruction.OpCode {
		case enums.OP_CODE_BYTE_ADD:
			programCounter, err = doAdd(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_SUBTRACT:
			programCounter, err = doSubtract(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_MULTIPLY:
			programCounter, err = doMultiply(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_DIVIDE:
			programCounter, err = doDivide(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_MODULO:
			programCounter, err = doModulo(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_POWER:
			programCounter, err = doPower(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BINARY_ADD:
			programCounter, err = doBinaryAdd(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BINARY_SUBTRACT:
			programCounter, err = doBinarySubtract(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_GREATER:
			programCounter, err = doGreater(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_GREATER_EQUAL:
			programCounter, err = doGreaterEqual(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_EQUAL:
			programCounter, err = doEqual(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_LESS_EQUAL:
			programCounter, err = doLessEqual(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_LESS:
			programCounter, err = doLess(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_INPUT_BLOCK:
			programCounter, err = doInputBlock(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_INPUT_NON_BLOCK:
			programCounter, err = doInputNonBlock(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_OUTPUT:
			programCounter, err = doOutput(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_POSITIVE:
			programCounter, err = doBranchPositive(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_NOT_POSITIVE:
			programCounter, err = doBranchNotPositive(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_ZERO:
			programCounter, err = doBranchZero(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_NOT_ZERO:
			programCounter, err = doBranchNotZero(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_NEGATIVE:
			programCounter, err = doBranchNegative(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_NOT_NEGATIVE:
			programCounter, err = doBranchNotNegative(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_GOTO:
			programCounter, err = doGoto(instruction, programCounter)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_STOP:
			programCounter, err = doStop(instruction, programCounter)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func doAdd(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)
	dataTypeC := utils.GetVariableDataType(c)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in add operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in add operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		memstore.IntData[c] = memstore.IntData[a] + memstore.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		memstore.FloatData[c] = memstore.FloatData[a] + memstore.FloatData[b]
	case enums.DATATYPE_BYTE_STRING:
		memstore.StringData[c] = memstore.StringData[a] + memstore.StringData[b]
	default:
		return 0, fmt.Errorf("add operation is not yet implemented for type %s", dataTypeA)
	}

	programCounter += 1

	return programCounter, nil
}

func doSubtract(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)
	dataTypeC := utils.GetVariableDataType(c)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in subtract operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in subtract operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		memstore.IntData[c] = memstore.IntData[a] - memstore.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		memstore.FloatData[c] = memstore.FloatData[a] - memstore.FloatData[b]
	default:
		return 0, fmt.Errorf("subtract operation is not yet implemented for type %s", dataTypeA)
	}

	programCounter += 1

	return programCounter, nil
}

func doMultiply(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)
	dataTypeC := utils.GetVariableDataType(c)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in multiply operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in multiply operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		memstore.IntData[c] = memstore.IntData[a] * memstore.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		memstore.FloatData[c] = memstore.FloatData[a] * memstore.FloatData[b]
	default:
		return 0, fmt.Errorf("multiply operation is not yet implemented for type %s", dataTypeA)
	}

	programCounter += 1

	return programCounter, nil
}

func doDivide(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)
	dataTypeC := utils.GetVariableDataType(c)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in divide operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in divide operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		memstore.IntData[c] = memstore.IntData[a] / memstore.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		memstore.FloatData[c] = memstore.FloatData[a] / memstore.FloatData[b]
	default:
		return 0, fmt.Errorf("divide operation is not yet implemented for type %s", dataTypeA)
	}

	programCounter += 1

	return programCounter, nil
}

func doModulo(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)
	dataTypeC := utils.GetVariableDataType(c)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in modulo operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in modulo operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		memstore.IntData[c] = memstore.IntData[a] % memstore.IntData[b]
	default:
		return 0, fmt.Errorf("modulo operation is not yet implemented for type %s", dataTypeA)
	}

	programCounter += 1

	return programCounter, nil
}

func doPower(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)
	dataTypeC := utils.GetVariableDataType(c)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in power operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in add operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		memstore.IntData[c] = int(math.Pow(float64(memstore.IntData[a]), float64(memstore.IntData[b])))
	case enums.DATATYPE_BYTE_FLOAT:
		memstore.FloatData[c] = math.Pow(memstore.FloatData[a], memstore.FloatData[b])
	default:
		return 0, fmt.Errorf("power operation is not yet implemented for type %s", dataTypeA)
	}

	programCounter += 1

	return programCounter, nil
}

func doBinaryAdd(instruction parser.Instruction, programCounter int) (int, error) {
	// a := string(instruction.Args[0])
	// b := string(instruction.Args[1])
	// c := string(instruction.Args[2])

	// dataTypeA := utils.GetVariableDataType(a)
	// dataTypeB := utils.GetVariableDataType(b)
	// dataTypeC := utils.GetVariableDataType(c)

	// dataTypeAString := enums.ByteToDataType(dataTypeA)
	// dataTypeBString := enums.ByteToDataType(dataTypeB)
	// dataTypeCString := enums.ByteToDataType(dataTypeC)

	// if dataTypeA != dataTypeB {
	// 	return 0, fmt.Errorf("data type mismatch in add operation: %s, %s", dataTypeAString, dataTypeBString)
	// }

	// if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
	// 	return 0, fmt.Errorf("data type mismatch in add operation storage: %s, %s", dataTypeAString, dataTypeCString)
	// }

	// switch dataTypeA {
	// case enums.DATATYPE_BYTE_INT:
	// 	memstore.IntData[c] = memstore.IntData[a] + memstore.IntData[b]
	// case enums.DATATYPE_BYTE_FLOAT:
	// 	memstore.FloatData[c] = memstore.FloatData[a] + memstore.FloatData[b]
	// case enums.DATATYPE_BYTE_STRING:
	// 	memstore.StringData[c] = memstore.StringData[a] + memstore.StringData[b]
	// default:
	// 	return 0, fmt.Errorf("add operation is not yet implemented for type %s", dataTypeA)
	// }

	programCounter += 1

	return programCounter, nil
}

func doBinarySubtract(instruction parser.Instruction, programCounter int) (int, error) {
	// a := string(instruction.Args[0])
	// b := string(instruction.Args[1])
	// c := string(instruction.Args[2])

	// dataTypeA := utils.GetVariableDataType(a)
	// dataTypeB := utils.GetVariableDataType(b)
	// dataTypeC := utils.GetVariableDataType(c)

	// dataTypeAString := enums.ByteToDataType(dataTypeA)
	// dataTypeBString := enums.ByteToDataType(dataTypeB)
	// dataTypeCString := enums.ByteToDataType(dataTypeC)

	// if dataTypeA != dataTypeB {
	// 	return 0, fmt.Errorf("data type mismatch in add operation: %s, %s", dataTypeAString, dataTypeBString)
	// }

	// if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
	// 	return 0, fmt.Errorf("data type mismatch in add operation storage: %s, %s", dataTypeAString, dataTypeCString)
	// }

	// switch dataTypeA {
	// case enums.DATATYPE_BYTE_INT:
	// 	memstore.IntData[c] = memstore.IntData[a] + memstore.IntData[b]
	// case enums.DATATYPE_BYTE_FLOAT:
	// 	memstore.FloatData[c] = memstore.FloatData[a] + memstore.FloatData[b]
	// case enums.DATATYPE_BYTE_STRING:
	// 	memstore.StringData[c] = memstore.StringData[a] + memstore.StringData[b]
	// default:
	// 	return 0, fmt.Errorf("add operation is not yet implemented for type %s", dataTypeA)
	// }

	programCounter += 1

	return programCounter, nil
}

func doGreater(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != enums.DATATYPE_BYTE_INT || dataTypeA != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in add operation: %s", dataTypeAString)
	}
	if dataTypeB != enums.DATATYPE_BYTE_INT || dataTypeB != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in add operation: %s", dataTypeBString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		switch dataTypeB {
		case enums.DATATYPE_BYTE_INT:
			memstore.BoolData[c] = memstore.IntData[a] > memstore.IntData[b]
		case enums.DATATYPE_BYTE_FLOAT:
			memstore.BoolData[c] = float64(memstore.IntData[a]) > memstore.FloatData[b]
		}
	case enums.DATATYPE_BYTE_FLOAT:
		switch dataTypeB {
		case enums.DATATYPE_BYTE_INT:
			memstore.BoolData[c] = memstore.FloatData[a] > float64(memstore.IntData[b])
		case enums.DATATYPE_BYTE_FLOAT:
			memstore.BoolData[c] = memstore.FloatData[a] > memstore.FloatData[b]
		}
	}

	programCounter += 1

	return programCounter, nil
}

func doGreaterEqual(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != enums.DATATYPE_BYTE_INT || dataTypeA != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in add operation: %s", dataTypeAString)
	}
	if dataTypeB != enums.DATATYPE_BYTE_INT || dataTypeB != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in greater equal operation: %s", dataTypeBString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		switch dataTypeB {
		case enums.DATATYPE_BYTE_INT:
			memstore.BoolData[c] = memstore.IntData[a] >= memstore.IntData[b]
		case enums.DATATYPE_BYTE_FLOAT:
			memstore.BoolData[c] = float64(memstore.IntData[a]) >= memstore.FloatData[b]
		}
	case enums.DATATYPE_BYTE_FLOAT:
		switch dataTypeB {
		case enums.DATATYPE_BYTE_INT:
			memstore.BoolData[c] = memstore.FloatData[a] >= float64(memstore.IntData[b])
		case enums.DATATYPE_BYTE_FLOAT:
			memstore.BoolData[c] = memstore.FloatData[a] >= memstore.FloatData[b]
		}
	}

	programCounter += 1

	return programCounter, nil
}

func doEqual(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in equal operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		memstore.BoolData[c] = memstore.BoolData[a] == memstore.BoolData[b]
	case enums.DATATYPE_BYTE_INT:
		memstore.BoolData[c] = memstore.IntData[a] == memstore.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		memstore.BoolData[c] = memstore.FloatData[a] == memstore.FloatData[b]
	case enums.DATATYPE_BYTE_STRING:
		memstore.BoolData[c] = memstore.StringData[a] == memstore.StringData[b]
	default:
		return 0, fmt.Errorf("equal operation is not yet implemented for type %s", dataTypeA)
	}

	programCounter += 1

	return programCounter, nil
}

func doLessEqual(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != enums.DATATYPE_BYTE_INT || dataTypeA != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in less equal operation: %s", dataTypeAString)
	}
	if dataTypeB != enums.DATATYPE_BYTE_INT || dataTypeB != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in less equal operation: %s", dataTypeBString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		switch dataTypeB {
		case enums.DATATYPE_BYTE_INT:
			memstore.BoolData[c] = memstore.IntData[a] <= memstore.IntData[b]
		case enums.DATATYPE_BYTE_FLOAT:
			memstore.BoolData[c] = float64(memstore.IntData[a]) <= memstore.FloatData[b]
		}
	case enums.DATATYPE_BYTE_FLOAT:
		switch dataTypeB {
		case enums.DATATYPE_BYTE_INT:
			memstore.BoolData[c] = memstore.FloatData[a] <= float64(memstore.IntData[b])
		case enums.DATATYPE_BYTE_FLOAT:
			memstore.BoolData[c] = memstore.FloatData[a] <= memstore.FloatData[b]
		}
	}

	programCounter += 1

	return programCounter, nil
}

func doLess(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a)
	dataTypeB := utils.GetVariableDataType(b)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != enums.DATATYPE_BYTE_INT || dataTypeA != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in less operation: %s", dataTypeAString)
	}
	if dataTypeB != enums.DATATYPE_BYTE_INT || dataTypeB != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in less operation: %s", dataTypeBString)
	}

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		switch dataTypeB {
		case enums.DATATYPE_BYTE_INT:
			memstore.BoolData[c] = memstore.IntData[a] < memstore.IntData[b]
		case enums.DATATYPE_BYTE_FLOAT:
			memstore.BoolData[c] = float64(memstore.IntData[a]) < memstore.FloatData[b]
		}
	case enums.DATATYPE_BYTE_FLOAT:
		switch dataTypeB {
		case enums.DATATYPE_BYTE_INT:
			memstore.BoolData[c] = memstore.FloatData[a] < float64(memstore.IntData[b])
		case enums.DATATYPE_BYTE_FLOAT:
			memstore.BoolData[c] = memstore.FloatData[a] < memstore.FloatData[b]
		}
	}

	programCounter += 1

	return programCounter, nil
}

func doInputBlock(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])

	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')

	// convert CRLF to LF
	inputString = strings.Replace(inputString, "\n", "", -1)

	memstore.StringData[a] = inputString

	programCounter += 1

	return programCounter, nil
}

func doInputNonBlock(instruction parser.Instruction, programCounter int) (int, error) {
	// a := string(instruction.Args[0])
	// b := string(instruction.Args[1])
	// c := string(instruction.Args[2])

	// dataTypeA := utils.GetVariableDataType(a)
	// dataTypeB := utils.GetVariableDataType(b)

	// dataTypeAString := enums.ByteToDataType(dataTypeA)
	// dataTypeBString := enums.ByteToDataType(dataTypeB)

	// if dataTypeA != enums.DATATYPE_BYTE_INT || dataTypeA != enums.DATATYPE_BYTE_FLOAT {
	// 	return 0, fmt.Errorf("data type invalid in less operation: %s", dataTypeAString)
	// }
	// if dataTypeB != enums.DATATYPE_BYTE_INT || dataTypeB != enums.DATATYPE_BYTE_FLOAT {
	// 	return 0, fmt.Errorf("data type invalid in less operation: %s", dataTypeBString)
	// }

	// switch dataTypeA {
	// case enums.DATATYPE_BYTE_INT:
	// 	switch dataTypeB {
	// 	case enums.DATATYPE_BYTE_INT:
	// 		memstore.BoolData[c] = memstore.IntData[a] < memstore.IntData[b]
	// 	case enums.DATATYPE_BYTE_FLOAT:
	// 		memstore.BoolData[c] = memstore.IntData[a] < memstore.FloatData[b]
	// 	}
	// case enums.DATATYPE_BYTE_FLOAT:
	// 	switch dataTypeB {
	// 	case enums.DATATYPE_BYTE_INT:
	// 		memstore.BoolData[c] = memstore.FloatData[a] < memstore.IntData[b]
	// 	case enums.DATATYPE_BYTE_FLOAT:
	// 		memstore.BoolData[c] = memstore.FloatData[a] < memstore.FloatData[b]
	// 	}
	// }

	programCounter += 1

	return programCounter, nil
}

func doOutput(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])

	dataTypeA := utils.GetVariableDataType(a)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		fmt.Printf("%v", memstore.BoolData[a])
	case enums.DATATYPE_BYTE_INT:
		fmt.Printf("%v", memstore.IntData[a])
	case enums.DATATYPE_BYTE_FLOAT:
		fmt.Printf("%v", memstore.FloatData[a])
	case enums.DATATYPE_BYTE_STRING:
		fmt.Printf("%v", memstore.StringData[a])
	case enums.DATATYPE_BYTE_LIST:
		fmt.Printf("%v", memstore.ListData[a])
	case enums.DATATYPE_BYTE_DICT:
		fmt.Printf("%v", memstore.DictData[a])
	}

	programCounter += 1

	return programCounter, nil
}

func doBranchPositive(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		if memstore.BoolData[a] {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_INT:
		if memstore.IntData[a] > 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if memstore.FloatData[a] > 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch positive operation is not defined for type %s", dataTypeA)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchNotPositive(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		if !memstore.BoolData[a] {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_INT:
		if memstore.IntData[a] <= 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if memstore.FloatData[a] <= 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch not positive operation is not defined for type %s", dataTypeA)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchZero(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		if !memstore.BoolData[a] {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_INT:
		if memstore.IntData[a] == 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if memstore.FloatData[a] == 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch zero operation is not defined for type %s", dataTypeA)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchNotZero(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		if memstore.BoolData[a] {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_INT:
		if memstore.IntData[a] != 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if memstore.FloatData[a] != 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch not zero operation is not defined for type %s", dataTypeA)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchNegative(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		if memstore.IntData[a] < 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if memstore.FloatData[a] < 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch negative operation is not defined for type %s", dataTypeA)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchNotNegative(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		if memstore.IntData[a] >= 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if memstore.FloatData[a] >= 0 {
			programCounter = memstore.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch not negative operation is not defined for type %s", dataTypeA)
	}

	programCounter += 1
	return programCounter, nil

}

func doGoto(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])

	programCounter = memstore.LabelData[a]
	return programCounter, nil
}

func doStop(instruction parser.Instruction, programCounter int) (int, error) {
	a := string(instruction.Args[0])

	returnCode := memstore.IntData[a]

	os.Exit(returnCode)

	return 0, nil
}
