package example

import (
	"fmt"
	"math/rand"
)

//ExampleHello is
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

//ExampleSalutations is
func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}

//ExamplePerm is
func ExamplePerm() {
	for _, value := range rand.Perm(5) {
		fmt.Println(value)
	}
	// Unordered output:
	// 4
	// 2
	// 1
	// 3
	// 0
}
