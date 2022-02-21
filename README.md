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

goos: windows
goarch: amd64
pkg: github.com/jamesrom/priorityqueue
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkPush10Ascending-8         	 2580646	       441.2 ns/op	     465 B/op	       0 allocs/op
BenchmarkPush100Ascending-8        	  266526	      4356 ns/op	    4511 B/op	       0 allocs/op
BenchmarkPush1000Ascending-8       	   23054	     50782 ns/op	   47667 B/op	     744 allocs/op
BenchmarkPush10000Ascending-8      	    2086	    541466 ns/op	  538980 B/op	    9744 allocs/op
BenchmarkPush100000Ascending-8     	     205	   5680488 ns/op	 5489199 B/op	   99744 allocs/op
BenchmarkPush1000000Ascending-8    	      16	  62898462 ns/op	56073183 B/op	  999747 allocs/op
BenchmarkPush10Descending-8        	 3041575	       412.6 ns/op	     494 B/op	       0 allocs/op
BenchmarkPush100Descending-8       	  303991	      3980 ns/op	    4944 B/op	       0 allocs/op
BenchmarkPush1000Descending-8      	   25288	     45794 ns/op	   53505 B/op	     745 allocs/op
BenchmarkPush10000Descending-8     	    2474	    466319 ns/op	  563949 B/op	    9745 allocs/op
BenchmarkPush100000Descending-8    	     252	   4640873 ns/op	 5569142 B/op	   99745 allocs/op
BenchmarkPush1000000Descending-8   	      25	  46041216 ns/op	56091455 B/op	  999747 allocs/op
PASS
coverage: 34.5% of statements
ok  	github.com/jamesrom/priorityqueue	17.965s
