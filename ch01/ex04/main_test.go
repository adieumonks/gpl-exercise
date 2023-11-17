package main

import (
	"os"
	"strings"
	"testing"
)

func TestDup(t *testing.T) {
	counts, err := Dup([]string{"sample.txt"})
	if err != nil {
		t.Errorf("error is thrown %v", err)
	}
	for filename, count := range counts {
		if filename != "sample.txt" {
			t.Errorf("unexpected filename %s", filename)
		}
		for line, n := range count {
			if n < 1 {
				t.Errorf("unexpected linecount %d", n)
			}
			data, _ := os.ReadFile(filename)
			if !strings.Contains(string(data), line) {
				t.Errorf("unexpected line %s", line)
			}
		}
	}
}
