package main

import "testing"

func Test_reverse(t *testing.T) {
	a := [6]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	expected := [6]int{5, 4, 3, 2, 1, 0}
	if a != expected {
		t.Errorf("result: %v, expected: %v", a, expected)
	}
}
