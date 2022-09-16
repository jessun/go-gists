package main

import (
	"errors"
	"fmt"
)

type ErrorInvalidArgument struct {
	mysqlInstID string
}

func (e *ErrorInvalidArgument) Error() string {
	return fmt.Sprintf("Invalid argument: %s", e.mysqlInstID)
}

func NewInvalidArgumentError(mysqlInstID string) error {
	return &ErrorInvalidArgument{mysqlInstID: mysqlInstID}
}

type ErrorNotFound struct {
	mysqlInstID string
}

func (e *ErrorNotFound) Error() string {
	return fmt.Sprintf("Not found: %s", e.mysqlInstID)
}

func NewNotFoundError(mysqlInstID string) error {
	return &ErrorNotFound{mysqlInstID: mysqlInstID}
}

func main() {
	/// *****

	rawError := errors.New("simple error")
	e1 := fmt.Errorf("error: %w", rawError)
	e2 := fmt.Errorf("error: %w", rawError)
	e3 := fmt.Errorf("error: %w", e1)
	e4 := fmt.Errorf("%w", rawError)
	e5 := NewNotFoundError("test")

	if errors.Is(e1, rawError) {
		fmt.Println("e1 and raw_error are equal") // got
	} else {
		fmt.Println("e1 and raw_error are not equal")
	}

	if errors.Is(e2, rawError) {
		fmt.Println("e2 and raw_error are equal") // got
	} else {
		fmt.Println("e2 and raw_error are not equal")
	}
	if errors.Is(e3, e2) {
		fmt.Println("e2 and e3 are equal")
	} else {
		fmt.Println("e2 and e3 are not equal") // got
	}

	if errors.Is(e1, e4) {
		fmt.Println("e1 and e4 are equal")
	} else {
		fmt.Println("e1 and e4 are not equal") // got
	}

	var targetErrorIvalidArgument *ErrorInvalidArgument

	if errors.As(e1, &targetErrorIvalidArgument) {
		fmt.Println("e1 is ErrorInvalidArgument")
	} else {
		fmt.Println("e1 is not ErrorInvalidArgument") // got
	}

	var targetErrorNotFound *ErrorNotFound

	if errors.As(e5, &targetErrorNotFound) {
		fmt.Println("e4 is ErrorNotFound") // got
	} else {
		fmt.Println("e4 is not ErrorNotFound")
	}

}
