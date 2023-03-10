package memstore

import (
	"fmt"
	"net"
)

var LabelData map[string]int
var BoolData map[string]bool
var IntData map[string]int
var FloatData map[string]float64
var StringData map[string]string
var ListData map[string][]interface{}
var DictData map[string]map[interface{}]interface{}
var FuncData map[string]map[string]string
var PostalData map[string]map[string]string
var SocketData map[string]net.Conn

func Init() {
	LabelData = map[string]int{}
	BoolData = map[string]bool{}
	IntData = map[string]int{}
	FloatData = map[string]float64{}
	StringData = map[string]string{}
	ListData = map[string][]interface{}{}
	DictData = map[string]map[interface{}]interface{}{}
	FuncData = map[string]map[string]string{}
	PostalData = map[string]map[string]string{}
	SocketData = map[string]net.Conn{}
}

func PrintState() {
	fmt.Println()
	fmt.Printf("Label data: %v\n", LabelData)
	fmt.Printf("Bool data: %v\n", BoolData)
	fmt.Printf("Int data: %v\n", IntData)
	fmt.Printf("Float data: %v\n", FloatData)
	fmt.Printf("String data: %v\n", StringData)
	fmt.Printf("List data: %v\n", ListData)
	fmt.Printf("Dict data: %v\n", DictData)
	fmt.Printf("Func data: %v\n", FuncData)
	fmt.Printf("Postal data: %v\n", PostalData)
	fmt.Printf("Socket data: %v\n", SocketData)
}
