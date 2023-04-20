package gwrap_test

import (
	"testing"

	"github.com/muir/gwrap"

	"github.com/stretchr/testify/assert"
)

func TestAtomicValue(t *testing.T) {
	var i gwrap.AtomicValue[int]
	var sp gwrap.AtomicValue[*string]
	var ss gwrap.AtomicValue[[]string]

	ss.Store([]string{"a", "b"})
	assert.Equal(t, []string{"a", "b"}, ss.Load())

	assert.Equal(t, int(0), i.Load())
	assert.Equal(t, (*string)(nil), sp.Load())

	i.Store(7)
	assert.Equal(t, int(7), i.Swap(8))
	assert.Equal(t, int(8), i.Load())
	assert.False(t, i.CompareAndSwap(10, 11))
	assert.Equal(t, int(8), i.Load())
	assert.True(t, i.CompareAndSwap(8, 11))
	assert.Equal(t, int(11), i.Load())

	foo := "foo"
	bar := "bar"
	baz := "baz"
	assert.Equal(t, (*string)(nil), sp.Swap(&foo))
	assert.Equal(t, &foo, sp.Load())
	assert.False(t, sp.CompareAndSwap(&baz, &bar))
	assert.True(t, sp.CompareAndSwap(&foo, &bar))
	assert.Equal(t, &bar, sp.Swap(nil))
}
