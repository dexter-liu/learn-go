package main

import (
	"testing"

	calc "github.com/dexter-liu/learn-go/ch3/001.fatrate.refactor/calc"
)

func TestMain(t *testing.T) {
	var expectedBMI float64 = 18.9
	actualBMI := calc.CalcBMI(1.7, 69)
	if expectedBMI != actualBMI {
		t.Fail()
		// t.Errorf("expected %f, but got %f", expectedBMI, actualBMI)
	}
}
