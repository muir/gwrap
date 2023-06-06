package gwrap

import (
	"container/heap"
)

type innerHeap[T any] struct {
	slice []T
	less  func(T, T) bool
}

// Heap is implemented with container/heap and is not thread-safe
type Heap[T any] struct {
	i innerHeap[T]
}

// NewHeap creates a Heap from a comparison function.
func NewHeap[T any](less func(a T, b T) bool) *Heap[T] {
	i := innerHeap[T]{
		slice: make([]T, 0, 100),
		less:  less,
	}
	heap.Init(&i)
	return &Heap[T]{
		i: i,
	}
}

// Code below originated with the container/heap documentation

func (h *innerHeap[T]) Len() int           { return len(h.slice) }
func (h *innerHeap[T]) Less(i, j int) bool { return h.less(h.slice[i], h.slice[j]) }
func (h *innerHeap[T]) Swap(i, j int)      { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }

// Push is O(log n)
func (h *Heap[T]) Push(x T) {
	heap.Push(&h.i, x)
}

func (h *innerHeap[T]) Push(x any) {
	h.slice = append(h.slice, x.(T))
}

// Pop is O(log n)
func (h *Heap[T]) Pop() T {
	return heap.Pop(&h.i).(T)
}

func (h *innerHeap[T]) Pop() any {
	old := h.slice
	n := len(old)
	x := old[n-1]
	h.slice = old[0 : n-1]
	return x
}

func (h Heap[T]) Len() int {
	return len(h.i.slice)
}
