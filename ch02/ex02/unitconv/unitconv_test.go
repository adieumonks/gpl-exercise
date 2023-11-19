package unitconv_test

import (
	"fmt"
	"testing"

	"github.com/adieumonks/gpl-exercise/ch02/ex02/unitconv"
)

func TestCToF(t *testing.T) {
	if unitconv.CToF(unitconv.BoilingC) != 212.0 {
		t.Error(fmt.Sprint(unitconv.CToF(unitconv.BoilingC)))
	}
}

func TestFtoC(t *testing.T) {
	if unitconv.FToC(212.0) != unitconv.BoilingC {
		t.Error(fmt.Sprint(unitconv.FToC(212.0)))
	}
}

func TestCtoK(t *testing.T) {
	if unitconv.CToK(0) != -unitconv.Kelvin(unitconv.AbsoluteZeroC) {
		t.Error(fmt.Sprint(unitconv.CToK(0)))
	}
}

func TestKtoC(t *testing.T) {
	if unitconv.KToC(0) != unitconv.AbsoluteZeroC {
		t.Error(fmt.Sprint(unitconv.KToC(0)))
	}
}

func TestFtoK(t *testing.T) {
	expectd := unitconv.CToK(unitconv.FToC(0))
	if unitconv.FToK(0) != unitconv.Kelvin(expectd) {
		t.Error(fmt.Sprint(unitconv.FToK(0)))
	}
}

func TestKtoF(t *testing.T) {
	expectd := unitconv.CToF(unitconv.KToC(0))
	if unitconv.KToF(0) != expectd {
		t.Error(fmt.Sprint(unitconv.KToF(0)))
	}
}
