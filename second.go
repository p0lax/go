/*Closure function test*/
package main

import (
	"fmt"
)

type funcs func() int
const FUNCS_LIST_SIZE int = 10

func main() {
	x := make([]funcs, FUNCS_LIST_SIZE)
	for i := 0; i < FUNCS_LIST_SIZE; i ++{
		func(i int) {
			x[i] = func () int {
				return i;
			}
		}(i)
	}
	fmt.Println(x)
	for _, j := range x {
		fmt.Println(j())
	}
}

