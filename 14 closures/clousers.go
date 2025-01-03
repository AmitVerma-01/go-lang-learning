package main

import (
	"fmt"
)

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	counter1 := counter()
	counter2 := counter()
	counter3 := counter()
	
	fmt.Println(counter1(), counter2(), counter3())
	fmt.Println(counter1(),  counter3())
	fmt.Println(counter2(), counter3())
}