package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

const (
	// Plantuml character list
	plantumlMap = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
)

// Encodes a string using the DEFLATE algorithm
func deflateEncode(text string) []byte {
	var zbytes bytes.Buffer

	w := zlib.NewWriter(&zbytes)
	w.Write([]byte(text))
	w.Close()

	return zbytes.Bytes()
}

// Encodes a string using the Base64 algoritm, using plantumlMap character list
func base64Encode(text []byte) string {
	var base64buffer bytes.Buffer
	textLength := len(text)

	for i := 0; i < (3 - textLength%3); i++ {
		text = append(text, byte(0))
	}

	for i := 0; i < textLength; i += 3 {
		b1, b2, b3, b4 := text[i], text[i+1], text[i+2], byte(0)

		b4 = b3 & 0x3f
		b3 = ((b2 & 0xf) << 2) | (b3 >> 6)
		b2 = ((b1 & 0x3) << 4) | (b2 >> 4)
		b1 = b1 >> 2

		for _, b := range []byte{b1, b2, b3, b4} {
			base64buffer.WriteByte(byte(plantumlMap[b]))
		}
	}

	return string(base64buffer.Bytes())
}

func main() {
	// test
	fmt.Printf("~1%s\n", base64Encode(deflateEncode("@startuml\nB -> A : 111\n@enduml")))
}
