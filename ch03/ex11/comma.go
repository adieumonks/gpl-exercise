package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	if len(s) == 0 {
		return s
	}

	if s[0:1] == "+" || s[0:1] == "-" {
		return s[0:1] + comma(s[1:])
	}

	pIndex := strings.LastIndex(s, ".")
	if pIndex >= 0 {
		return comma(s[:pIndex]) + s[pIndex:]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma("123456789"))
	fmt.Println(comma("+123456789"))
	fmt.Println(comma("-123456789"))
	fmt.Println(comma("123456789.00000"))
	fmt.Println(comma("+123456789.00000"))
	fmt.Println(comma("-123456789.00000"))
}
