// +build js

package math

import (
	stdmath "math"
)

func Cos(x float32) float32 {
	return float32(stdmath.Cos(float64(x)))
}

func Sin(x float32) float32 {
	return float32(stdmath.Sin(float64(x)))
}
