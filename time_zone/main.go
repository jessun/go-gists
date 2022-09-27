package main

import (
	"fmt"
	"time"
)

func main() {
	s := "2022-09-20 15:20:22.898634"
	FORMAT_TIME_SQL := "2006-01-02 15:04:05.000000"
	t, err := time.Parse(FORMAT_TIME_SQL, s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", t)


	t2, err := time.ParseInLocation(FORMAT_TIME_SQL, s, time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", t2)
	t3, err := time.LoadLocation("+08:00")
	fmt.Printf("%v\n", t3)
}
