package gwrap

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

// PQItemEmbed may be embeded in the priority queue item's data
// structure, or the data structure can have a PQItemEmbed that
// it returns when the PQItem() method is called.
type PQItemEmbed[P constraints.Ordered] struct {
	priority P
	index    int // subtract 1 to get actual index so zero is invalid
}

type PQItem[P constraints.Ordered] interface {
	PQItem() *PQItemEmbed[P]
}

type innerPriorityQueue[P constraints.Ordered, T PQItem[P]] struct {
	slice []T
}

// PriorityQueue is implemented with container/heap and is not thread-safe.
type PriorityQueue[P constraints.Ordered, T PQItem[P]] struct {
	i innerPriorityQueue[P, T]
}

// NewPriorityQueue creates a PriorityQueue from a comparison function.
func NewPriorityQueue[P constraints.Ordered, T PQItem[P]]() *PriorityQueue[P, T] {
	i := innerPriorityQueue[P, T]{
		slice: make([]T, 0, 100),
	}
	heap.Init(&i)
	return &PriorityQueue[P, T]{
		i: i,
	}
}

func (pqe *PQItemEmbed[P]) PQItem() *PQItemEmbed[P] { return pqe }

func (h *innerPriorityQueue[P, T]) Len() int { return len(h.slice) }
func (h *innerPriorityQueue[P, T]) Less(i, j int) bool {
	return h.slice[i].PQItem().priority < h.slice[j].PQItem().priority
}
func (h *innerPriorityQueue[P, T]) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
	h.slice[i].PQItem().index = i + 1
	h.slice[j].PQItem().index = j + 1
}

// Push is O(log n)
func (h *PriorityQueue[P, T]) Enqueue(x T, priority P) {
	if x.PQItem().index != 0 {
		panic("cannot add an item to a priority queue if it is already in the queue")
	}
	x.PQItem().priority = priority
	heap.Push(&h.i, x)
}

func (h *innerPriorityQueue[P, T]) Push(x any) {
	t := x.(T)
	t.PQItem().index = len(h.slice) + 1
	h.slice = append(h.slice, t)
}

// Dequeue is O(log n)
func (h *PriorityQueue[P, T]) Dequeue() T {
	return heap.Pop(&h.i).(T)
}

func (h *innerPriorityQueue[P, T]) Pop() any {
	old := h.slice
	n := len(old)
	x := old[n-1]
	x.PQItem().index = 0
	h.slice = old[0 : n-1]
	return x
}

func (h PriorityQueue[P, T]) Len() int {
	return len(h.i.slice)
}

// Remove takes an item out of the priority queue
func (h *PriorityQueue[P, T]) Remove(x T) {
	i := x.PQItem().index
	if i == 0 {
		panic("cannot remove item from priority queue that is not present in queue")
	}
	x.PQItem().index = 0
	heap.Remove(&h.i, i-1)
}
