package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func getFloats(c chan []float64, size int) {
	c <- randFloats(size)
}

func createMatrix(c chan []float64, rows int, cols int) {
	primes := <-c

	mat.NewDense(rows, cols, primes)
}

func getMinNumber(rows float64, cols float64) {
	absRows := int(math.Abs(rows))
	absCols := int(math.Abs(cols))
	absSize := absRows * absCols

	if absRows < 1 || absCols < 1 {
		return
	}

	var c chan []float64 = make(chan []float64)

	go getFloats(c, absSize)
	go createMatrix(c, absRows, absCols)
}
