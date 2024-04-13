package utils

import (
	"crypto/rand"
	"io"
)

func EncodeToString(max int, randType string) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}

	if randType == "int" {
		for i := 0; i < len(b); i++ {
			b[i] = numTable[int(b[i])%len(numTable)]
		}
	} else {
		for i := 0; i < len(b); i++ {
			b[i] = stringTable[int(b[i])%len(stringTable)]
		}
	}

	return string(b)
}

var numTable = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var stringTable = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
