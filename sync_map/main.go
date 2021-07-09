package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}

	m.Store(1, "1")
	m.Store(2, "2")

	m.Range(func(key, value interface{}) bool {
		fmt.Println(key)
		return true
	})
}
