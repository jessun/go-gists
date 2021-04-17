package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Unix(0, 1618355617000000000)
	fmt.Printf("1618355617000000000 : %v\n", t1.Format(time.RFC3339))
	t2 := time.Unix(0, 1618384393000000000)
	fmt.Printf("1618355617000000000 : %v\n", t2.Format(time.RFC3339))

	t3 := time.Now()
	fmt.Printf("now : %v\n", t3.Format(time.RFC3339))

	t4 := time.Now().UTC()
	fmt.Printf("now : %v\n", t4.Format(time.RFC3339))
}
