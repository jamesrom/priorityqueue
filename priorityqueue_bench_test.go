package priorityqueue_test

import (
	"testing"

	"github.com/jamesrom/priorityqueue"
)

func benchmarkPushLinearIntAscending(b *testing.B, size int) {
	q := priorityqueue.NewWithComparator(func(a, b int) bool { return a > b })
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			q.Push(i)
		}
	}
}

func benchmarkPushLinearIntDescending(b *testing.B, size int) {
	q := priorityqueue.NewWithComparator(func(a, b int) bool { return a > b })
	for n := 0; n < b.N; n++ {
		for i := size; i > 0; i-- {
			q.Push(i)
		}
	}
}

// Push `size` integers into the queue, in ascending order. This simulates the
// worst-case: each element pushed is the highest priority and must swim to the top.
func BenchmarkPush10Ascending(b *testing.B)      { benchmarkPushLinearIntAscending(b, 10) }
func BenchmarkPush100Ascending(b *testing.B)     { benchmarkPushLinearIntAscending(b, 100) }
func BenchmarkPush1000Ascending(b *testing.B)    { benchmarkPushLinearIntAscending(b, 1000) }
func BenchmarkPush10000Ascending(b *testing.B)   { benchmarkPushLinearIntAscending(b, 10000) }
func BenchmarkPush100000Ascending(b *testing.B)  { benchmarkPushLinearIntAscending(b, 100000) }
func BenchmarkPush1000000Ascending(b *testing.B) { benchmarkPushLinearIntAscending(b, 1000000) }

// Push `size` integers into the queue, in descending order. This simulates the best case
// of an early-exit in the swim up process.
func BenchmarkPush10Descending(b *testing.B)      { benchmarkPushLinearIntDescending(b, 10) }
func BenchmarkPush100Descending(b *testing.B)     { benchmarkPushLinearIntDescending(b, 100) }
func BenchmarkPush1000Descending(b *testing.B)    { benchmarkPushLinearIntDescending(b, 1000) }
func BenchmarkPush10000Descending(b *testing.B)   { benchmarkPushLinearIntDescending(b, 10000) }
func BenchmarkPush100000Descending(b *testing.B)  { benchmarkPushLinearIntDescending(b, 100000) }
func BenchmarkPush1000000Descending(b *testing.B) { benchmarkPushLinearIntDescending(b, 1000000) }
