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
	ListData   map[string][]interface{}
	DictData   map[string]map[interface{}]interface{}
	FuncData   map[string]map[string]byte
	PostalData map[string]map[string]string
	SocketData map[string]net.Conn
	Parent     *VariableStore
}

var Vars VariableStore

func Init() {
	Vars = VariableStore{
		LabelData:  map[string]int{},
		BoolData:   map[string]bool{},
		IntData:    map[string]int{},
		FloatData:  map[string]float64{},
		StringData: map[string]string{},
		ListData:   map[string][]interface{}{},
		DictData:   map[string]map[interface{}]interface{}{},
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
