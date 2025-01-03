package main

import (
	"fmt"
	"strings"
)

// default false value int=0, float=0, string="", bool=false
func main() {

	var nums[4]int

	fmt.Println(nums)
	fmt.Println(len(nums))

	var val[5]float32

	fmt.Println(val)

	var strs = [5]string{"hello", "world", "go", "golang", "go"}
	fmt.Println(strings.Split(strs[0], "l"))
}