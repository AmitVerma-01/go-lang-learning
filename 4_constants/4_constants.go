package main

import "fmt"
const age = 20  
func main() {
	const name = "Amit"
	fmt.Println(name)

	//constants grouping
	const (
		pi = 3.14159265358979323846
		e  = 2.71828182845904523536
	)  
	fmt.Println(pi)
	fmt.Println(e)
	println(name, age)
}