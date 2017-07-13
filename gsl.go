package main

/*
#include <stdlib.h>
#include <gsl/gsl_errno.h>
#include <gsl/gsl_sf_bessel.h>
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
}
