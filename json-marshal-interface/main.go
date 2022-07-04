package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string      `json:"name"`
	Age  int         `json:"age"`
	Pvt  interface{} `json:"pvt"`
}

type PvtA struct {
	A int `json:"a"`
}

type PvtB struct {
	B int `json:"b"`
}

type PvtC struct {
	A PvtA `json:"a"`
	B PvtB `json:"b"`
}

func main() {
	p := Person{Name: "John", Age: 30, Pvt: PvtC{A: PvtA{A: 1}, B: PvtB{B: 2}}}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}
