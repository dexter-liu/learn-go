package calculator

import (
	gobmi "github.com/dexter-liu/go-bmi"
)

func CalcBMI(tall float64, weight float64) (bmi float64) {
	bmi, _ = gobmi.BMI(weight, tall)
	return
}
