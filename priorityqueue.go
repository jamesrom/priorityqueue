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

func New[T Prioritizable](items ...T) *PriorityQueue[T] {
	defaultComparator := func(i, j T) bool {
		return i.Priority() > j.Priority()
	}
	return NewWithComparator(defaultComparator, items...)
}

func NewWithOrderable[T constraints.Ordered](items ...T) *PriorityQueue[T] {
	defaultComparator := func(a, b T) bool { return a > b }
	return NewWithComparator(defaultComparator, items...)
}

func NewWithComparator[T any](fn func(T, T) bool, items ...T) *PriorityQueue[T] {
	h := heap[T]{
		elements:   items,
		comparator: fn,
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
