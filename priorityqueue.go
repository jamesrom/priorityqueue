package priorityqueue

import (
	"constraints"
	hp "container/heap"
	"sync"
)

type PriorityQueue[T any] struct {
	h   heap[T]
	mtx sync.RWMutex
}

type Prioritizable interface {
	Priority() int
}

// New - highest priority / descending order
func New[T Prioritizable](items ...T) *PriorityQueue[T] {
	greater := func(i, j T) bool {
		return i.Priority() > j.Priority()
	}
	return NewWithComparator(greater, items...)
}

// NewWithOrderable - highest priority / descending order
func NewWithOrderable[T constraints.Ordered](items ...T) *PriorityQueue[T] {
	greater := func(a, b T) bool { return a > b }
	return NewWithComparator(greater, items...)
}

// Less must describe a transitive ordering:
//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
type Less[T any] func(T, T) bool

// NewWithComparator - order defined by a custom less function
func NewWithComparator[T any](fn Less[T], items ...T) *PriorityQueue[T] {
	h := heap[T]{
		elements: items,
		less:     fn,
	}
	hp.Init(&h)
	return &PriorityQueue[T]{h: h}
}

func (p *PriorityQueue[T]) Push(x T) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	hp.Push(&p.h, x)
}

func (p *PriorityQueue[T]) Pop() T {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	return hp.Pop(&p.h).(T)
}

func (p *PriorityQueue[T]) Peek() T {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return p.h.elements[0]
}

func (p *PriorityQueue[T]) Len() int {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return len(p.h.elements)
}
