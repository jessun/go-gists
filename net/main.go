package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", addrs)
}
