package runner

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"quiver/enums"
	"quiver/vm/memstore"
	"quiver/vm/parser"
	"quiver/vm/utils"
	"strings"
)

func Run(instructions []parser.Instruction, vars *memstore.VariableStore) error {
	var err error
	programCounter := 0

	for programCounter < len(instructions) {
		instruction := instructions[programCounter]

		switch instruction.OpCode {
		case enums.OP_CODE_BYTE_ADD:
			programCounter, err = doAdd(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_SUBTRACT:
			programCounter, err = doSubtract(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_MULTIPLY:
			programCounter, err = doMultiply(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_DIVIDE:
			programCounter, err = doDivide(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_MODULO:
			programCounter, err = doModulo(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_POWER:
			programCounter, err = doPower(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BINARY_ADD:
			programCounter, err = doBinaryAdd(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BINARY_SUBTRACT:
			programCounter, err = doBinarySubtract(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_GREATER:
			programCounter, err = doGreater(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_GREATER_EQUAL:
			programCounter, err = doGreaterEqual(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_EQUAL:
			programCounter, err = doEqual(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_LESS_EQUAL:
			programCounter, err = doLessEqual(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_LESS:
			programCounter, err = doLess(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_INPUT_BLOCK:
			programCounter, err = doInputBlock(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_INPUT_NON_BLOCK:
			programCounter, err = doInputNonBlock(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_OUTPUT:
			programCounter, err = doOutput(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_POSITIVE:
			programCounter, err = doBranchPositive(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_NOT_POSITIVE:
			programCounter, err = doBranchNotPositive(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_ZERO:
			programCounter, err = doBranchZero(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_NOT_ZERO:
			programCounter, err = doBranchNotZero(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_NEGATIVE:
			programCounter, err = doBranchNegative(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_BRANCH_NOT_NEGATIVE:
			programCounter, err = doBranchNotNegative(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_GOTO:
			programCounter, err = doGoto(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_STOP:
			programCounter, err = doStop(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_COPY:
			programCounter, err = doCopy(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_DICT_ASSIGN:
			programCounter, err = doDictAssign(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_DICT_ACCESS:
			programCounter, err = doDictAccess(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_LIST_ASSIGN:
			programCounter, err = doListAssign(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		case enums.OP_CODE_BYTE_LIST_ACCESS:
			programCounter, err = doListAccess(instruction, programCounter, vars)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func doAdd(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)
	dataTypeC := utils.GetVariableDataType(c, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in add operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in add operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.IntData[c] = aVars.IntData[a] + bVars.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.FloatData[c] = aVars.FloatData[a] + bVars.FloatData[b]
	case enums.DATATYPE_BYTE_STRING:
		cVars.StringData[c] = aVars.StringData[a] + bVars.StringData[b]
	default:
		return 0, fmt.Errorf("add operation is not yet implemented for type %s", dataTypeAString)
	}

	programCounter += 1

	return programCounter, nil
}

func doSubtract(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)
	dataTypeC := utils.GetVariableDataType(c, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in subtract operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in subtract operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.IntData[c] = aVars.IntData[a] - bVars.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.FloatData[c] = aVars.FloatData[a] - bVars.FloatData[b]
	default:
		return 0, fmt.Errorf("subtract operation is not yet implemented for type %s", dataTypeAString)
	}

	programCounter += 1

	return programCounter, nil
}

func doMultiply(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)
	dataTypeC := utils.GetVariableDataType(c, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in multiply operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in multiply operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.IntData[c] = aVars.IntData[a] * bVars.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.FloatData[c] = aVars.FloatData[a] * bVars.FloatData[b]
	default:
		return 0, fmt.Errorf("multiply operation is not yet implemented for type %s", dataTypeAString)
	}

	programCounter += 1

	return programCounter, nil
}

func doDivide(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)
	dataTypeC := utils.GetVariableDataType(c, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in divide operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in divide operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.IntData[c] = aVars.IntData[a] / bVars.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.FloatData[c] = aVars.FloatData[a] / bVars.FloatData[b]
	default:
		return 0, fmt.Errorf("divide operation is not yet implemented for type %s", dataTypeAString)
	}

	programCounter += 1

	return programCounter, nil
}

func doModulo(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)
	dataTypeC := utils.GetVariableDataType(c, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in modulo operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in modulo operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.IntData[c] = aVars.IntData[a] % bVars.IntData[b]
	default:
		return 0, fmt.Errorf("modulo operation is not yet implemented for type %s", dataTypeAString)
	}

	programCounter += 1

	return programCounter, nil
}

func doPower(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)
	dataTypeC := utils.GetVariableDataType(c, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)
	dataTypeCString := enums.ByteToDataType(dataTypeC)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in power operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	if dataTypeC != enums.DATATYPE_BYTE_NULL && dataTypeC != dataTypeA {
		return 0, fmt.Errorf("data type mismatch in add operation storage: %s, %s", dataTypeAString, dataTypeCString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.IntData[c] = int(math.Pow(float64(aVars.IntData[a]), float64(bVars.IntData[b])))
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.FloatData[c] = math.Pow(aVars.FloatData[a], bVars.FloatData[b])
	default:
		return 0, fmt.Errorf("power operation is not yet implemented for type %s", dataTypeAString)
	}

	programCounter += 1

	return programCounter, nil
}

func doBinaryAdd(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	// a := string(instruction.Args[0])
	// b := string(instruction.Args[1])
	// c := string(instruction.Args[2])

	// dataTypeA := utils.GetVariableDataType(a, vars)
	// dataTypeB := utils.GetVariableDataType(b, vars)
	// dataTypeC := utils.GetVariableDataType(c, vars)

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

func doBinarySubtract(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	// a := string(instruction.Args[0])
	// b := string(instruction.Args[1])
	// c := string(instruction.Args[2])

	// dataTypeA := utils.GetVariableDataType(a, vars)
	// dataTypeB := utils.GetVariableDataType(b, vars)
	// dataTypeC := utils.GetVariableDataType(c, vars)

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

func doGreater(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != enums.DATATYPE_BYTE_INT && dataTypeA != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in add operation: %s", dataTypeAString)
	}
	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in power operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.BoolData[c] = aVars.IntData[a] > bVars.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.BoolData[c] = aVars.FloatData[a] > bVars.FloatData[b]
	}

	programCounter += 1

	return programCounter, nil
}

func doGreaterEqual(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != enums.DATATYPE_BYTE_INT && dataTypeA != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in add operation: %s", dataTypeAString)
	}
	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in power operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.BoolData[c] = aVars.IntData[a] >= bVars.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.BoolData[c] = aVars.FloatData[a] >= bVars.FloatData[b]
	}

	programCounter += 1

	return programCounter, nil
}

func doEqual(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in equal operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		cVars.BoolData[c] = aVars.BoolData[a] == bVars.BoolData[b]
	case enums.DATATYPE_BYTE_INT:
		cVars.BoolData[c] = aVars.IntData[a] == bVars.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.BoolData[c] = aVars.FloatData[a] == bVars.FloatData[b]
	case enums.DATATYPE_BYTE_STRING:
		cVars.BoolData[c] = aVars.StringData[a] == bVars.StringData[b]
	default:
		return 0, fmt.Errorf("equal operation is not yet implemented for type %s", dataTypeAString)
	}

	programCounter += 1

	return programCounter, nil
}

func doLessEqual(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != enums.DATATYPE_BYTE_INT && dataTypeA != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in less equal operation: %s", dataTypeAString)
	}
	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in power operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.BoolData[c] = aVars.IntData[a] <= bVars.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.BoolData[c] = aVars.FloatData[a] <= bVars.FloatData[b]
	}

	programCounter += 1

	return programCounter, nil
}

func doLess(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])
	c := string(instruction.Args[2])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeB := utils.GetVariableDataType(b, vars)

	dataTypeAString := enums.ByteToDataType(dataTypeA)
	dataTypeBString := enums.ByteToDataType(dataTypeB)

	if dataTypeA != enums.DATATYPE_BYTE_INT && dataTypeA != enums.DATATYPE_BYTE_FLOAT {
		return 0, fmt.Errorf("data type invalid in less operation: %s", dataTypeAString)
	}
	if dataTypeA != dataTypeB {
		return 0, fmt.Errorf("data type mismatch in power operation: %s, %s", dataTypeAString, dataTypeBString)
	}

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)
	cVars := utils.GetVariableContext(c, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		cVars.BoolData[c] = aVars.IntData[a] < bVars.IntData[b]
	case enums.DATATYPE_BYTE_FLOAT:
		cVars.BoolData[c] = aVars.FloatData[a] < bVars.FloatData[b]
	}

	programCounter += 1

	return programCounter, nil
}

func doInputBlock(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])

	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')

	// convert CRLF to LF
	inputString = strings.Replace(inputString, "\n", "", -1)

	aVars := utils.GetVariableContext(a, enums.DATATYPE_BYTE_STRING, vars)

	aVars.StringData[a] = inputString

	programCounter += 1

	return programCounter, nil
}

func doInputNonBlock(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	// a := string(instruction.Args[0])
	// b := string(instruction.Args[1])
	// c := string(instruction.Args[2])

	// dataTypeA := utils.GetVariableDataType(a, vars)
	// dataTypeB := utils.GetVariableDataType(b, vars)

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

func doOutput(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])

	dataTypeA := utils.GetVariableDataType(a, vars)

	aVars := utils.GetVariableContext(a, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		fmt.Printf("%v", aVars.BoolData[a])
	case enums.DATATYPE_BYTE_INT:
		fmt.Printf("%v", aVars.IntData[a])
	case enums.DATATYPE_BYTE_FLOAT:
		fmt.Printf("%v", aVars.FloatData[a])
	case enums.DATATYPE_BYTE_STRING:
		fmt.Printf("%v", aVars.StringData[a])
	case enums.DATATYPE_BYTE_LIST:
		fmt.Printf("%v", aVars.ListData[a])
	case enums.DATATYPE_BYTE_DICT:
		fmt.Printf("%v", aVars.DictData[a])
	}

	programCounter += 1

	return programCounter, nil
}

func doBranchPositive(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeAString := enums.ByteToDataType(dataTypeA)

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		if aVars.BoolData[a] {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_INT:
		if aVars.IntData[a] > 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if aVars.FloatData[a] > 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch positive operation is not defined for type %s", dataTypeAString)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchNotPositive(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeAString := enums.ByteToDataType(dataTypeA)

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		if !aVars.BoolData[a] {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_INT:
		if aVars.IntData[a] <= 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if aVars.FloatData[a] <= 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch not positive operation is not defined for type %s", dataTypeAString)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchZero(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeAString := enums.ByteToDataType(dataTypeA)

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		if !aVars.BoolData[a] {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_INT:
		if aVars.IntData[a] == 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if aVars.FloatData[a] == 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch zero operation is not defined for type %s", dataTypeAString)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchNotZero(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeAString := enums.ByteToDataType(dataTypeA)

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_BOOL:
		if aVars.BoolData[a] {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_INT:
		if aVars.IntData[a] != 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if aVars.FloatData[a] != 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch not zero operation is not defined for type %s", dataTypeAString)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchNegative(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeAString := enums.ByteToDataType(dataTypeA)

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		if aVars.IntData[a] < 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if aVars.FloatData[a] < 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch negative operation is not defined for type %s", dataTypeAString)
	}

	programCounter += 1
	return programCounter, nil

}

func doBranchNotNegative(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])
	b := string(instruction.Args[1])

	dataTypeA := utils.GetVariableDataType(a, vars)
	dataTypeAString := enums.ByteToDataType(dataTypeA)

	aVars := utils.GetVariableContext(a, dataTypeA, vars)
	bVars := utils.GetVariableContext(b, dataTypeA, vars)

	switch dataTypeA {
	case enums.DATATYPE_BYTE_INT:
		if aVars.IntData[a] >= 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	case enums.DATATYPE_BYTE_FLOAT:
		if aVars.FloatData[a] >= 0 {
			programCounter = bVars.LabelData[b]
			return programCounter, nil
		}
	default:
		return 0, fmt.Errorf("branch not negative operation is not defined for type %s", dataTypeAString)
	}

	programCounter += 1
	return programCounter, nil

}

func doGoto(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])

	aVars := utils.GetVariableContext(a, enums.DATATYPE_BYTE_LABEL, vars)

	programCounter = aVars.LabelData[a]
	return programCounter, nil
}

func doStop(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	a := string(instruction.Args[0])

	aVars := utils.GetVariableContext(a, enums.DATATYPE_BYTE_INT, vars)

	returnCode := aVars.IntData[a]

	os.Exit(returnCode)

	return 0, nil
}

func doCopy(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	from := string(instruction.Args[0])
	to := string(instruction.Args[1])

	fromType, err := utils.FindVariableType(from, vars)
	if err != nil {
		return 0, err
	}
	toType, err := utils.FindVariableType(to, vars)
	if err != nil {
		toType = fromType
	}

	fromContext := utils.GetVariableContext(from, fromType, vars)
	toContext := utils.GetVariableContext(to, toType, vars)

	switch fromType {
	case enums.DATATYPE_BYTE_BOOL:
		toContext.BoolData[to] = fromContext.BoolData[from]
	case enums.DATATYPE_BYTE_INT:
		toContext.IntData[to] = fromContext.IntData[from]
	case enums.DATATYPE_BYTE_FLOAT:
		toContext.FloatData[to] = fromContext.FloatData[from]
	case enums.DATATYPE_BYTE_STRING:
		toContext.StringData[to] = fromContext.StringData[from]
	case enums.DATATYPE_BYTE_DICT:
		toContext.DictData[to] = fromContext.DictData[from]
	case enums.DATATYPE_BYTE_LIST:
		toContext.ListData[to] = fromContext.ListData[from]
	}

	programCounter += 1
	return programCounter, nil
}

func doDictAssign(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	from := string(instruction.Args[0])
	key := string(instruction.Args[1])
	to := string(instruction.Args[2])

	keyNameContext := utils.GetVariableContext(key, enums.DATATYPE_BYTE_STRING, vars)
	key = keyNameContext.StringData[key]

	fromType, err := utils.FindVariableType(from, vars)
	if err != nil {
		return 0, err
	}
	fromContext := utils.GetVariableContext(from, fromType, vars)
	toContext := utils.GetVariableContext(to, enums.DATATYPE_BYTE_DICT, vars)

	switch fromType {
	case enums.DATATYPE_BYTE_BOOL:
		toContext.DictData[to].BoolData[key] = fromContext.BoolData[from]
	case enums.DATATYPE_BYTE_INT:
		toContext.DictData[to].IntData[key] = fromContext.IntData[from]
	case enums.DATATYPE_BYTE_FLOAT:
		toContext.DictData[to].FloatData[key] = fromContext.FloatData[from]
	case enums.DATATYPE_BYTE_STRING:
		toContext.DictData[to].StringData[key] = fromContext.StringData[from]
	case enums.DATATYPE_BYTE_DICT:
		toContext.DictData[to].DictData[key] = fromContext.DictData[from]
	case enums.DATATYPE_BYTE_LIST:
		toContext.DictData[to].ListData[key] = fromContext.ListData[from]
	}

	programCounter += 1
	return programCounter, nil
}

func doDictAccess(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	from := string(instruction.Args[0])
	key := string(instruction.Args[1])
	to := string(instruction.Args[2])

	keyNameContext := utils.GetVariableContext(key, enums.DATATYPE_BYTE_STRING, vars)
	key = keyNameContext.StringData[key]

	fromContext := utils.GetVariableContext(from, enums.DATATYPE_BYTE_DICT, vars)
	fromType := utils.FindDictVariableType(key, fromContext.DictData[from])
	toContext := utils.GetVariableContext(to, fromType, vars)

	switch fromType {
	case enums.DATATYPE_BYTE_BOOL:
		toContext.BoolData[to] = fromContext.DictData[from].BoolData[key]
	case enums.DATATYPE_BYTE_INT:
		toContext.IntData[to] = fromContext.DictData[from].IntData[key]
	case enums.DATATYPE_BYTE_FLOAT:
		toContext.FloatData[to] = fromContext.DictData[from].FloatData[key]
	case enums.DATATYPE_BYTE_STRING:
		toContext.StringData[to] = fromContext.DictData[from].StringData[key]
	case enums.DATATYPE_BYTE_DICT:
		toContext.DictData[to] = fromContext.DictData[from].DictData[key]
	case enums.DATATYPE_BYTE_LIST:
		toContext.ListData[to] = fromContext.DictData[from].ListData[key]
	}

	programCounter += 1
	return programCounter, nil
}

func doListAssign(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	from := string(instruction.Args[0])
	keyName := string(instruction.Args[1])
	to := string(instruction.Args[2])

	keyNameContext := utils.GetVariableContext(keyName, enums.DATATYPE_BYTE_STRING, vars)
	key := keyNameContext.IntData[keyName]

	fromType, err := utils.FindVariableType(from, vars)
	if err != nil {
		return 0, err
	}
	fromContext := utils.GetVariableContext(from, fromType, vars)
	toContext := utils.GetVariableContext(to, enums.DATATYPE_BYTE_LIST, vars)

	switch fromType {
	case enums.DATATYPE_BYTE_BOOL:
		toContext.ListData[to].BoolData[key] = fromContext.BoolData[from]
	case enums.DATATYPE_BYTE_INT:
		toContext.ListData[to].IntData[key] = fromContext.IntData[from]
	case enums.DATATYPE_BYTE_FLOAT:
		toContext.ListData[to].FloatData[key] = fromContext.FloatData[from]
	case enums.DATATYPE_BYTE_STRING:
		toContext.ListData[to].StringData[key] = fromContext.StringData[from]
	case enums.DATATYPE_BYTE_DICT:
		toContext.ListData[to].DictData[key] = fromContext.DictData[from]
	case enums.DATATYPE_BYTE_LIST:
		toContext.ListData[to].ListData[key] = fromContext.ListData[from]
	}

	programCounter += 1
	return programCounter, nil
}

func doListAccess(instruction parser.Instruction, programCounter int, vars *memstore.VariableStore) (int, error) {
	from := string(instruction.Args[0])
	keyName := string(instruction.Args[1])
	to := string(instruction.Args[2])

	keyNameContext := utils.GetVariableContext(keyName, enums.DATATYPE_BYTE_STRING, vars)
	key := keyNameContext.IntData[keyName]

	fromContext := utils.GetVariableContext(from, enums.DATATYPE_BYTE_LIST, vars)
	fromType := utils.FindListVariableType(key, fromContext.ListData[from])
	toContext := utils.GetVariableContext(to, fromType, vars)

	switch fromType {
	case enums.DATATYPE_BYTE_BOOL:
		toContext.BoolData[to] = fromContext.ListData[from].BoolData[key]
	case enums.DATATYPE_BYTE_INT:
		toContext.IntData[to] = fromContext.ListData[from].IntData[key]
	case enums.DATATYPE_BYTE_FLOAT:
		toContext.FloatData[to] = fromContext.ListData[from].FloatData[key]
	case enums.DATATYPE_BYTE_STRING:
		toContext.StringData[to] = fromContext.ListData[from].StringData[key]
	case enums.DATATYPE_BYTE_DICT:
		toContext.DictData[to] = fromContext.ListData[from].DictData[key]
	case enums.DATATYPE_BYTE_LIST:
		toContext.ListData[to] = fromContext.ListData[from].ListData[key]
	}

	programCounter += 1
	return programCounter, nil
}
