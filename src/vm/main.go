package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

func read_qvc(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err

	}
	defer file.Close()

	stats, err := file.Stat()

	if err != nil {
		return nil, err
	}

	var size int64 = stats.Size()

	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)

	_, err = buffer.Read(bytes)

	return bytes, err
}

func main() {
	args := os.Args[1:]

	bytecode, err := read_qvc(args[0])

	if err != nil {
		panic(err)
	}

	dataLength := binary.BigEndian.Uint64(bytecode[:8])
	bytecode = bytecode[8:]

	fmt.Printf("%d", dataLength)
}
