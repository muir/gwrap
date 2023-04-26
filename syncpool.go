package gwrap

import (
	"sync"
)

type SyncPool[T any] struct {
	Pool sync.Pool
	New  func() T
}

func NewSyncPool[T any](new func() T) *SyncPool[T] {
	return &SyncPool[T]{
		New: new,
	}
}

func (p *SyncPool[T]) Get() T {
	i := p.Pool.Get()
	if i == nil {
		if p.New != nil {
			return p.New()
		}
		var zero T
		return zero
	}
	return i.(T)
}

func (p *SyncPool[T]) Put(i T) {
	p.Pool.Put(i)
}
