package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

import "yy.little.cat/you"

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
	you.Card()
	fmt.Println("人生苦短，Let's Go!")
	fmt.Println("The time is ", time.Now())
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Println("My favorite number is ", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(Add(9, 10))
	a, b := Swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println(Split(17))
	fmt.Println("Name in main ", name)
	school := 0
	fmt.Println("Your name in main is ", school)

	InitialValue()
	TypePrint()

	fmt.Println(NeedInt(Small))
	fmt.Println(NeedFloat(Small))
	fmt.Println(NeedFloat(Big))
	Loop()
	LookSlice()
	AppendSlice()
	UseMap()
	ComputeUse()
	Closures()
	Fibonacci()
}



