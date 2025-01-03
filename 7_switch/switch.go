package main

import (
	"fmt"
	"time"
)

// func main() {
// 	i := 10

// 	switch i {
// 	case 10:
// 		fmt.Println("i is 10")
// 	case 20:
// 		fmt.Println("i is 20")
// 	default:
// 		fmt.Println("i is not 10 or 20")
// 	}
// }

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a workday")
	}

	// type swtich 

	whoAmI := func (i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("i is int")
		case string:
			fmt.Println("i is string")
		default:
			fmt.Println("i is not int or string")
		}
	}

	whoAmI(10)
	whoAmI("hello")
	whoAmI(true)
	whoAmI(nil)



}