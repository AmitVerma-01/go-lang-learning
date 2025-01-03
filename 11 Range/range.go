package main

import (
	"fmt"
)


// iterate over data structures
func main() {
	// iterate over a slice
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Println(i, v)
	}
	sum:=0
	// iterate over a map
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m {
		fmt.Println(k, v)
		sum+=v
	}
	fmt.Println(sum)

	// iterate over a string
	str := "hello"
	for i, v := range str {
		fmt.Println(i, string(v))
	}

}