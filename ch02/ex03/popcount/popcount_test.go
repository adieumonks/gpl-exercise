package popcount_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/adieumonks/gpl-exercise/ch02/ex03/popcount"
)

func TestPopCount(t *testing.T) {
	if popcount.PopCount(0) != 0 {
		t.Error(fmt.Sprint(popcount.PopCount(0)))
	}

	if popcount.PopCount(1) != 1 {
		t.Error(fmt.Sprint(popcount.PopCount(1)))
	}

	if popcount.PopCount(10) != 2 {
		t.Error(fmt.Sprint(popcount.PopCount(10)))
	}

	if popcount.PopCount(math.MaxUint64) != 64 {
		t.Error(fmt.Sprint(popcount.PopCount(math.MaxUint64)))
	}
}

func TestPopCountWithLoop(t *testing.T) {
	if popcount.PopCountWithLoop(0) != 0 {
		t.Error(fmt.Sprint(popcount.PopCountWithLoop(0)))
	}

	if popcount.PopCountWithLoop(1) != 1 {
		t.Error(fmt.Sprint(popcount.PopCountWithLoop(1)))
	}

	if popcount.PopCountWithLoop(10) != 2 {
		t.Error(fmt.Sprint(popcount.PopCountWithLoop(10)))
	}

	if popcount.PopCountWithLoop(math.MaxUint64) != 64 {
		t.Error(fmt.Sprint(popcount.PopCountWithLoop(math.MaxUint64)))
	}
}

var results []int

func BenchmarkPopCount(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, popcount.PopCount(uint64(i)))
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, popcount.PopCount(uint64(i)))
	}
}
