package test

import (
	"fmt"
)

//Add is
func Add(n ...int) int {
	if len(n) == 0 {
		fmt.Println("there is no input")
		return 0
	}
	sum := 0
	for _, i := range n {
		sum += i
	}
	return sum
}
