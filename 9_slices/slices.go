package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(len(nums))
	


	val := make([]int, 0,4)
	println(cap(val))
	val = append(val, 1, 2, 3, 4, 5)
	println(cap(val))
	// for i, v := range val {
	// 	fmt.Println(i, v)
	// }

	// for i, v := range nums {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println("----------------")
	// nums = append(nums, 6, 7, 8)
	// nums = append(nums, 9)
	// for i, v := range nums {
	// 	fmt.Println(i, v)
	// }

	// copy function

	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	// for i, v := range nums2 {
	// 	fmt.Println(i, v)
	// }
	fmt.Println(nums[3:])

}