package gwrap_test

import (
	"testing"

	"github.com/muir/gwrap"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueueHappy(t *testing.T) {
	type item struct {
		gwrap.PQItemEmbed[int]
		value int
	}

	h := gwrap.NewPriorityQueue[int, *item]()
	assert.Equal(t, 0, h.Len())

	h.Enqueue(&item{value: 30}, 3)
	h.Enqueue(&item{value: 20}, 2)
	h.Enqueue(&item{value: 40}, 4)

	// 20 30 40

	assert.Equal(t, 3, h.Len())

	assert.Equal(t, 20, h.Dequeue().value)

	// 30 40

	h.Enqueue(&item{value: 60}, 6)
	h.Enqueue(&item{value: 10}, 1)
	i50 := &item{value: 50}
	h.Enqueue(i50, 5)

	// 10 30 40 50 60

	assert.Equal(t, 10, h.Dequeue().value)
	assert.Equal(t, 30, h.Dequeue().value)

	// 40 50 60

	assert.Equal(t, 3, h.Len(), "length, before remove")

	h.Remove(i50)

	// 40 60

	assert.Equal(t, 2, h.Len(), "length, post remove")

	assert.Equal(t, 40, h.Dequeue().value, "1st dequeue after remove")

	i60 := h.Dequeue()
	assert.Equal(t, 60, i60.value, "2nd dequeue after remove")

	h.Enqueue(i60, 6)

	assert.Equal(t, 1, h.Len(), "final length")
}

func TestPriorityQueuePanics(t *testing.T) {
	type item struct {
		gwrap.PQItemEmbed[int]
		value int
	}

	h := gwrap.NewPriorityQueue[int, *item]()

	h.Enqueue(&item{value: 30}, 3)
	h.Enqueue(&item{value: 20}, 2)
	h.Enqueue(&item{value: 40}, 4)

	assert.Panics(t, func() {
		h.Remove(&item{value: 90})
	}, "remove value never inserted")

	iTwice := &item{value: 50}
	h.Enqueue(iTwice, 5)

	assert.Panics(t, func() {
		h.Enqueue(iTwice, 5)
	}, "enqueue an already-queued item")
}
