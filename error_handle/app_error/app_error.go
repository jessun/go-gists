package apperror

import (
	"errors"
	"fmt"
)

type customError struct {
	Err     error
	ErrCode int
}

const ErrCodeWrongAnswer = 1

func (e *customError) Error() string {
	return fmt.Sprintf("(%v)%v", e.ErrCode, e.Err.Error())
}

type generalMathTool interface {
	calculate(a, b, expected int) error
}

type simpleMultiplier struct{}

func (m *simpleMultiplier) calculate(a, b, expected int) error {
	if a*b == expected {
		return nil
	}
	return errors.New("wrong answer")
}

type multiplierWithAnalyzer struct{}

func (m *multiplierWithAnalyzer) calculateRaw(a, b, expected int) *customError {
	if a*b == expected {
		return nil
	}

	fmt.Printf("a[%v], b[%v], expected[%v]", a, b, expected)

	return &customError{errors.New("wrong answer"), ErrCodeWrongAnswer}
}

type multiplierWithAnalyzerWrapper func(a, b, c int) *customError

func (fn multiplierWithAnalyzerWrapper) calculate(a, b, c int) error {
	if err := fn(a, b, c); err != nil {
		// do something with err
		return err
	}
	// do something else

	return nil
}

func Calculate(t generalMathTool, a, b, c int) error {
	return t.calculate(a, b, c)
}

func Question1() {

	var m1 = &simpleMultiplier{}

	if err := Calculate(m1, 1, 2, 3); err != nil {
		fmt.Printf("m1 error: %v", err)
	}

	var m2 = &multiplierWithAnalyzer{}

	if err := Calculate(multiplierWithAnalyzerWrapper(m2.calculateRaw), 1, 2, 3); err != nil {
		fmt.Printf("m1 error: %v", err)
	}

}
