// Copyright © 2015 Yoshiki Shibata. All rights reserved.

package tempconv_test

import (
	"fmt"
	"testing"

	"github.com/YoshikiShibata/gpl/ch02/ex01/tempconv"
)

func TestCToF(t *testing.T) {
	if tempconv.CToF(tempconv.BoilingC) != 212.0 {
		t.Error(fmt.Sprint(tempconv.CToF(tempconv.BoilingC)))
	}
}

func TestKToC(t *testing.T) {
	if tempconv.KToC(0) != tempconv.AbsoluteZeroC {
		t.Error(fmt.Sprint(tempconv.KToC(0)))
	}
}

func TestCToK(t *testing.T) {
	if tempconv.CToK(0) != -tempconv.Kelvin(tempconv.AbsoluteZeroC) {
		t.Error(fmt.Sprint(tempconv.CToK(0)))
	}
}

func TestFToK(t *testing.T) {
	expected := tempconv.FToC(0) - tempconv.AbsoluteZeroC

	if tempconv.FToK(0) != tempconv.Kelvin(expected) {
		t.Error(fmt.Sprint(tempconv.FToK(0)))
	}
}
