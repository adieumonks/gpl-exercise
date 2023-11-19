package popcount_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/adieumonks/gpl-exercise/ch02/ex05/popcount"
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

func TestPopCountByBitShift(t *testing.T) {
	if popcount.PopCountByBitShift(0) != 0 {
		t.Error(fmt.Sprint(popcount.PopCountByBitShift(0)))
	}

	if popcount.PopCountByBitShift(1) != 1 {
		t.Error(fmt.Sprint(popcount.PopCountByBitShift(1)))
	}

	if popcount.PopCountByBitShift(10) != 2 {
		t.Error(fmt.Sprint(popcount.PopCountByBitShift(10)))
	}

	if popcount.PopCountByBitShift(math.MaxUint64) != 64 {
		t.Error(fmt.Sprint(popcount.PopCountByBitShift(math.MaxUint64)))
	}
}

func TestPopCountByClearingBit(t *testing.T) {
	if popcount.PopCountByClearingBit(0) != 0 {
		t.Error(fmt.Sprint(popcount.PopCountByClearingBit(0)))
	}

	if popcount.PopCountByClearingBit(1) != 1 {
		t.Error(fmt.Sprint(popcount.PopCountByClearingBit(1)))
	}

	if popcount.PopCountByClearingBit(10) != 2 {
		t.Error(fmt.Sprint(popcount.PopCountByClearingBit(10)))
	}

	if popcount.PopCountByClearingBit(math.MaxUint64) != 64 {
		t.Error(fmt.Sprint(popcount.PopCountByClearingBit(math.MaxUint64)))
	}
}

var results []int

func BenchmarkPopCount(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, popcount.PopCount(uint64(i)))
	}
}

func BenchmarkPopCountByBitShift(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, popcount.PopCountByBitShift(uint64(i)))
	}
}

func BenchmarkPopCountByClearingBit(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, popcount.PopCountByClearingBit(uint64(i)))
	}
}
