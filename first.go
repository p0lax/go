package main

import (
	"fmt"
)

type Test1 struct {
	a, b, c float64
	str string
}

func (test *Test1) A() int {
	return 10 + 1
}

type Test2 struct {
	Name string
	Test1 Test1
}

func (test *Test2) A() int {
	return 10 + 1
}

func main() {
	test := Test2{Name: "11111", Test1: Test1{0, 0, 0, "1111"}}
	fmt.Println(test)
}

