package fileio

import (
	"bufio"
	"encoding/binary"
	"os"
)

func ReadFile(path string) ([]string, error) {
	lines := []string{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func WriteFile(path string, data []byte) error {
	outFile, err := os.Create(path)
	if err != nil {
		return err
	}
	for _, datum := range data {
		err := binary.Write(outFile, binary.BigEndian, datum)
		if err != nil {
			return err
		}
	}
	outFile.Close()
	return nil
}
