package main

/*
#include <stdio.h>
#include <stdlib.h>
#include "/home/konstantin/Programs/OpenBLAS2/include/cblas.h"
#cgo LDFLAGS: -I /home/konstantin/Programs/OpenBLAS2/include/ -L/home/konstantin/Programs/OpenBLAS2/lib/ -lopenblas -lpthread -lgfortran
*/
import "C"

func mmOpenBLAS(A, B, D *[]float64) {
	const (
		n = 1024
	)
	a := *A
	b := *B
	c := *D

	//C.cblas_dgemm(C.CblasColMajor, C.CblasNoTrans, C.CblasTrans, n, n, n, C.double(1.0), mA, n, mB, n, C.double(1.0), mC, n)
	C.cblas_dgemm(C.CblasColMajor, C.CblasNoTrans, C.CblasTrans, n, n, n, C.double(1.0), (*C.double)(&a[0]), n, (*C.double)(&b[0]), n, C.double(1.0), (*C.double)(&c[0]), n)
}

/*
func main() {
	//	C.cblas_dgemm(C.CblasRowMajor, C.CblasTrans, C.CblasNoTrans, C.INVALID, 0, 0,
	//		ALPHA, A, 1, B, 1, BETA, C, 1)
	A := [6]float64{1.0, 2.0, 1.0, -3.0, 4.0, -1.0}
	B := [6]float64{1.0, 2.0, 1.0, -3.0, 4.0, -1.0}
	D := [9]float64{.5, .5, .5, .5, .5, .5, .5, .5, .5}
	C.cblas_dgemm(C.CblasColMajor, C.CblasNoTrans, C.CblasTrans, 3, 3, 2, 1, (*C.double)(&A[0]), 3, (*C.double)(&B[0]), 3, 2, (*C.double)(&D[0]), 3)
	for i := 0; i < 9; i++ {
		fmt.Println(D[i])
	}
}
*/
