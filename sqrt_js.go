// +build js

package math

import (
	stdmath "math"
)

func Sqrt(x float32) float32 {
	return float32(stdmath.Sqrt(float64(x)))
}
