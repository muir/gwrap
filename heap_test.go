package gwrap_test

import (
	"testing"

	"github.com/muir/gwrap"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := gwrap.NewHeap(func(a float64, b float64) bool {
		return a < b
	})
	assert.Equal(t, 0, h.Len())
	h.Push(1.7)
	h.Push(5.9)
	h.Push(3.4)
	assert.Equal(t, 3, h.Len())
	assert.Equal(t, 1.7, h.Pop())
	assert.Equal(t, 2, h.Len())
	h.Push(2.7)
	h.Push(4.2)
	assert.Equal(t, 2.7, h.Pop())
	assert.Equal(t, 3.4, h.Pop())
	assert.Equal(t, 4.2, h.Pop())
	assert.Equal(t, 5.9, h.Pop())
	assert.Equal(t, 0, h.Len())
}
