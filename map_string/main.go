package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]int{"a": 1, "b": 2}
	s, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}
