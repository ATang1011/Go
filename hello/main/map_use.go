package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func UseMap() {
	type Vertex struct {
		Lat, Long float64
	}
	m := make(map[string]Vertex)
	m["Haha"] = Vertex{12120.001, 45512167}
	m["Ahha"] = Vertex{12120.001, 45512167}
	fmt.Println("map ", m)
	delete(m, "Haha")

	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	a := strings.Fields(s)
	var m = make(map[string]int)
	for _, value := range a {
		_, ok := m[value]
		if ok == true {
			m[value] += 1
		}else {
			m[value] = 1
		}
	}
	return m
}