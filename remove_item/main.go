package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	for i := range s {
		ss := s[i]
		if ss == 1 {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}

	fmt.Printf("%v", s)
}
