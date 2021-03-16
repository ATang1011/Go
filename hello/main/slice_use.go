package main

import "fmt"

func LookSlice() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	//切片的长度就是它所包含的元素个数。
	//切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	//截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)

	var w []int
	fmt.Println(w, len(w), cap(w))
	if w == nil {
		fmt.Println("nil!")
	}

	a := make([]int, 4)
	b := make([]int, 5, 5)
	c := b[:2]
	d := c[2:5]
	printSlice(a)
	printSlice(b)
	printSlice(c)
	printSlice(d)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	//切片，像对数组的引用，修改切片中的值，也同样修改了原数组以及其他切片中对应的值。
	//没有长度的数组
	var s2 []int = primes[1:4]
	fmt.Println(s2)
}

func AppendSlice() {

	var s []int
	printSlice(s)

	// 添加一个空切片
	s1 := append(s, 0)
	printSlice(s1)

	// 这个切片会按需增长
	s2 := append(s, 1)
	printSlice(s2)

	// 可以一次性添加多个元素
	//s = append(s, 2, 3)
	//printSlice(s)
	s3 := append(s, 2, 3, 4)
	printSlice(s3)
	s4 := append(s3, 2, 3, 4, 7)
	printSlice(s4)

	//切片的容量？

	for i, v := range s4 {
		fmt.Printf("2 ** %d == %d\n", i, v)
	}

	for i := range s4 {
		fmt.Println(i)
	}

	for _, value := range s4 {
		fmt.Println(value)
	}

	//pic.Show(Pic)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for x := range a {
		b := make([]uint8, dx)
		for y := range b {
			b[y] = uint8(x * y -  1)
		}
		a[x] = b
	}
	return a
}