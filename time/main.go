package main

import (
	"fmt"
	"regexp"
)

func main() {
	// creat a game object
	// start the game

	// game loop

	// creat a slice of strings
	// t := "2006-01-02 15:04:05"
	// s := "2021-11-01 00:00:00"
	// ss, err := time.ParseInLocation(t, s, time.Local)
	// if err != nil {
	//     panic(err)
	// }

	reg := regexp.MustCompile(`debug: set binlog file expiration time\[(.*)\]`)
	s := "debug: set binlog file expiration time[2021-11-23 00:00:00]"
	if reg.MatchString(s) {
		matches := reg.FindAllStringSubmatch(s, -1)
		fmt.Printf("%v", matches)
	}
}
