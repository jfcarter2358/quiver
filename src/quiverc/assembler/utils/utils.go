package utils

import (
	"encoding/binary"
	"fmt"
	"quiverc/enums"
	"strconv"
	"strings"
)

func IntToByteArray(val int, length int) []byte {
	switch length {
	case 1:
		return []byte{byte(val)}
	case 2:
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, uint16(val))
		return b
	case 4:
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(val))
		return b
	case 8:
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(val))
		return b
	}
	return []byte{}
}

func GetDataType(val string) (byte, error) {
	if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
		return enums.DATATYPE_BYTE_STRING, nil
	}
	if val == "true" || val == "false" {
		return enums.DATATYPE_BYTE_BOOL, nil
	}
	if _, err := strconv.ParseInt(val, 10, 64); err == nil {
		return enums.DATATYPE_BYTE_INT, nil
	}
	if _, err := strconv.ParseFloat(val, 64); err == nil {
		return enums.DATATYPE_BYTE_FLOAT, nil
	}

	return byte(0), fmt.Errorf("value '%s' does not match any known datatype", val)
}

func uint64ToLenBytes(val uint64, length int) []byte {
	byteArr := make([]byte, length)

	for i := 0; i < length; i++ {
		f := 8 * i
		byteArr[i] = byte(val >> f)
	}

	return byteArr
}

func int64ToLenBytes(val int64, length int) []byte {
	return uint64ToLenBytes(uint64(val), length)
}
