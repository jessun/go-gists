package main

import (
	"fmt"
)

type P struct {
	value string
}

func (p *P) Print() {
	fmt.Printf("p.value[%v]\n", p.value)
}

type Q struct {
	value string
}

func (q Q) Print() {
	fmt.Printf("q.value[%v]\n", q.value)
}

type R struct {
	value string
}

func (r *R) Print() {
	fmt.Printf("r.value[%v]\n", r.value)
}

func main() {
	// ===============================
	p1 := P{value: "p1"}
	defer p1.Print()

	p2 := &P{value: "p2"}
	defer p2.Print()

	p1 = P{value: "p1.update"}
	p2 = &P{value: "p2.update"}

	defer fmt.Println()

	// ===============================
	q1 := Q{value: "q1"}
	defer q1.Print()

	q2 := &Q{value: "q2"}
	defer q2.Print()

	q1 = Q{value: "q1.update"}
	q2 = &Q{value: "q2.update"}

	defer fmt.Println()

	// ===============================
	r1 := R{value: "r1"}
	defer r1.Print()

	r2 := &R{value: "r2"}
	defer r2.Print()

	r1.value = "r1.update"
	r2.value = "r2.update"
}
