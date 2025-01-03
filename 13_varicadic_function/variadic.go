package main
import (
	"fmt"
)

func sum(a ...int) int {
	var total int
	for _, v := range a {
		total += v
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	total := sum(nums...)
	fmt.Println(total)
	a := sum(1, 2, 3, 4, 5)
	b := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	c := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	fmt.Println(a, b, c)
}