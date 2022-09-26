package main

import "fmt"

func main() {
	l := 5

	s := make([]string, l)

	for i := range s {
		s[i] = "a"
	}

	fmt.Printf("%#v", s)
	fmt.Printf("%+v", s)
}
