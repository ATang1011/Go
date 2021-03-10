package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

var (
	name  string     = "ZYT"
)

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func main() {
	fmt.Println("人生苦短，Let's Go!")
	fmt.Println("The time is ", time.Now())
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Println("My favorite number is ", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(9, 10))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println("Name in main ", name)
	school := 0
	fmt.Println("Your name in main is ", school)

	initialValue()
	typePrint()

	//defer 语句会将函数推迟到外层函数返回之后执行。推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	//defer typeTransport()

	const pi = 3.14
	fmt.Println("Pi is ", pi)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	for i := 0; i < 10; i++ {
		//defer fmt.Println("For i is ", i)
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x.")
		fallthrough
	case "Linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("Others os.")
	}

	timeDay()

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

	primes := [6]int{2, 3, 5, 7, 11, 13}

	//切片
	var s []int = primes[1:4]
	fmt.Println(s)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	var name = "YTY"
	fmt.Println("Name in func ", name)
	return y, x
}

func split(sum int) (y, x int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func initialValue()  {
	var (
		isCID bool       = false
		age   int        = 29
		z     complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", isCID, isCID)
	fmt.Printf("Type: %T Value: %v\n", age, age)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func typePrint()  {
	var (
		i int
		f float64
		w bool
		s string
	)
	fmt.Printf("%v %v %v %q\n", i, f, w, s)
}

func typeTransport() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func timeDay()  {
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