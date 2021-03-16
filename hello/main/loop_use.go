package main

import "fmt"
import "runtime"

func Loop()  {
	//defer 语句会将函数推迟到外层函数返回之后执行。推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	//defer typeTransport()
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

	TimeDay()
}