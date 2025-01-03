package main

import (
	// "math/rand"
	"fmt"
	// "time"
)

// func processMessage(messageChan chan int) {
// 	for num := range messageChan {
// 		fmt.Println("processMessage", num)
// 	}
// 	// fmt.Println("processMessage", <-messageChan)
// }
// func main() {
// 	messageChan := make(chan int)

// 	go processMessage(messageChan)
// 	for {
// 		messageChan <- rand.Intn(10)
// 		time.Sleep(1 * time.Second)
// 	}
// 	// time.Sleep(5 * time.Second)
// }

func sum(result chan int, nums ...int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	result <- sum
}

func main() {
	result := make(chan int) 
	go sum(result, 1, 2, 3, 4, 5)
	fmt.Println(<-result)
}