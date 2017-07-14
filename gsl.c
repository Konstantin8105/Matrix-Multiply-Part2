#include <time.h>
#include <stdio.h>
#include <gsl/gsl_blas.h>

int
main (void)
{
    int n =  1024;
    int size = n*n;
    double *aq = malloc (sizeof (double) * size);
    double *bq = malloc (sizeof (double) * size);
    double *cq = malloc (sizeof (double) * size);
    
    int i,j;
    
    for ( i=0;i<n;i++){
        for ( j=0;j<n;j++){
            aq[i+j*n] = 4.0 *(i+9+j);
            bq[i+j*n] = 4.0 *(11-i-2*j);
        }
    }

    gsl_matrix_view A = gsl_matrix_view_array(aq, n, n);
    gsl_matrix_view B = gsl_matrix_view_array(bq, n, n);
    gsl_matrix_view C = gsl_matrix_view_array(cq, n, n);

    /* Compute C = A B */
    clock_t begin = clock();
    gsl_blas_dgemm (CblasNoTrans, CblasNoTrans,
                  1.0, &A.matrix, &B.matrix,
                  0.0, &C.matrix);
    clock_t end = clock();
    double time_spent = (double)(end - begin) / CLOCKS_PER_SEC;

    printf (" time = %.8g sec.\n", time_spent);
    return 0;
}
