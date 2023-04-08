//go:build go1.20

package gwrap_test

import (
	"testing"

	"github.com/muir/gwrap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompareMap(t *testing.T) {
	var m gwrap.CompareMap[string, int]

	total := func() int {
		var total int
		m.Range(func(_ string, v int) bool {
			total += v
			return true
		})
		return total
	}

	m.Store("foo", 2)
	require.Equal(t, 2, total())
	m.Store("bar", 8)
	require.Equal(t, 10, total())

	ok := m.CompareAndSwap("bar", 11, 100)
	assert.False(t, ok)
	require.Equal(t, 10, total())

	ok = m.CompareAndSwap("bar", 8, 98)
	assert.True(t, ok)
	require.Equal(t, 100, total())

	ok = m.CompareAndDelete("bar", 99)
	assert.False(t, ok)
	require.Equal(t, 100, total())

	ok = m.CompareAndDelete("bar", 98)
	assert.True(t, ok)
	require.Equal(t, 2, total())
}

func TestSwap(t *testing.T) {
	var m gwrap.SyncMap[string, int]

	total := func() int {
		var total int
		m.Range(func(_ string, v int) bool {
			total += v
			return true
		})
		return total
	}

	m.Store("foo", 7)
	require.Equal(t, 7, total())
	m.Store("bar", 10)
	require.Equal(t, 17, total())

	p, ok := m.Swap("bar", 8)
	assert.True(t, ok)
	assert.Equal(t, 10, p)
	require.Equal(t, 15, total())

	p, ok = m.Swap("baz", -5)
	assert.False(t, ok)
	assert.Equal(t, 0, p)
	require.Equal(t, 10, total())
}
