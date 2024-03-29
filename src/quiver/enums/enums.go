package enums

import "fmt"

/* -------- DOT CODES -------- */

const DOT_CODE_NAME_LITERAL = ".LITERAL"
const DOT_CODE_NAME_LABEL = ".LABEL"

const DOT_CODE_BYTE_LITERAL = byte(0)
const DOT_CODE_BYTE_LABEL = byte(1)

/* -------- DATA TYPES -------- */

const DATATYPE_NAME_NULL = "null"
const DATATYPE_NAME_BOOL = "bool"
const DATATYPE_NAME_INT = "int"
const DATATYPE_NAME_FLOAT = "float"
const DATATYPE_NAME_STRING = "string"
const DATATYPE_NAME_LIST = "list"
const DATATYPE_NAME_DICT = "dict"
const DATATYPE_NAME_LABEL = "label"

const DATATYPE_BYTE_NULL = byte(0)
const DATATYPE_BYTE_BOOL = byte(1)
const DATATYPE_BYTE_INT = byte(2)
const DATATYPE_BYTE_FLOAT = byte(3)
const DATATYPE_BYTE_STRING = byte(4)
const DATATYPE_BYTE_LIST = byte(5)
const DATATYPE_BYTE_DICT = byte(6)
const DATATYPE_BYTE_LABEL = byte(7)

/* -------- OP CODES -------- */

const OP_CODE_NAME_ADD = "ADD"
const OP_CODE_NAME_SUBTRACT = "SUBTRACT"
const OP_CODE_NAME_MULTIPLY = "MULTIPLY"
const OP_CODE_NAME_DIVIDE = "DIVIDE"
const OP_CODE_NAME_MODULO = "MODULO"
const OP_CODE_NAME_POWER = "POWER"
const OP_CODE_NAME_BINARY_ADD = "BINARY_ADD"
const OP_CODE_NAME_BINARY_SUBTRACT = "BINARY_SUBTRACT"
const OP_CODE_NAME_GREATER = "GREATER"
const OP_CODE_NAME_GREATER_EQUAL = "GREATER_EQUAL"
const OP_CODE_NAME_EQUAL = "EQUAL"
const OP_CODE_NAME_LESS_EQUAL = "LESS_EQUAL"
const OP_CODE_NAME_LESS = "LESS"
const OP_CODE_NAME_INPUT_BLOCK = "INPUT_BLOCK"
const OP_CODE_NAME_INPUT_NON_BLOCK = "INPUT_NON_BLOCK"
const OP_CODE_NAME_OUTPUT = "OUTPUT"
const OP_CODE_NAME_BRANCH_POSITIVE = "BRANCH_POSITIVE"
const OP_CODE_NAME_BRANCH_NOT_POSITIVE = "BRANCH_NOT_POSITIVE"
const OP_CODE_NAME_BRANCH_ZERO = "BRANCH_ZERO"
const OP_CODE_NAME_BRANCH_NOT_ZERO = "BRANCH_NOT_ZERO"
const OP_CODE_NAME_BRANCH_NEGATIVE = "BRANCH_NEGATIVE"
const OP_CODE_NAME_BRANCH_NOT_NEGATIVE = "BRANCH_NOT_NEGATIVE"
const OP_CODE_NAME_GOTO = "GOTO"
const OP_CODE_NAME_STOP = "STOP"
const OP_CODE_NAME_COPY = "COPY"
const OP_CODE_NAME_DICT_ACCESS = "DICT_ACCESS"
const OP_CODE_NAME_DICT_ASSIGN = "DICT_ASSIGN"
const OP_CODE_NAME_LIST_ACCESS = "LIST_ACCESS"
const OP_CODE_NAME_LIST_ASSIGN = "LIST_ASSIGN"

const OP_CODE_BYTE_ADD = byte(0)
const OP_CODE_BYTE_SUBTRACT = byte(1)
const OP_CODE_BYTE_MULTIPLY = byte(2)
const OP_CODE_BYTE_DIVIDE = byte(3)
const OP_CODE_BYTE_MODULO = byte(4)
const OP_CODE_BYTE_POWER = byte(5)
const OP_CODE_BYTE_BINARY_ADD = byte(6)
const OP_CODE_BYTE_BINARY_SUBTRACT = byte(7)
const OP_CODE_BYTE_GREATER = byte(8)
const OP_CODE_BYTE_GREATER_EQUAL = byte(9)
const OP_CODE_BYTE_EQUAL = byte(10)
const OP_CODE_BYTE_LESS_EQUAL = byte(11)
const OP_CODE_BYTE_LESS = byte(12)
const OP_CODE_BYTE_INPUT_BLOCK = byte(13)
const OP_CODE_BYTE_INPUT_NON_BLOCK = byte(14)
const OP_CODE_BYTE_OUTPUT = byte(15)
const OP_CODE_BYTE_BRANCH_POSITIVE = byte(16)
const OP_CODE_BYTE_BRANCH_NOT_POSITIVE = byte(17)
const OP_CODE_BYTE_BRANCH_ZERO = byte(18)
const OP_CODE_BYTE_BRANCH_NOT_ZERO = byte(19)
const OP_CODE_BYTE_BRANCH_NEGATIVE = byte(20)
const OP_CODE_BYTE_BRANCH_NOT_NEGATIVE = byte(21)
const OP_CODE_BYTE_GOTO = byte(22)
const OP_CODE_BYTE_STOP = byte(23)
const OP_CODE_BYTE_COPY = byte(24)
const OP_CODE_BYTE_DICT_ACCESS = byte(25)
const OP_CODE_BYTE_DICT_ASSIGN = byte(26)
const OP_CODE_BYTE_LIST_ACCESS = byte(27)
const OP_CODE_BYTE_LIST_ASSIGN = byte(28)

/* -------- CONVENIENCE FUNCTIONS -------- */

func ByteToDotCode(dotCodeByte byte) (string, error) {
	switch dotCodeByte {
	case DOT_CODE_BYTE_LITERAL:
		return DOT_CODE_NAME_LITERAL, nil
	case DOT_CODE_BYTE_LABEL:
		return DOT_CODE_NAME_LABEL, nil
	}
	return "", fmt.Errorf("unknown dot code %v", dotCodeByte)
}

func ByteToDataType(dataTypeCodeByte byte) string {
	switch dataTypeCodeByte {
	case DATATYPE_BYTE_BOOL:
		return DATATYPE_NAME_BOOL
	case DATATYPE_BYTE_INT:
		return DATATYPE_NAME_INT
	case DATATYPE_BYTE_FLOAT:
		return DATATYPE_NAME_FLOAT
	case DATATYPE_BYTE_STRING:
		return DATATYPE_NAME_STRING
	case DATATYPE_BYTE_LIST:
		return DATATYPE_NAME_LIST
	case DATATYPE_BYTE_DICT:
		return DATATYPE_NAME_DICT
	}
	return DATATYPE_NAME_NULL
}
