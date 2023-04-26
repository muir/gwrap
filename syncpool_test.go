package gwrap_test

import (
	"testing"

	"github.com/muir/gwrap"

	"github.com/stretchr/testify/assert"
)

func TestSyncPool(t *testing.T) {
	type foo struct {
		a int
	}

	var x gwrap.SyncPool[foo]

	assert.Equal(t, foo{}, x.Get())
	for i := 1; i < 10000; i++ {
		x.Put(foo{a: i})
	}
	var count int
	for i := 1; i < 10000; i++ {
		xi := x.Get()
		if xi.a != 0 {
			count++
		}
	}
	assert.NotEmpty(t, count)

	y := gwrap.NewSyncPool(func() []string {
		return make([]string, 0, 4)
	})

	assert.Equal(t, 4, cap(y.Get()))

	z := gwrap.SyncPool[[]int]{
		New: func() []int { return make([]int, 3) },
	}

	assert.Equal(t, 3, len(z.Get()))
}
