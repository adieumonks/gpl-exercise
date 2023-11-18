package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestFetchAll(t *testing.T) {
	buffer := &bytes.Buffer{}
	data, _ := os.ReadFile("alexa.txt")
	sites := strings.Split(string(data), "\n")
	fetchAll(sites, buffer)
	fmt.Println(buffer)
}
