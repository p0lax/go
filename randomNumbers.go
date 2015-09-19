package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createArray(len int) []int {
	var result []int;
	for i := 0; i < len; i ++ {
		result = append(result, rand.Intn(10))
	}
	return result
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(createArray(5))
}