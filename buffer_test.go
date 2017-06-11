package main_test

import (
	"testing"
)

func TestBuffer4(t *testing.T) {
	if !isSame(mmBuffer4) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkBuffer4(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmBuffer4(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmBuffer4 - added 4 buffers
func mmBuffer4(A, B, C *[][]float64) {
	n := len(*A)
	// Create buffers
	buffer0 := make([]float64, n, n)
	buffer1 := make([]float64, n, n)
	buffer2 := make([]float64, n, n)
	buffer3 := make([]float64, n, n)
	// Now, we use (i+=4), for avoid
	// dublicate calculations
	for i := 0; i < n; i += 4 {
		for j := 0; j < n; j++ {
			// Put in buffer row of matrix [A]
			buffer0[j] = (*A)[i+0][j]
			buffer1[j] = (*A)[i+1][j]
			buffer2[j] = (*A)[i+2][j]
			buffer3[j] = (*A)[i+3][j]
		}
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				(*C)[i+0][j] += buffer0[k] * (*B)[k][j]
				(*C)[i+1][j] += buffer1[k] * (*B)[k][j]
				(*C)[i+2][j] += buffer2[k] * (*B)[k][j]
				(*C)[i+3][j] += buffer3[k] * (*B)[k][j]
			}
		}
	}
}

func TestBuffer4Assembly(t *testing.T) {
	if !isSame(mmBuffer4Assembly) {
		t.Errorf("Algorithm is not correct")
	}
}

func BenchmarkBuffer4Assembly(b *testing.B) {
	// Stop the timer for avoid add time of generate matrix
	b.StopTimer()
	A, B, C := generateMatrix()
	// Now, we are ready for start timer our benchmark
	b.StartTimer()
	// We cannot control for amount of benchmark test,
	// but it is not important
	for t := 0; t < b.N; t++ {
		// Start of algorithm
		mmBuffer4Assembly(&A, &B, &C)
		// Finish of algorithm
	}
}

// mmBuffer4Assembly - added 4 buffers
func mmBuffer4Assembly(A, B, C *[][]float64) {
	n := len(*A)
	// Create buffers
	buffer0 := make([]float64, n, n)
	buffer1 := make([]float64, n, n)
	buffer2 := make([]float64, n, n)
	buffer3 := make([]float64, n, n)
	// Now, we use (i+=4), for avoid
	// dublicate calculations
	var sum0, sum1, sum2, sum3 float64
	for i := 0; i < n; i += 4 {
		for j := 0; j < n; j++ {
			// Put in buffer row of matrix [A]
			buffer0[j] = (*A)[i+0][j]
			buffer1[j] = (*A)[i+1][j]
			buffer2[j] = (*A)[i+2][j]
			buffer3[j] = (*A)[i+3][j]
		}
		for j := 0; j < n; j++ {
			sum0 = 0.0
			sum1 = 0.0
			sum2 = 0.0
			sum3 = 0.0
			for k := 0; k < n; k++ {
				sum0 += buffer0[k] * (*B)[k][j]
				sum1 += buffer1[k] * (*B)[k][j]
				sum2 += buffer2[k] * (*B)[k][j]
				sum3 += buffer3[k] * (*B)[k][j]
			}
			(*C)[i+0][j] = sum0
			(*C)[i+1][j] = sum1
			(*C)[i+2][j] = sum2
			(*C)[i+3][j] = sum3
		}
	}
}
