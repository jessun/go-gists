package main

import (
	"fmt"
	"regexp"
)

func main() {
	x := `[maxmemory-policy noeviction]`
	re, err := regexp.Compile(`\[maxmemory-policy (.*)\]`)
	if err != nil {
		panic(err)
	}
	r := re.FindStringSubmatch(x)
	fmt.Printf("r: %v", r[1])

}
