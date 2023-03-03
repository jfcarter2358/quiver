package utils

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"
	"vm/enums"
	"vm/memstore"
)

func CoerceByteInt(data []byte) int {
	var returnVal int

	switch len(data) {
	case 1:
		returnVal = int(data[0])
	case 2:
		returnVal = int(binary.BigEndian.Uint16(data))
	case 4:
		returnVal = int(binary.BigEndian.Uint32(data))
	case 8:
		returnVal = int(binary.BigEndian.Uint64(data))
	}

	return returnVal
}

func CoerceBool(data []byte) (bool, error) {
	stringRep := string(bytes.Trim(data, "\x00"))
	val, err := strconv.ParseBool(stringRep)
	return val, err
}

func CoerceInt(data []byte) (int, error) {
	stringRep := string(bytes.Trim(data, "\x00"))
	val, err := strconv.Atoi(stringRep)
	return int(val), err
}

func CoerceFloat(data []byte) (float64, error) {
	stringRep := string(bytes.Trim(data, "\x00"))
	val, err := strconv.ParseFloat(stringRep, 64)
	return val, err
}

func CoerceString(data []byte) string {
	stringRep := string(bytes.Trim(data, "\x00"))
	stringRep = strings.ReplaceAll(stringRep, "\\n", "\n")
	return strings.Trim(stringRep, "\"")
}

func GetVariableDataType(name string) byte {
	if _, ok := memstore.BoolData[name]; ok {
		return enums.DATATYPE_BYTE_BOOL
	}
	if _, ok := memstore.IntData[name]; ok {
		return enums.DATATYPE_BYTE_INT
	}
	if _, ok := memstore.FloatData[name]; ok {
		return enums.DATATYPE_BYTE_FLOAT
	}
	if _, ok := memstore.StringData[name]; ok {
		return enums.DATATYPE_BYTE_STRING
	}
	if _, ok := memstore.ListData[name]; ok {
		return enums.DATATYPE_BYTE_LIST
	}
	return enums.DATATYPE_BYTE_NULL
}
