# priorityqueue
A fast, generic, thread safe priority queue for Go 1.18+.

## Features

 - Minimal API
 - Generic initializers – anything that can be ordered can be put into the queue.
 - Built-in thread safety
 - Speed – see benchmarks

## TODO

 - [ ] Documentation
 - [x] Benchmarks
 - [ ] Thread safety tests

## Under the hood
The priority queue implements a binary heap (using `container/heap` from the standard library).

### Benchmarks

```sh
goos: windows
goarch: amd64
pkg: github.com/jamesrom/priorityqueue
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkPushMaxOntoSize10-8          	 8987396	       142.6 ns/op	      51 B/op	       0 allocs/op
BenchmarkPushMaxOntoSize100-8         	 9521926	       139.3 ns/op	      49 B/op	       0 allocs/op
BenchmarkPushMaxOntoSize1000-8        	 9483894	       136.5 ns/op	      49 B/op	       1 allocs/op
BenchmarkPushMaxOntoSize10000-8       	 9374215	       137.5 ns/op	      49 B/op	       1 allocs/op
BenchmarkPushMaxOntoSize100000-8      	 9338542	       137.8 ns/op	      49 B/op	       1 allocs/op
BenchmarkPushMaxOntoSize1000000-8     	 9193189	       144.9 ns/op	      56 B/op	       1 allocs/op
BenchmarkPushMaxOntoSize10000000-8    	 8510607	       147.6 ns/op	      40 B/op	       1 allocs/op
BenchmarkPushMaxOntoSize100000000-8   	 7896320	       151.7 ns/op	       8 B/op	       1 allocs/op
BenchmarkPushMinOntoSize10-8          	26356190	        42.10 ns/op	      53 B/op	       1 allocs/op
BenchmarkPushMinOntoSize100-8         	26087636	        42.51 ns/op	      54 B/op	       1 allocs/op
BenchmarkPushMinOntoSize1000-8        	26652747	        42.06 ns/op	      53 B/op	       1 allocs/op
BenchmarkPushMinOntoSize10000-8       	25247264	        42.34 ns/op	      55 B/op	       1 allocs/op
BenchmarkPushMinOntoSize100000-8      	26076072	        42.03 ns/op	      53 B/op	       1 allocs/op
BenchmarkPushMinOntoSize1000000-8     	27891862	        41.44 ns/op	      49 B/op	       1 allocs/op
BenchmarkPushMinOntoSize10000000-8    	31578614	        41.45 ns/op	      51 B/op	       1 allocs/op
BenchmarkPushMinOntoSize100000000-8   	31583269	        41.27 ns/op	      44 B/op	       1 allocs/op
BenchmarkPopFromSize10-8              	 6015972	       204.1 ns/op	       7 B/op	       0 allocs/op
BenchmarkPopFromSize100-8             	 6779817	       220.7 ns/op	       7 B/op	       0 allocs/op
BenchmarkPopFromSize1000-8            	 6266318	       216.6 ns/op	       8 B/op	       1 allocs/op
BenchmarkPopFromSize10000-8           	 6382936	       207.3 ns/op	       8 B/op	       1 allocs/op
BenchmarkPopFromSize100000-8          	 6611161	       210.2 ns/op	       8 B/op	       1 allocs/op
BenchmarkPopFromSize1000000-8         	 6250816	       211.9 ns/op	       8 B/op	       1 allocs/op
BenchmarkPopFromSize10000000-8        	 5357144	       220.8 ns/op	       8 B/op	       1 allocs/op
BenchmarkPopFromSize100000000-8       	 4819297	       267.4 ns/op	       8 B/op	       1 allocs/op
BenchmarkPeekFromSize10-8             	95812207	        12.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize100-8            	92128395	        12.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize1000-8           	92311952	        12.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize10000-8          	92383020	        12.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize100000-8         	92302011	        12.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize1000000-8        	95803792	        12.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize10000000-8       	92275749	        12.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize100000000-8      	92310531	        12.88 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 69.0% of statements
ok  	github.com/jamesrom/priorityqueue	390.646s
```
