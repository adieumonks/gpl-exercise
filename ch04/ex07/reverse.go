package main

import "unicode/utf8"

func reverse(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func Reverse(bytes []byte) {
	var size int
	for i := 0; i < len(bytes); i += size {
		_, size = utf8.DecodeRune(bytes[i:])
		reverse(bytes[i : i+size])
	}
	reverse(bytes)
}
