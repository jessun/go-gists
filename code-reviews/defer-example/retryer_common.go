package main

import "time"

// ==========================
type RetrySignal = int

const (
	FuncSuccess RetrySignal = iota
	FuncFailure
	FuncComplete
	FuncError
)

func minDuration(a time.Duration, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}

// ==========================
