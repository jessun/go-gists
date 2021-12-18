package main

import (
	"fmt"
	"sync/atomic"
)

var g = atomic.Value{}

func init() {

}

func main() {
	a := g.Load()
	v, ok := a.(string)
	fmt.Println(v, ok)
}
