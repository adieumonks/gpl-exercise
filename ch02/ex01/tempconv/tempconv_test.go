package tempconv_test

import (
	"fmt"
	"testing"

	"github.com/adieumonks/gpl-exercise/ch02/ex01/tempconv"
)

func TestCToF(t *testing.T) {
	if tempconv.CToF(tempconv.BoilingC) != 212.0 {
		t.Error(fmt.Sprint(tempconv.CToF(tempconv.BoilingC)))
	}
}

func TestFtoC(t *testing.T) {
	if tempconv.FToC(212.0) != tempconv.BoilingC {
		t.Error(fmt.Sprint(tempconv.FToC(212.0)))
	}
}

func TestCtoK(t *testing.T) {
	if tempconv.CToK(0) != -tempconv.Kelvin(tempconv.AbsoluteZeroC) {
		t.Error(fmt.Sprint(tempconv.CToK(0)))
	}
}

func TestKtoC(t *testing.T) {
	if tempconv.KToC(0) != tempconv.AbsoluteZeroC {
		t.Error(fmt.Sprint(tempconv.KToC(0)))
	}
}

func TestFtoK(t *testing.T) {
	expectd := tempconv.CToK(tempconv.FToC(0))
	if tempconv.FToK(0) != tempconv.Kelvin(expectd) {
		t.Error(fmt.Sprint(tempconv.FToK(0)))
	}
}

func TestKtoF(t *testing.T) {
	expectd := tempconv.CToF(tempconv.KToC(0))
	if tempconv.KToF(0) != expectd {
		t.Error(fmt.Sprint(tempconv.KToF(0)))
	}
}
