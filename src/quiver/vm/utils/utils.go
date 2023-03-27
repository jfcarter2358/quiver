package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"quiver/enums"
	"quiver/vm/memstore"
	"strconv"
	"strings"
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

func CoerceDict(data []byte) (map[interface{}]interface{}, error) {
	var val map[interface{}]interface{}
	err := json.Unmarshal(data, &val)
	return val, err
}

func CoerceList(data []byte) ([]interface{}, error) {
	var val []interface{}
	err := json.Unmarshal(data, &val)
	return val, err
}

func GetVariableDataType(name string, vars *memstore.VariableStore) byte {
	if _, ok := vars.BoolData[name]; ok {
		return enums.DATATYPE_BYTE_BOOL
	}
	if _, ok := vars.IntData[name]; ok {
		return enums.DATATYPE_BYTE_INT
	}
	if _, ok := vars.FloatData[name]; ok {
		return enums.DATATYPE_BYTE_FLOAT
	}
	if _, ok := vars.StringData[name]; ok {
		return enums.DATATYPE_BYTE_STRING
	}
	if _, ok := vars.ListData[name]; ok {
		return enums.DATATYPE_BYTE_LIST
	}
	return enums.DATATYPE_BYTE_NULL
}

func GetVariableContext(key string, dataType byte, vars *memstore.VariableStore) *memstore.VariableStore {
	baseStore := vars
	tempStore := vars

	switch dataType {
	case enums.DATATYPE_BYTE_BOOL:
		for tempStore != nil {
			if _, ok := tempStore.BoolData[key]; !ok {
				tempStore = tempStore.Parent
				continue
			}
			break
		}
		if tempStore == nil {
			return baseStore
		}
		return tempStore
	case enums.DATATYPE_BYTE_FLOAT:
		for tempStore != nil {
			if _, ok := tempStore.FloatData[key]; !ok {
				tempStore = tempStore.Parent
				continue
			}
			break
		}
		if tempStore == nil {
			return baseStore
		}
		return tempStore
	case enums.DATATYPE_BYTE_INT:
		for tempStore != nil {
			if _, ok := tempStore.IntData[key]; !ok {
				tempStore = tempStore.Parent
				continue
			}
			break
		}
		if tempStore == nil {
			return baseStore
		}
		return tempStore
	case enums.DATATYPE_BYTE_STRING:
		for tempStore != nil {
			if _, ok := tempStore.StringData[key]; !ok {
				tempStore = tempStore.Parent
				continue
			}
			break
		}
		if tempStore == nil {
			return baseStore
		}
		return tempStore
	case enums.DATATYPE_BYTE_DICT:
		for tempStore != nil {
			if _, ok := tempStore.DictData[key]; !ok {
				tempStore = tempStore.Parent
				continue
			}
			break
		}
		if tempStore == nil {
			return baseStore
		}
		return tempStore
	case enums.DATATYPE_BYTE_LIST:
		for tempStore != nil {
			if _, ok := tempStore.ListData[key]; !ok {
				tempStore = tempStore.Parent
				continue
			}
			break
		}
		if tempStore == nil {
			return baseStore
		}
		return tempStore
	case enums.DATATYPE_BYTE_LABEL:
		for tempStore != nil {
			if _, ok := tempStore.LabelData[key]; !ok {
				tempStore = tempStore.Parent
				continue
			}
			break
		}
		if tempStore == nil {
			return baseStore
		}
		return tempStore
	}
	return baseStore
}
