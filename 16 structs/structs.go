package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type Student struct {
	Person
	School string
}

func newPerson(name string, age int) *Person {
	return &Person{name, age}
}
func (P *Person) changeAge(age int){
	P.Age = age
}
func main() {
	p := Person{"John", 20}
	fmt.Println(p)
	p.Name = "Jane"
	p.changeAge(21)
	fmt.Println(p)
	q := *newPerson("John", 20)
	fmt.Println(q)
	s := Student{Person: Person{"John", 20}, School: "UCLA"}
	fmt.Println(s)
}
