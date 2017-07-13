package main

import (
	"testing"
)

func BenchmarkBuffer(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrixSingle()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmBuffer(&A, &B, &C)
		// Finish of algorithm
	}
}
