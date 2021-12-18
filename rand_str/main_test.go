package main

import (
	"testing"
)

func BenchmarkGenRandomThreadID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genRandomThreadID()
	}
}

func BenchmarkGenRandomThreadIDV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genRandomThreadIDV2(9)
	}

}

func BenchmarkGenRandomThreadIDV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genRandomThreadIDV4(9)
	}

}
