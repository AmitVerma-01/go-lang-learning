package main

import "fmt"

func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type Stack[T int | string | bool] struct {
	items []int
}


// main demonstrates the use of type parameters in a function
// and in a struct.
func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// printSlice(nums)
	// names := []string{"GoLang","Typescript", "Jane", "Jack"}
	// printSlice(names)
	// isValid := []bool{true, false, true, false}
	// printSlice(isValid)

	s := Stack[int]{[]int{1, 2, 3, 4, 5}}
	printSlice(s.items)
}