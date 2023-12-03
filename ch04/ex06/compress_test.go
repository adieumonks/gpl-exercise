package main

import (
	"testing"
)

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_compress(t *testing.T) {
	data := []struct {
		s        []byte
		expected []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte(" "), []byte(" ")},
		{[]byte("  "), []byte(" ")},
		{[]byte("   "), []byte(" ")},
		{[]byte(" \t\n\v\r\f\u0085\u00A0"), []byte(" ")},
	}

	for _, d := range data {
		result := compress(d.s)
		if !equal(result, d.expected) {
			t.Errorf("result: %v, expected: %v\n", result, d.expected)
		}
	}
}
