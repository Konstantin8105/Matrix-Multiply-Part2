# Matrix Multiply Part2

Let's continue.
Now, we compare "GNU GSL + Cgo vs Parallel Buffer" algorithms.
```command
BenchmarkBuffer-4   	     200	 397804236 ns/op
BenchmarkGSL-4      	      50	1692919032 ns/op
```

Now, we add to comparing "OpenBLAS" algorithm.
```command
BenchmarkBLAS-4     	    1000	  89582252 ns/op
BenchmarkBuffer-4   	     200	 397804236 ns/op
BenchmarkGSL-4      	      50	1692919032 ns/op
```
Note: [OpenBLAS](https://github.com/xianyi/OpenBLAS) little bit unstable.

#### TODO

* cgo for http://eigen.tuxfamily.org/index.php?title=Benchmark
* Strassen's algorithm O^2.8074 (~3.8 times different for 1024 elements)
