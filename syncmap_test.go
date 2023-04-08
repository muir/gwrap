package gwrap_test

import (
	"testing"

	"github.com/muir/gwrap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSyncMap(t *testing.T) {
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
	m.Store("baz", -5)
	require.Equal(t, 12, total())

	v, ok := m.LoadOrStore("baz", 5)
	assert.True(t, ok)
	assert.Equal(t, -5, v)
	require.Equal(t, 12, total())

	v, ok = m.LoadOrStore("bloop", 10)
	assert.False(t, ok)
	assert.Equal(t, 10, v)
	require.Equal(t, 22, total())

	v, ok = m.LoadAndDelete("bloop")
	assert.True(t, ok)
	assert.Equal(t, 10, v)
	require.Equal(t, 12, total())

	v, ok = m.LoadAndDelete("blarp")
	assert.False(t, ok)
	assert.Equal(t, 0, v)
	require.Equal(t, 12, total())

	v, ok = m.Load("blarp")
	assert.False(t, ok)
	assert.Equal(t, 0, v)
	require.Equal(t, 12, total())

	v, ok = m.Load("foo")
	assert.True(t, ok)
	assert.Equal(t, 7, v)
	require.Equal(t, 12, total())

	m.Delete("foo")
	require.Equal(t, 5, total())
}
