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
		/*
		  double x = 5.0;
		  double expected = -0.17759677131433830434739701;

		  double y = gsl_sf_bessel_J0 (x);

		  printf ("J0(5.0) = %.18f\n", y);
		  printf ("exact   = %.18f\n", expected);
		*/
		x := 5.0
		expected := -0.17759677131433830434739701
		y := C.gsl_sf_bessel_J0(C.double(x))
		fmt.Println("y        = ", y)
		fmt.Println("expected = ", expected)
	}
	{
		/*
		  double x = 5.0;
		  gsl_sf_result result;

		  double expected = -0.17759677131433830434739701;

		  int status = gsl_sf_bessel_J0_e (x, &result);

		  printf ("status  = %s\n", gsl_strerror(status));
		  printf ("J0(5.0) = %.18f\n"
		          "      +/- % .18f\n",
		          result.val, result.err);
		  printf ("exact   = %.18f\n", expected);
		*/
		x := 5.0
		var result C.gsl_sf_result
		expected := -0.17759677131433830434739701
		status := C.gsl_sf_bessel_J0_e(C.double(x), &result)
		fmt.Println("status   = ", C.GoString(C.gsl_strerror(status)))
		fmt.Println("result   = ", float64(result.val), "\terror = ", float64(result.err))
		fmt.Println("expected = ", expected)
	}
	{
		/*
		  double a[] = { 0.11, 0.12, 0.13,
		                 0.21, 0.22, 0.23 };

		  double b[] = { 1011, 1012,
		                 1021, 1022,
		                 1031, 1032 };

		  double c[] = { 0.00, 0.00,
		                 0.00, 0.00 };

		  gsl_matrix_view A = gsl_matrix_view_array(a, 2, 3);
		  gsl_matrix_view B = gsl_matrix_view_array(b, 3, 2);
		  gsl_matrix_view C = gsl_matrix_view_array(c, 2, 2);

		  // Compute C = A B

		  gsl_blas_dgemm (CblasNoTrans, CblasNoTrans,
		                  1.0, &A.matrix, &B.matrix,
		                  0.0, &C.matrix);

		  printf ("[ %g, %g\n", c[0], c[1]);
		  printf ("  %g, %g ]\n", c[2], c[3]);
		*/
		a := []float64{0.11, 0.12, 0.13,
			0.21, 0.22, 0.23}

		b := []float64{1011, 1012,
			1021, 1022,
			1031, 1032}

		c := []float64{0.00, 0.00,
			0.00, 0.00}

		var Am C.gsl_matrix_view = C.gsl_matrix_view_array((*C.double)(&a[0]), 2, 3)
		var Bm C.gsl_matrix_view = C.gsl_matrix_view_array((*C.double)(&b[0]), 3, 2)
		var Cm C.gsl_matrix_view = C.gsl_matrix_view_array((*C.double)(&c[0]), 2, 2)

		// before build execute:
		// export GODEBUG=cgocheck=0
		C.gsl_blas_dgemm(C.CblasNoTrans, C.CblasNoTrans,
			1.0, &Am.matrix, &Bm.matrix,
			0.0, &Cm.matrix)
		fmt.Println("[", c[0], ",", c[1], "]")
		fmt.Println("[", c[2], ",", c[3], "]")
	}
}
