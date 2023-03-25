package fileio

import (
	"bufio"
	"os"
)

func ReadQVC(path string) ([]byte, error) {
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
