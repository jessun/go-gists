package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	a := uuid.NodeID()
	fmt.Printf("%x\n", a)
}
