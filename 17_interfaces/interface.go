package main
import (
	"fmt"
)

type Shape struct {
	name string
	side float64
}

func (s Shape) area() float64 {
	return s.side * s.side
}

func (s Shape) perimeter() float64 {
	return 2 * s.side
}

func main() {
	s := Shape{name: "square", side: 5}
	a := s.area()
	p := s.perimeter()
	fmt.Println(a, p)
}