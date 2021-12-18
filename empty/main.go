package main

import (
	"fmt"
	"strings"
)

func main() {
	a := ""
	b := strings.Split(a, ",")
	fmt.Printf("%#v", len(b))
}
