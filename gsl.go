package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <gsl/gsl_errno.h>
#include <gsl/gsl_sf_bessel.h>
#include <gsl/gsl_blas.h>
#cgo LDFLAGS: -lm -lgsl -lgslcblas
*/

import "C"
import (
	"fmt"
)

func main() {

	{
		x := 5.0
		expected := -0.17759677131433830434739701
		y := C.gsl_sf_bessel_J0(C.double(x))
		fmt.Println("y        = ", y)
		fmt.Println("expected = ", expected)
	}

	{
		x := 5.0
		var result C.gsl_sf_result
		expected := -0.17759677131433830434739701
		status := C.gsl_sf_bessel_J0_e(C.double(x), &result)
		fmt.Println("status   = ", C.GoString(C.gsl_strerror(status)))
		fmt.Println("result   = ", float64(result.val), "\terror = ", float64(result.err))
		fmt.Println("expected = ", expected)
	}

	{
		// https://www.gnu.org/software/gsl/doc/html/blas.html#examples
		a := []float64{0.11, 0.12, 0.13,
			0.21, 0.22, 0.23}
		mA := C.gsl_matrix_alloc(2, 3)
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				C.gsl_matrix_set(mA, C.size_t(i), C.size_t(j), C.double(a[i*3+j]))
			}
		}

		b := []float64{1011, 1012,
			1021, 1022,
			1031, 1032}
		mB := C.gsl_matrix_alloc(3, 2)
		for i := 0; i < 3; i++ {
			for j := 0; j < 2; j++ {
				C.gsl_matrix_set(mB, C.size_t(i), C.size_t(j), C.double(b[i*2+j]))
			}
		}

		mC := C.gsl_matrix_alloc(2, 2)

		C.gsl_blas_dgemm(C.CblasNoTrans, C.CblasNoTrans, C.double(1.0), mA, mB, C.double(0.0), mC)

		fmt.Println("[", C.gsl_matrix_get(mC, 0, 0), C.gsl_matrix_get(mC, 0, 1), "]")
		fmt.Println("[", C.gsl_matrix_get(mC, 1, 0), C.gsl_matrix_get(mC, 1, 1), "]")

		C.gsl_matrix_free(mA)
		C.gsl_matrix_free(mB)
		C.gsl_matrix_free(mC)
	}
}

// TODO : remove dublicate of memory
// link   : https://askubuntu.com/questions/623339/altough-i-installed-gsl-library-g-cannot-compile-my-code
// install: sudo apt-get install libgsl0-dev
func mmGSL(A, B, D *[][]float64) {
	/*
		n := C.size_t(len(*A))

		mA := new(C.gsl_matrix)
		a := &((*A)[0][0])
		mA.data = (*C.double)((a)) // (*C.double)(unsafe.Pointer(&a)) //(&((*A)[0]))
		mA.size1 = n
		mA.size2 = n
		mA.tda = n
		mA.block = nil
		mA.owner = C.int(0)

		mB := new(C.gsl_matrix)
		b := &((*B)[0][0])
		mA.data = (*C.double)((b)) //(*C.double)(unsafe.Pointer(&b)) //(&((*A)[0]))
		mB.size1 = n
		mB.size2 = n
		mB.tda = n
		mB.block = nil
		mB.owner = C.int(0)

		mC := new(C.gsl_matrix)
		c := &((*D)[0][0])
		mA.data = (*C.double)((c)) //(*C.double)(unsafe.Pointer(&c)) //(&((*A)[0]))
		mC.size1 = n
		mC.size2 = n
		mC.tda = n
		mC.block = nil
		mC.owner = C.int(0)

		C.gsl_blas_dgemm(C.CblasNoTrans, C.CblasNoTrans, C.double(1.0), mA, mB, C.double(0.0), mC)
	*/

	n := C.size_t(len(*A))
	mA := C.gsl_matrix_alloc(n, n)
	for i := 0; C.size_t(i) < C.size_t(n); i++ {
		for j := 0; C.size_t(j) < C.size_t(n); j++ {
			C.gsl_matrix_set(mA, C.size_t(i), C.size_t(j), C.double((*A)[i][j]))
		}
	}

	mB := C.gsl_matrix_alloc(n, n)
	for i := 0; C.size_t(i) < C.size_t(n); i++ {
		for j := 0; C.size_t(j) < C.size_t(n); j++ {
			C.gsl_matrix_set(mB, C.size_t(i), C.size_t(j), C.double((*B)[i][j]))
		}
	}

	mC := C.gsl_matrix_alloc(n, n)

	C.gsl_blas_dgemm(C.CblasNoTrans, C.CblasNoTrans, C.double(1.0), mA, mB, C.double(0.0), mC)

	for i := 0; C.size_t(i) < C.size_t(n); i++ {
		for j := 0; C.size_t(j) < C.size_t(n); j++ {
			(*D)[i][j] = float64(C.gsl_matrix_get(mC, C.size_t(i), C.size_t(j)))
		}
	}

	C.gsl_matrix_free(mA)
	C.gsl_matrix_free(mB)
	C.gsl_matrix_free(mC)

}
