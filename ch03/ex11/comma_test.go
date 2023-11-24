package main

import (
	"testing"
)

var data = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"1", "1"},
	{"12", "12"},
	{"123", "123"},
	{"1234", "1,234"},
	{"1.", "1."},
	{"+1", "+1"},
	{"-1", "-1"},
	{"1.2345", "1.2345"},
	{"+1.2345", "+1.2345"},
	{"-1.2345", "-1.2345"},
	{"123456789.123456789", "123,456,789.123456789"},
	{"+123456789.123456789", "+123,456,789.123456789"},
	{"-123456789.123456789", "-123,456,789.123456789"},
}

func Test_comma(t *testing.T) {
	for _, d := range data {
		result := comma(d.input)
		if result != d.expected {
			t.Errorf("result %s, expected %s", result, d.expected)
		}
	}
}
