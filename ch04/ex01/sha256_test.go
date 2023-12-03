package main

import (
	"crypto/sha256"
	"testing"
)

var data = []struct {
	d1       []byte
	d2       []byte
	expected int
}{
	{[]byte("x"), []byte("x"), 0},
	{[]byte("X"), []byte("X"), 0},
	{[]byte("x"), []byte("X"), 125},
}

func Test_totalPopCount(t *testing.T) {
	for _, d := range data {
		sum1 := sha256.Sum256(d.d1)
		sum2 := sha256.Sum256(d.d2)

		xorBytes := make([]byte, 0, sha256.Size)
		for i := 0; i < sha256.Size; i++ {
			xorBytes = append(xorBytes, sum1[i]^sum2[i])
		}
		result := totalPopCount(xorBytes)
		if result != d.expected {
			t.Errorf("result: %d, expected: %d", result, d.expected)
		}
	}
}
