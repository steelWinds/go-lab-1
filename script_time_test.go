package main

import (
	"fmt"

	"testing"
)

type testTable = struct {
	input float64
}

var tableInputs = []testTable{
	{input: 1},
	{input: 10},
	{input: 100},
}

func BenchmarkTime(b *testing.B) {
	for _, v := range tableInputs {
		b.Run(fmt.Sprintf("input_size_%d", int(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				getMinNumber(v.input, v.input)
			}
		})
	}
}
