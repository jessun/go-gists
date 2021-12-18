package main

import "fmt"

func main() {
	m := map[string]int{}

	m["a"]++
	m["b"]++

	fmt.Printf("%#v", m)

	a := []string{}
	a = append(a, "a")
	a = append(a, "b")
	for i := range a {
	}
}
