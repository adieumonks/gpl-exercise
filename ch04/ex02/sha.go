package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	shaType := flag.String("sha", "256", `"256" "384" "512"`)
	flag.Parse()

	switch *shaType {
	case "256", "384", "512":
	default:
		fmt.Printf("invalid sha: %s\n", *shaType)
		flag.PrintDefaults()
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)
	var t string
	if sc.Scan() {
		t = sc.Text()
	}
	switch *shaType {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(t)))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(t)))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(t)))
	}

}
