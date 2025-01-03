package main

import "fmt"

type Colour string

const (
	RED   Colour = "Red"
	GREEN        = "Green"
	BLUE         = "Blue"
)

func getColor(c Colour) string {
	return string(c)
}
func main() {
	fmt.Println(getColor(RED))
}
