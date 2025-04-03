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

	m.Store("abc", 9)
	assert.False(t, m.CompareAndSwap("abc", 10, 8))
	v, ok = m.Load("abc")
	assert.True(t, ok)
	assert.Equal(t, 9, v)
	assert.True(t, m.CompareAndSwap("abc", 9, 8))
	v, ok = m.Load("abc")
	assert.True(t, ok)
	assert.Equal(t, 8, v)

	t.Log("now testing invalid values inserted by bypassing the API")
	m.Map.Store("invalid", "comparable")
	_, ok = m.Load("invalid")
	require.False(t, ok)
	t.Log("seems like it isn't there")
	_, ok = m.Map.Load("invalid")
	require.True(t, ok)
	t.Log("but it is")
	_, ok = m.LoadAndDelete("invalid")
	assert.False(t, ok)
	t.Log("seems like it wasn't there")
	_, ok = m.Map.Load("invalid")
	require.False(t, ok)
	t.Log("now it's really gone")

	t.Log("now we replace an invalid type value that's comparable")
	m.Map.Store("invalid", "comparable")
	_, ok = m.Map.Load("invalid")
	require.True(t, ok)
	v, ok = m.LoadOrStore("invalid", 22)
	require.False(t, ok)
	require.Equal(t, 22, v)
	v, ok = m.Load("invalid")
	require.True(t, ok)
	require.Equal(t, 22, v)

	t.Log("now we replace an invalid type value that's not comparable")
	var blankAny any
	m.Map.Store("invalid", blankAny)
	_, ok = m.Map.Load("invalid")
	require.True(t, ok)
	v, ok = m.LoadOrStore("invalid", 23)
	require.False(t, ok)
	require.Equal(t, 23, v)
	v, ok = m.Load("invalid")
	require.True(t, ok)
	require.Equal(t, 23, v)
}
