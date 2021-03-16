package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func ComputeUse() {
	hyhot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hyhot))
	fmt.Println(hyhot(5, 12))
	fmt.Println(compute(math.Pow))
}