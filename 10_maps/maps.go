package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	println(len(m))
	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "one")
	println(len(m))


	mp := map[string]int{"one": 1, "two": 2, "three": 3}
	println(mp["one"])
	println(mp["two"])

	j,k:=mp["one"]

	if k {
		println("found",k,j)
	} else {
		println("not found",j)
	}

	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	m2 := map[string]int{"four": 4, "five": 5, "six": 6}

	fmt.Println(maps.Equal(m1, m2))
}
