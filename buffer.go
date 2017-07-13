package main

import (
	"math"
	"runtime"
	"sync"
)

func mmBuffer(A, B, C *[]float64) {
	n := int(math.Sqrt(float64(len(*A))))
	// Found amount allowable parallelism
	threads := runtime.GOMAXPROCS(0)
	if threads > runtime.NumCPU() {
		threads = runtime.NumCPU()
	}
	// Create workgroup
	var wg sync.WaitGroup
	// Run calculation in goroutines
	for t := 0; t < threads; t++ {
		// Add one goroutine in workgroup
		wg.Add(1)
		// The value "init" is a number of thread
		// that created for offset of loop
		go func(init int) {
			// Change waitgroup after work done
			defer wg.Done()
			// Inialize addition variables
			var sum00, sum01, sum02, sum03 float64
			// Create buffers
			amountBuffers := 4
			buffer0 := make([]float64, n, n)
			buffer1 := make([]float64, n, n)
			buffer2 := make([]float64, n, n)
			buffer3 := make([]float64, n, n)
			// Calculate amount of calculation part
			// for that goroutine
			amountParts := n / amountBuffers
			for i := init; i < amountParts; i += threads {
				for j := 0; j < n; j++ {
					// Put in buffer row of matrix [A]
					buffer0[j] = (*A)[(i*amountBuffers+0)+j*n]
					buffer1[j] = (*A)[(i*amountBuffers+1)+j*n]
					buffer2[j] = (*A)[(i*amountBuffers+2)+j*n]
					buffer3[j] = (*A)[(i*amountBuffers+3)+j*n]
				}
				for j := 0; j < n; j++ {
					sum00 = 0.0
					sum01 = 0.0
					sum02 = 0.0
					sum03 = 0.0
					b := (*B)[j*n : j*n+n]
					for k := 0; k < n; k++ {
						sum00 += buffer0[k] * b[k]
						sum01 += buffer1[k] * b[k]
						sum02 += buffer2[k] * b[k]
						sum03 += buffer3[k] * b[k]

					}
					(*C)[(i*amountBuffers+0)+j*n] = sum00
					(*C)[(i*amountBuffers+1)+j*n] = sum01
					(*C)[(i*amountBuffers+2)+j*n] = sum02
					(*C)[(i*amountBuffers+3)+j*n] = sum03
				}
			}
		}(t)
	}
	wg.Wait()
}
