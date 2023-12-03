package main

import (
	"testing"
)

func equal(a, b []string) bool {
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

func Test_reduct(t *testing.T) {
	data := []struct {
		s        []string
		expected []string
	}{
		{
			[]string{},
			[]string{},
		},
		{
			[]string{"a"},
			[]string{"a"},
		},
		{
			[]string{"a", "b"},
			[]string{"a", "b"},
		},
		{
			[]string{"a", "a"},
			[]string{"a"},
		},
		{
			[]string{"a", "a", "a"},
			[]string{"a"},
		},
		{
			[]string{"a", "a", "b"},
			[]string{"a", "b"},
		},
		{
			[]string{"a", "a", "b", "b"},
			[]string{"a", "b"},
		},
		{
			[]string{"a", "a", "b", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			[]string{"a", "a", "b", "b", "c", "c"},
			[]string{"a", "b", "c"},
		},
	}

	for _, d := range data {
		result := reduct(d.s)
		if !equal(result, d.expected) {
			t.Errorf("result: %v, expected: %v", result, d.expected)
		}
	}
}
