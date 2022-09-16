package main

import "fmt"

var _ ServiceErr = (*ClassAErr)(nil)

// An Error represents a network error.
type ServiceErr interface {
	error
	Timeout() bool   // Is the error a timeout?
	Temporary() bool // Is the error temporary?
}

type ClassAErr struct {
}

func (e *ClassAErr) Error() string {
	return "ClassAErr"
}

func (e *ClassAErr) Timeout() bool {
	return true
}

func (e *ClassAErr) Temporary() bool {
	return true
}

type ClassBErr struct{}

func (e *ClassBErr) Error() string {
	return "ClassBErr"
}

func (e *ClassBErr) Timeout() bool {
	return true
}

func (e *ClassBErr) Temporary() bool {
	return false
}

func analysisError(e error) {
	if e == nil {
		return
	}

	if err, ok := e.(ServiceErr); ok {
		if err.Timeout() {
			// do something
			fmt.Printf("Timeout true\n")
		}
		if err.Temporary() {
			// do something
			fmt.Printf("Temporary true\n")
		} else {
			fmt.Printf("Temporary false\n")
		}
	}

}

func main() {
	fmt.Printf("begin to analysis error a\n")
	analysisError(&ClassAErr{})
	fmt.Printf("analysis error a over\n")

	fmt.Println("=====================")

	fmt.Printf("begin to analysis error b\n")
	analysisError(&ClassBErr{})
	fmt.Printf("analysis error b over\n")
}
