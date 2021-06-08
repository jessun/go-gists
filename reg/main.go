package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "test"

	ss := []string{"test", "atest", "TEST"}

	re, err := regexp.Compile(s)
	if err != nil {
		panic(err)
	}

	for i := range ss {
		item := ss[i]
		bl := re.MatchString(item)
		fmt.Printf("re[%v], item[%v], match[%v]\n", s, item, bl)
	}

}
