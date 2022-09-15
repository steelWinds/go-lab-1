package main

import (
	"math/rand"

	"time"
)

func randFloats(n int) []float64 {
	rand.Seed(time.Now().UnixNano())

	floats := make([]float64, n)

	for i := range floats {
		floats[i] = float64(rand.Intn(21) - 10)
	}

	return floats
}
