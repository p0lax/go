package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := make([]interface{}, 10, 10)
	for i := 0; i < 10; i ++{
		x[i] = func () int {
			return i;
		}
	}
	fmt.Println(x)
	for _, j := range x {
		f := reflect.ValueOf(j)
		fmt.Println(f)
	}
}

