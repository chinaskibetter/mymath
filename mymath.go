package mymath

import "math"

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Yn(x int, y float64) float64 {
	return math.Yn(x, y)
}
