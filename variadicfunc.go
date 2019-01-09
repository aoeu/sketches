package main

import (
	"fmt"
)

var ints = make([]int, 0)

func put(ints []int, vals ...int) []int {
	ints = append(ints, vals...)
	return ints
}

func main() {
	fmt.Println(ints) // []
	ints = put(ints, 1)
	fmt.Println(ints) // [1]
	ints = put(ints, 2, 3)
	fmt.Println(ints) // [1 2 3]
}
