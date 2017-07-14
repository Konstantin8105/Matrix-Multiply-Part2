package main

import "C"

/*
#include <stdio.h>
#include <stdlib.h>
#include "/home/konstantin/Programs/OpenBLAS2/include/cblas.h"
#cgo LDFLAGS: -I /home/konstantin/Programs/OpenBLAS2/include/ -L/home/konstantin/Programs/OpenBLAS2/lib/ -lopenblas -lpthread -lgfortran
*/

func mmOpenBLAS(A, B, D *[][]float64) {
	n := C.blasint(len(*A))
	C.cblas_dgemm(C.CblasColMajor, C.CblasNoTrans, C.CblasTrans, n, n, n, C.double(1.0), (*C.double)(&((*A)[0][0])), n, (*C.double)(&((*B)[0][0])), n, C.double(1.0), (*C.double)(&((*D)[0][0])), n)
}
