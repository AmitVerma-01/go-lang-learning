package main

import "fmt"

// for -> only construct in go for looping 
func main(){
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	println("value of i after loop", i)

	// classic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	i=0
	// infinite loop
	for {
		fmt.Println("infinite loop")
		if i==10 {
			break
		} 
		i++
	}
	
	// for range
	for i := range 5 { 
		fmt.Println(i)
	}
}