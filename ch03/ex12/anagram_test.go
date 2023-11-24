package main

import "testing"

var data = []struct {
	s1       string
	s2       string
	expected bool
}{
	{"", "", true},
	{"abcd", "cbda", true},
	{"abcd", "aaaa", false},
	{"あいうえお", "うえあおい", true},
	{"あいうえお", "あああああ", false},
}

func Test_anagram(t *testing.T) {
	for _, d := range data {
		result := anagram(d.s1, d.s2)
		if result != d.expected {
			t.Errorf("s1 %v, s2 %v, expected %v, result %v", d.s1, d.s2, d.expected, result)
		}
	}
}
