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

	m.Store("foo", 7)
	require.Equal(t, 7, total())
	m.Store("bar", 8)
	require.Equal(t, 15, total())

	ok := m.CompareAndSwap("bar", 11, 100)
	assert.False(t, ok)
	require.Equal(t, 15, total())

	ok = m.CompareAndSwap("bar", 8, 98)
	assert.True(t, ok)
	require.Equal(t, 105, total())

	ok = m.CompareAndDelete("bar", 99)
	assert.False(t, ok)
	require.Equal(t, 105, total())

	ok = m.CompareAndDelete("bar", 98)
	assert.True(t, ok)
	require.Equal(t, 7, total())
}
