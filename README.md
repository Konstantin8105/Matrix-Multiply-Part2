# Matrix Multiply Part2

Let's continue.
Now, we compare "GNU GSL + Cgo vs Parallel Buffer" algorithms.
```command
BenchmarkBuffer-4   	     200	 447207595 ns/op
BenchmarkGSL-4      	      50	1606703441 ns/op
PASS
ok  	github.com/Konstantin8105/Matrix-Multiply-Part2	215.910s
```


#### TODO

* cgo for http://eigen.tuxfamily.org/index.php?title=Benchmark
* Strassen's algorithm O^2.8074 (~3.8 times different for 1024 elements)
