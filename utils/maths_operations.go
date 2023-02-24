package utils

import "math"

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func CalculateFees(operation string, number float64) float64 {
	var fee float64

	switch operation {
	case "withdraw":
		fee = RoundFloat(number*0.01, 2)
	case "checkBalance":
		fee = 0.36
	}

	return fee
}
