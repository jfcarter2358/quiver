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
	BoolData   map[int]bool
	IntData    map[int]int
	FloatData  map[int]float64
	StringData map[int]string
	DictData   map[int]DictDataStore
	ListData   map[int]ListDataStore
}

type DictDataStore struct {
	BoolData   map[string]bool
	IntData    map[string]int
	FloatData  map[string]float64
	StringData map[string]string
	DictData   map[string]DictDataStore
	ListData   map[string]ListDataStore
}

var Vars VariableStore

func (d *DictDataStore) Init() {
	d.BoolData = map[string]bool{}
	d.IntData = map[string]int{}
	d.FloatData = map[string]float64{}
	d.StringData = map[string]string{}
	d.DictData = map[string]DictDataStore{}
	d.ListData = map[string]ListDataStore{}
}

func (l *ListDataStore) Init() {
	l.BoolData = map[int]bool{}
	l.IntData = map[int]int{}
	l.FloatData = map[int]float64{}
	l.StringData = map[int]string{}
	l.DictData = map[int]DictDataStore{}
	l.ListData = map[int]ListDataStore{}
}

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
