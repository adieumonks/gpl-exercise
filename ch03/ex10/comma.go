package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	bc := len(s) % 3
	if bc == 0 {
		bc = 3
	}
	for i := 0; i < len(s); i++ {
		if bc == 0 {
			buf.WriteByte(',')
			bc = 3
		}
		buf.WriteByte(s[i])
		bc--
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("123456789"))
}
