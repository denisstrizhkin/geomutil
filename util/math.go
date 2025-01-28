package util

import (
	"math"
)

func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func Inf(sign int) float32 {
	return float32(math.Inf(sign))
}

func Sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

func Cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func DegToRad(deg float32) float32 {
	return deg * math.Pi / 180.0
}
