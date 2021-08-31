package main

import "fmt"

type sub struct {
	name string
	c    chan struct{}
}

func main() {

	s1 := sub{
		name: "1",
		c:    make(chan struct{}, 0),
	}

	counter := 0

	for {
		select {
		case a := <-s1.c:
			fmt.Println(a)
			return
		default:
			counter++
			fmt.Println("counter: ", counter)

			if counter > 100 {
				close(s1.c)
			}

		}
	}
}
