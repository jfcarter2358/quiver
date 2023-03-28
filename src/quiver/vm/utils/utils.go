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

func CoerceDict(data []byte, keyType, valType byte) (memstore.DictDataStore, error) {
	

	output := memstore.DictDataStore{}

	switch keyType {
	case enums.DATATYPE_BYTE_INT:
		data = val.(map[int]interface{})
		output.KeyType = enums.DATATYPE_BYTE_INT
		switch valType {
		case enums.DATATYPE_BYTE_BOOL:
			output.ValType = enums.DATATYPE_BYTE_BOOL

			var val map[bool]bool
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.IntBoolData = val.(map[int]bool)
		case enums.DATATYPE_BYTE_INT:
			output.ValType = enums.DATATYPE_BYTE_INT

			var val map[bool]int
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}


			output.IntIntData = val.(map[int]int)
		case enums.DATATYPE_BYTE_FLOAT:
			output.ValType = enums.DATATYPE_BYTE_FLOAT

			var val map[bool]float64
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}


			output.IntFloatData = val.(map[int]float64)
		case enums.DATATYPE_BYTE_STRING:
			output.ValType = enums.DATATYPE_BYTE_STRING

			var val map[bool]string
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}


			output.IntStringData = val.(map[int]string)
		case enums.DATATYPE_BYTE_DICT:
			output.ValType = enums.DATATYPE_BYTE_DICT
	
			var val map[bool]interface{}
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.IntDictData = val.(map[int]memstore.DictDataStore)
		case enums.DATATYPE_BYTE_LIST:
			ouptut.ValType = enums.DATATYPE_BYTE_LIST
	
			var val map[bool]interface{}
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.IntListData = val.(map[int]memstore.ListDataStore)
		}
	case enums.DATATYPE_BYTE_FLOAT:
		output.KeyType = enums.DATATYPE_BYTE_FLOAT
		switch valType {
		case enums.DATATYPE_BYTE_BOOL:
			output.ValType = enums.DATATYPE_BYTE_BOOL

			var val map[float64]bool
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.FloatBoolData = val.(map[float64]bool)
		case enums.DATATYPE_BYTE_INT:
			output.ValType = enums.DATATYPE_BYTE_INT
	
			var val map[float64]int
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.FloatIntData = val.(map[float64]int)
		case enums.DATATYPE_BYTE_FLOAT:
			output.ValType = enums.DATATYPE_BYTE_FLOAT

			var val map[float64]float64
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.FloatFloatData = val.(map[float64]float64)
		case enums.DATATYPE_BYTE_STRING:
			output.ValType = enums.DATATYPE_BYTE_STRING
	
			var val map[float64]string
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.FloatStringData = val.(map[float64]string)
		case enums.DATATYPE_BYTE_DICT:
			output.ValType = enums.DATATYPE_BYTE_DICT
	
			var val map[float64]interface{}
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.FloatDictData = val.(map[float64]memstore.DictDataStore)
		case enums.DATATYPE_BYTE_LIST:
			ouptut.ValType = enums.DATATYPE_BYTE_LIST
	
			var val map[float64]interface{}
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.FloatListData = val.(map[float64]memstore.ListDataStore)
		}
	case enums.DATATYPE_BYTE_STRING:
		output.KeyType = enums.DATATYPE_BYTE_STRING
		data = val.(map[string]interface{})
		switch valType {
		case enums.DATATYPE_BYTE_BOOL:
			output.ValType = enums.DATATYPE_BYTE_BOOL
	
			var val map[string]bool
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.StringBoolData = val.(map[string]bool)
		case enums.DATATYPE_BYTE_INT:
			output.ValType = enums.DATATYPE_BYTE_INT

			var val map[string]int
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.StringIntData = val.(map[string]int)
		case enums.DATATYPE_BYTE_FLOAT:
			output.ValType = enums.DATATYPE_BYTE_FLOAT
	
			var val map[string]float64
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.StringFloatData = val.(map[string]float64)
		case enums.DATATYPE_BYTE_STRING:
			output.ValType = enums.DATATYPE_BYTE_STRING
	
			var val map[string]string
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.StringStringData = val.(map[string]string)
		case enums.DATATYPE_BYTE_DICT:
			output.ValType = enums.DATATYPE_BYTE_DICT

			var val map[string]interface{}
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.StringDictData = val.(map[string]memstore.DictDataStore)
		case enums.DATATYPE_BYTE_LIST:
			ouptut.ValType = enums.DATATYPE_BYTE_LIST
	
			var val map[string}]interface{}
			err := json.Unmarshal(data, &val)

			if err != nil {
				return memstore.DictDataStore{}, err
			}

			output.StringListData = val.(map[string]memstore.ListDataStore)
		}
	}
	return output, nil
}

func CoerceList(data []byte, valType byte) (memstore.ListDataStore, error) {
	var val map[interface{}]interface{}
	err := json.Unmarshal(data, &val)
	if err != nil {
		return memstore.ListDataStore{}, err
	}

	output := memstore.ListDataStore{}

	switch valType {
	case enums.DATATYPE_BYTE_BOOL:
		output.ValType = enums.DATATYPE_BYTE_BOOL
		output.BoolData = val.([]bool)
	case enums.DATATYPE_BYTE_INT:
		output.ValType = enums.DATATYPE_BYTE_INT
		output.IntData = val.([]int)
	case enums.DATATYPE_BYTE_FLOAT:
		output.ValType = enums.DATATYPE_BYTE_FLOAT
		output.FloatData = val.([]float64)
	case enums.DATATYPE_BYTE_STRING:
		output.ValType = enums.DATATYPE_BYTE_STRING
		output.StringData = val.([]string)
	case enums.DATATYPE_BYTE_DICT:
		output.ValType = enums.DATATYPE_BYTE_DICT
		output.DictData = val.([]memstore.DictDataStore)
	case enums.DATATYPE_BYTE_LIST:
		ouptut.ValType = enums.DATATYPE_BYTE_LIST
		output.ListData = val.([]memstore.ListDataStore)
	}
	return output, nil
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
