package main

/*
#include <stdio.h>
#include <stdlib.h>
#include "/home/konstantin/Programs/OpenBLAS2/include/cblas.h"
#cgo LDFLAGS: -I /home/konstantin/Programs/OpenBLAS2/include/ -L/home/konstantin/Programs/OpenBLAS2/lib/ -lopenblas -lpthread -lgfortran
*/
import "C"

/*
func main() {
	//	C.cblas_dgemm(C.CblasRowMajor, C.CblasTrans, C.CblasNoTrans, C.INVALID, 0, 0,
	//		ALPHA, A, 1, B, 1, BETA, C, 1)
	/*
	   int i=0;
	   double A[6] = {1.0,2.0,1.0,-3.0,4.0,-1.0};
	   double B[6] = {1.0,2.0,1.0,-3.0,4.0,-1.0};
	   double C[9] = {.5,.5,.5,.5,.5,.5,.5,.5,.5};
	   cblas_dgemm(CblasColMajor, CblasNoTrans, CblasTrans,3,3,2,1,A, 3, B, 3,2,C,3);

	   for(i=0; i<9; i++)
	     printf("%lf ", C[i]);
	   printf("\n");
*/ /*
	A := [6]float64{1.0, 2.0, 1.0, -3.0, 4.0, -1.0}
	B := [6]float64{1.0, 2.0, 1.0, -3.0, 4.0, -1.0}
	C := [9]float64{.5, .5, .5, .5, .5, .5, .5, .5, .5}
	C.cblas_dgemm(C.CblasColMajor, C.CblasNoTrans, C.CblasTrans, 3, 3, 2, 1, (*C.double)(&A[0]), 3, (*C.double)(&B[0]), 3, 2, (*C.double)(&C[0]), 3)
	for i := 0; i < 9; i++ {
		fmt.Println(C[i])
	}
}*/

func mmBLAS(A, B, D *[][]float64) {
	n := C.blasint(len(*A))
	C.cblas_dgemm(C.CblasColMajor, C.CblasNoTrans, C.CblasTrans, n, n, n, C.double(1.0), (*C.double)(&((*A)[0][0])), n, (*C.double)(&((*B)[0][0])), n, C.double(1.0), (*C.double)(&((*D)[0][0])), n)
}
