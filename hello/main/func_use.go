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

func adder() func(int) int {
	sum := 0
	fmt.Println(&sum)
	return func(x int) int {
		sum += x
		return sum
	}
}

func Closures() {
	//理解为匿名函数？，是一个"内联"语句或表达式。
	//匿名函数的优越性在于可以直接使用函数内的变量，不必声明
	pos, neg, aha := adder(), adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
			aha(i*2),
		)
	}

	pos1 := adder()
	fmt.Println(pos1(3))
}

func fibonacci() func() int {
	s := []int{0, 1}
	return func() int {
		next := s[len(s) - 1] + s[len(s) - 2]
		s = append(s, next)
		return next
	}
}

func Fibonacci()  {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}