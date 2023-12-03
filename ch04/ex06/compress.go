package main

import (
	"unicode"
	"unicode/utf8"
)

func compress(bytes []byte) []byte {
	var result []byte
	var size int
	inSpace := false
	for i := 0; i < len(bytes); i += size {
		r, s := utf8.DecodeRune(bytes[i:])
		if !unicode.IsSpace(r) {
			result = append(result, bytes[i:i+s]...)
			inSpace = false
		} else {
			if !inSpace {
				result = append(result, byte(0x20))
				inSpace = true
			}
		}
		size = s
	}
	return result
}
