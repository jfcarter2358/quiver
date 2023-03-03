package memstore

import "fmt"

var LabelData map[string]int
var BoolData map[string]bool
var IntData map[string]int
var FloatData map[string]float64
var StringData map[string]string
var ListData map[string][]interface{}
var DictData map[string]map[interface{}]interface{}

func Init() {
	LabelData = map[string]int{}
	BoolData = map[string]bool{}
	IntData = map[string]int{}
	FloatData = map[string]float64{}
	StringData = map[string]string{}
	ListData = map[string][]interface{}{}
	DictData = map[string]map[interface{}]interface{}{}
}

func PrintState() {
	fmt.Printf("Label data: %v\n", LabelData)
	fmt.Printf("Bool data: %v\n", BoolData)
	fmt.Printf("Int data: %v\n", IntData)
	fmt.Printf("Float data: %v\n", FloatData)
	fmt.Printf("String data: %v\n", StringData)
	fmt.Printf("List data: %v\n", ListData)
	fmt.Printf("Dict data: %v\n", DictData)
}
