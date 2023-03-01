package utils

import "strconv"

func IntToByteArray(val int, length int) []byte {
	bytes := []byte(strconv.Itoa(val))

	for len(bytes) < length {
		bytes = append([]byte{byte(0)}, bytes...)
	}

	return bytes

}
