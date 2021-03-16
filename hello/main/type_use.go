package main

import (
	"math"
	"math/cmplx"
	"fmt"
	"time"
)

func Add(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	var name = "YTY"
	fmt.Println("Name in func ", name)
	return y, x
}

func Split(sum int) (y, x int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func InitialValue()  {
	var (
		isCID bool       = false
		age   int        = 29
		z     complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", isCID, isCID)
	fmt.Printf("Type: %T Value: %v\n", age, age)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func TypePrint()  {
	var (
		i int
		f float64
		w bool
		s string
	)
	fmt.Printf("%v %v %v %q\n", i, f, w, s)

	type aha struct {
		a string
		h string
	}

	ahha := aha{"aa", "hh"}
	fmt.Println(aha{"a", "ha"})
	fmt.Println(ahha.a)
	p := &ahha
	p.a = "ahhha"
	fmt.Println(ahha.a)

	const pi = 3.14
	fmt.Println("Pi is ", pi)
}

func TypeTransport() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}

func NeedInt(x int) int { return x*10 + 1 }
func NeedFloat(x float64) float64 {
	return x * 0.1
}

func TimeDay()  {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Printf("Type: %T Value: %v\n", today, today)

	//优雅的 if-else if-else...
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
