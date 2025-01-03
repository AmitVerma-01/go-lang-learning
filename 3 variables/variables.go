package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	// var c int = a + b
	var c = a + b
	fmt.Println(c)
	var isAdult bool = true
	if isAdult {
		fmt.Println("adult")
	} else {
		fmt.Println("child")
	}

	name := "Amit"
	fmt.Println(name)

	var price float32 = 10.5
	// var price  = 10.5
	// var price float64 = 10.5
	// price := 10.5
	fmt.Println(price)
}