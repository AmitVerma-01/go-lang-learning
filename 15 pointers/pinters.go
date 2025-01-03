package main

import (
	"fmt"
)

func changeNum(num *int) {
	*num = 10
}

func main() {
	n:=20
	changeNum(&n)
	fmt.Println(n)
}