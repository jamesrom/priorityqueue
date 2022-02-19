# priorityqueue
A fast, generic, thread safe priority queue for Go 1.18+.

## Features

 - Minimal API.
 - Generic initializers. Anything that can be ordered can be put into the queue.
 - Built-in thread safety.
 - Speed. See benchmarks.

## TODO

 - [ ] Documentation
 - [ ] Benchmarks
 - [ ] Thread safety tests

## Under the hood
The priority queue implements a binary heap (using `container/heap` from the standard library).
