package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	Control = "Control"
	Digit   = "Digit"
	Graphic = "Graphic"
	Letter  = "Letter"
	Lower   = "Lower"
	Mark    = "Mark"
	Number  = "Number"
	Print   = "Print"
	Punct   = "Punct"
	Space   = "Space"
	Symbol  = "Symbol"
	Title   = "Title"
	Upper   = "Upper"
)

var functions = map[string]func(rune) bool{
	Control: unicode.IsControl,
	Digit:   unicode.IsDigit,
	Graphic: unicode.IsGraphic,
	Letter:  unicode.IsLetter,
	Lower:   unicode.IsLower,
	Mark:    unicode.IsMark,
	Number:  unicode.IsNumber,
	Print:   unicode.IsPrint,
	Punct:   unicode.IsPunct,
	Space:   unicode.IsSpace,
	Symbol:  unicode.IsSymbol,
	Title:   unicode.IsTitle,
	Upper:   unicode.IsUpper,
}

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	categories := make(map[string]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		for c, f := range functions {
			if f(r) {
				categories[c]++
			}
		}
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Print("\ncategory\tcount\n")
	for c, n := range categories {
		fmt.Printf("%s\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
