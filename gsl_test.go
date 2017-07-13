package main_test

import (
	"testing"
)

func BenchmarkGSL(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for i := 0; i < b.N; i++ {
		// Start of algorithm
		mmSimple(&A, &B, &C)
		// Finish of algorithm
	}
}
