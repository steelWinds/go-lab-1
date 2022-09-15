package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func getMinNumber(rows float64, cols float64) float64 {
	absRows := int(math.Abs(rows))
	absCols := int(math.Abs(cols))
	absSize := absRows * absCols

	if absRows < 1 || absCols < 1 {
		return 0
	}

	primes := randFloats(absSize)

	matrix := mat.NewDense(absRows, absCols, primes)

	return mat.Min(matrix)
}
