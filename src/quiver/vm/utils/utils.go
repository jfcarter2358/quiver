package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
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

func CoerceDict(byteData []byte) (memstore.DictDataStore, error) {
	output := memstore.DictDataStore{}
	trimmed := bytes.Trim(byteData, "\x00")

	var data map[string]interface{}
	err := json.Unmarshal(trimmed, &data)

	if err != nil {
		return memstore.DictDataStore{}, err
	}

	output = CoerceDictInterface(data)

	return output, nil
}

func CoerceList(byteData []byte) (memstore.ListDataStore, error) {
	output := memstore.ListDataStore{}
	trimmed := bytes.Trim(byteData, "\x00")

	var data []interface{}
	err := json.Unmarshal(trimmed, &data)

	if err != nil {
		return memstore.ListDataStore{}, err
	}

	output = CoerceListInterface(data)

	return output, nil
}
func CoerceDictInterface(data map[string]interface{}) memstore.DictDataStore {
	output := memstore.DictDataStore{}
	output.Init()

	for key, val := range data {
		valType := GetDataTypeForInterface(val)

		switch valType {
		case enums.DATATYPE_BYTE_BOOL:
			output.BoolData[key] = val.(bool)
		case enums.DATATYPE_BYTE_INT:
			output.IntData[key] = val.(int)
		case enums.DATATYPE_BYTE_FLOAT:
			output.FloatData[key] = val.(float64)
		case enums.DATATYPE_BYTE_STRING:
			output.StringData[key] = val.(string)
		case enums.DATATYPE_BYTE_DICT:
			output.DictData[key] = CoerceDictInterface(val.(map[string]interface{}))
		case enums.DATATYPE_BYTE_LIST:
			output.ListData[key] = CoerceListInterface(val.([]interface{}))
		}
	}

	return output
}

func CoerceListInterface(data []interface{}) memstore.ListDataStore {
	output := memstore.ListDataStore{}
	output.Init()

	for idx, val := range data {
		valType := GetDataTypeForInterface(val)

		switch valType {
		case enums.DATATYPE_BYTE_BOOL:
			output.BoolData[idx] = val.(bool)
		case enums.DATATYPE_BYTE_INT:
			output.IntData[idx] = val.(int)
		case enums.DATATYPE_BYTE_FLOAT:
			output.FloatData[idx] = val.(float64)
		case enums.DATATYPE_BYTE_STRING:
			output.StringData[idx] = val.(string)
		case enums.DATATYPE_BYTE_DICT:
			output.DictData[idx] = CoerceDictInterface(val.(map[string]interface{}))
		case enums.DATATYPE_BYTE_LIST:
			coercedVal, _ := val.([]interface{})
			output.ListData[idx] = CoerceListInterface(coercedVal)
		}
	}

	return output
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

func GetDataTypeForInterface(val interface{}) byte {
	if _, ok := val.(string); ok {
		return enums.DATATYPE_BYTE_STRING
	}

	if _, ok := val.(float64); ok {
		return enums.DATATYPE_BYTE_FLOAT
	}

	if _, ok := val.(int); ok {
		return enums.DATATYPE_BYTE_INT
	}

	if _, ok := val.(bool); ok {
		return enums.DATATYPE_BYTE_BOOL
	}

	if _, ok := val.([]interface{}); ok {
		return enums.DATATYPE_BYTE_LIST
	}

	return enums.DATATYPE_BYTE_DICT
}

func FindVariableType(name string, vars *memstore.VariableStore) (byte, error) {
	tempStore := vars

	for tempStore != nil {
		if _, ok := vars.BoolData[name]; ok {
			return enums.DATATYPE_BYTE_BOOL, nil
		}
		if _, ok := vars.IntData[name]; ok {
			return enums.DATATYPE_BYTE_INT, nil
		}
		if _, ok := vars.FloatData[name]; ok {
			return enums.DATATYPE_BYTE_FLOAT, nil
		}
		if _, ok := vars.StringData[name]; ok {
			return enums.DATATYPE_BYTE_STRING, nil
		}
		if _, ok := vars.DictData[name]; ok {
			return enums.DATATYPE_BYTE_DICT, nil
		}
		if _, ok := vars.ListData[name]; ok {
			return enums.DATATYPE_BYTE_LIST, nil
		}
		tempStore = tempStore.Parent
	}

	return byte(0), fmt.Errorf("variable %s does not exist", name)
}

func FindDictVariableType(name string, dict memstore.DictDataStore) byte {
	if _, ok := dict.BoolData[name]; ok {
		return enums.DATATYPE_BYTE_BOOL
	}
	if _, ok := dict.IntData[name]; ok {
		return enums.DATATYPE_BYTE_INT
	}
	if _, ok := dict.FloatData[name]; ok {
		return enums.DATATYPE_BYTE_FLOAT
	}
	if _, ok := dict.StringData[name]; ok {
		return enums.DATATYPE_BYTE_STRING
	}
	if _, ok := dict.DictData[name]; ok {
		return enums.DATATYPE_BYTE_DICT
	}
	return enums.DATATYPE_BYTE_LIST
}

func FindListVariableType(idx int, list memstore.ListDataStore) byte {
	if _, ok := list.BoolData[idx]; ok {
		return enums.DATATYPE_BYTE_BOOL
	}
	if _, ok := list.IntData[idx]; ok {
		return enums.DATATYPE_BYTE_INT
	}
	if _, ok := list.FloatData[idx]; ok {
		return enums.DATATYPE_BYTE_FLOAT
	}
	if _, ok := list.StringData[idx]; ok {
		return enums.DATATYPE_BYTE_STRING
	}
	if _, ok := list.DictData[idx]; ok {
		return enums.DATATYPE_BYTE_DICT
	}
	return enums.DATATYPE_BYTE_LIST
}
