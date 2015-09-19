package main

import (
	"fmt"
	"test"
)

func main() {
	xs := []float64{1,2,3,4}
	avg := test.Average(xs)
	fmt.Println(avg)
}