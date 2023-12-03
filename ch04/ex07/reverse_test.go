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

func TestReverse(t *testing.T) {
	data := []struct {
		s        []byte
		expected []byte
	}{
		{
			[]byte{},
			[]byte{},
		},
		{
			[]byte("a"),
			[]byte("a"),
		},
		{
			[]byte("ab"),
			[]byte("ba"),
		},
		{
			[]byte("abc"),
			[]byte("cba"),
		},
		{
			[]byte("こんにちは 世界"),
			[]byte("界世 はちにんこ"),
		},
		{
			[]byte("Hello World こんにちは 世界"),
			[]byte("界世 はちにんこ dlroW olleH"),
		},
	}

	for _, d := range data {
		Reverse(d.s)
		if !equal(d.s, d.expected) {
			t.Errorf("result: %v, expected: %v", d.s, d.expected)
		}
	}
}
