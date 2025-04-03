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
