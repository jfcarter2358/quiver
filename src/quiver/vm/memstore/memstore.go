package memstore

import (
	"fmt"
	"net"
)

type VariableStore struct {
	LabelData  map[string]int
	BoolData   map[string]bool
	IntData    map[string]int
	FloatData  map[string]float64
	StringData map[string]string
	ListData   map[string]ListDataStore
	DictData   map[string]DictDataStore
	FuncData   map[string]map[string]byte
	PostalData map[string]map[string]string
	SocketData map[string]net.Conn
	Parent     *VariableStore
}

type ListDataStore struct {
	ValType    byte
	BoolData   []bool
	IntData    []int
	FloatData  []float64
	StringData []string
	DictData   []DictDataStore
	ListData   []ListDataStore
}

type DictDataStore struct {
	KeyType          byte
	ValType          byte
	IntBoolData      map[int]bool
	IntIntData       map[int]int
	IntFloatData     map[int]float64
	IntStringData    map[int]string
	IntDictData      map[int]DictDataStore
	IntListData      map[int]ListDataStore
	FloatBoolData    map[float64]bool
	FloatIntData     map[float64]int
	FloatFloatData   map[float64]float64
	FloatStringData  map[float64]string
	FloatDictData    map[float64]DictDataStore
	FloatListData    map[float64]ListDataStore
	StringBoolData   map[string]bool
	StringIntData    map[string]int
	StringFloatData  map[string]float64
	StringStringData map[string]string
	StringDictData   map[string]DictDataStore
	StringListData   map[string]ListDataStore
}

var Vars VariableStore

func Init() {
	Vars = VariableStore{
		LabelData:  map[string]int{},
		BoolData:   map[string]bool{},
		IntData:    map[string]int{},
		FloatData:  map[string]float64{},
		StringData: map[string]string{},
		ListData:   map[string]ListDataStore{},
		DictData:   map[string]DictDataStore{},
		FuncData:   map[string]map[string]byte{},
		PostalData: map[string]map[string]string{},
		SocketData: map[string]net.Conn{},
		Parent:     nil,
	}
}

func PrintState() {
	fmt.Println()
	fmt.Printf("Label data: %v\n", Vars.LabelData)
	fmt.Printf("Bool data: %v\n", Vars.BoolData)
	fmt.Printf("Int data: %v\n", Vars.IntData)
	fmt.Printf("Float data: %v\n", Vars.FloatData)
	fmt.Printf("String data: %v\n", Vars.StringData)
	fmt.Printf("List data: %v\n", Vars.ListData)
	fmt.Printf("Dict data: %v\n", Vars.DictData)
	fmt.Printf("Func data: %v\n", Vars.FuncData)
	fmt.Printf("Postal data: %v\n", Vars.PostalData)
	fmt.Printf("Socket data: %v\n", Vars.SocketData)
}
