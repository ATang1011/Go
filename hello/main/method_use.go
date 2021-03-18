package main

import (
	"fmt"
	"math"
	"yy.little.cat/you"
)

//method : 方法只是个带接收者参数的函数。
//你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
//error: cannot define new methods on non-local type you.YouInt
//error: cannot use y (type you.YouInt) as type int in return argument

//func (y you.YouInt)AbsInt() int  {
//	return y
//}

type Vertex struct {
	X, Y float64
}

type Myfloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f Myfloat)Absf() Myfloat  {
	if f < 0 {
		return -f
	}else {
		return f
	}
}

func (hi aha) AbsAha() string {
	return hi.a + " & " + hi.h
}

func (v *Vertex)scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Method() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	(&v).scale(10)
	fmt.Println(v)

	f := Myfloat(-math.Sqrt2)
	fmt.Println(f.Absf())

	hi := aha{"Tom", "Jerry"}
	fmt.Println(hi.AbsAha())

	you.Card()
	fmt.Println(you.YouInt(8))
}

//https://tour.go-zh.org/methods/4

