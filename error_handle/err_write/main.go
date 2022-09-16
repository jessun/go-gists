package main

import (
	"fmt"
	"io"
)

type screenWriter struct {
}

func (sw *screenWriter) Write(p []byte) (n int, err error) {
	errStr := "2"

	if string(p) == errStr {
		return len(p), fmt.Errorf("%v is error", errStr)
	}

	fmt.Printf("screenWriter: %s\n\n", string(p))

	return len(p), nil
}

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) write(p []byte) {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(p)
}

func main() {
	ew := errWriter{w: &screenWriter{}}
	ew.write([]byte("1"))
	ew.write([]byte("2"))
	ew.write([]byte("3"))
	ew.write([]byte("4"))
	ew.write([]byte("2"))
	ew.write([]byte("5"))
	if ew.err != nil {
		fmt.Printf("found error: %v", ew.err)
	}
}
