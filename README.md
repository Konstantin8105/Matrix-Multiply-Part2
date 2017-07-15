# Matrix Multiply Part2

Let's continue.
Now, we compare "GNU GSL + Cgo vs Parallel Buffer" algorithms.
```command
BenchmarkBuffer-4   	     200	 397804236 ns/op
BenchmarkGSL-4      	      50	1692919032 ns/op
```

Now, we add to comparing "OpenBLAS" algorithm.
```command
BenchmarkBuffer-4   	     300	 377738429 ns/op
BenchmarkGSL-4      	      50	1664413660 ns/op
BenchmarkBLAS-4     	     100	1249711011 ns/op
```
Note about [OpenBLAS](https://github.com/xianyi/OpenBLAS):
* little bit unstable
* function `cblas_dgemm` use `const int` insteand of `int`


#### TODO

* Strassen's algorithm O^2.8074 (~3.8 times different for 1024 elements)
