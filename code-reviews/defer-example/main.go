package main

import (
	"fmt"
)

// P
type P struct{ value string }

func (p *P) Print() { fmt.Printf("p.value[%v]\n", p.value) }

// Q
type Q struct{ value string }

func (q Q) Print() { fmt.Printf("q.value[%v]\n", q.value) }

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
	p3 := P{value: "p3"}
	defer p3.Print()

	p4 := &P{value: "p4"}
	defer p4.Print()

	p3.value = "p3.update"
	p4.value = "p4.update"
}
