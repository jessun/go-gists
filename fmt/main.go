package main

import "fmt"

type i interface {
	GetString() string
}

type a struct {
	s string
}

func (p *a) GetString() string {
	return p.s
}

func printInterface(v i) {
	fmt.Printf("%#v", v)
}

func main() {
	var newA i = &a{s: "hello"}
    printInterface(newA)
}
