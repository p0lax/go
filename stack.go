package main

import (
	"fmt"
	"stacker/stack"
)

func main () {
	var haystack stack.Stack
	haystack.Push("111")
	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}