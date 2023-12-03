package main

import "testing"

func equal(a, b []int) bool {
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

func Test_rotate(t *testing.T) {
	data := []struct {
		s        []int
		diff     int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, []int{3, 4, 5, 6, 7, 8, 9, 10, 1, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, []int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}},
	}

	for _, d := range data {
		rotate(d.s, d.diff)
		if !equal(d.s, d.expected) {
			t.Errorf("result %v, expected %v", d.s, d.expected)
		}
	}
}
