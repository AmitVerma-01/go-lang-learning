package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	a := add(1, 2)
	b := add(3, 4)
	

	fmt.Println(a, b)
	c := prod(add, 1, 2)
	fmt.Println(c)
}

func prod(fn func(c,d int) int, a, b int) int {
	return fn(a,b) * fn(a,b)
}