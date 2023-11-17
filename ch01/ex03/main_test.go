package main

import "testing"

func TestEcho1(t *testing.T) {
	result := echo1([]string{"a", "b", "c", "d", "e", "f", "g"})
	if result != "a b c d e f g" {
		t.Errorf("unexpected result %s", result)
	}
}

func TestEcho2(t *testing.T) {
	result := echo2([]string{"a", "b", "c", "d", "e", "f", "g"})
	if result != "a b c d e f g" {
		t.Errorf("unexpected result %s", result)
	}
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1([]string{"a", "b", "c", "d", "e", "f", "g"})
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2([]string{"a", "b", "c", "d", "e", "f", "g"})
	}
}
