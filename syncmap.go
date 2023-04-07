package gwrap

import "sync"

type SyncMap[K comparable, V any] struct {
	sync.Map
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	var zero V
	var v any
	v, ok = m.Map.Load(key)
	if ok {
		return v.(V), true
	}
	return zero, false
}

func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	var v any
	v, loaded = m.Map.LoadAndDelete(key)
	if loaded {
		return v.(V), true
	}
	var zero V
	return zero, false
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	var v any
	v, loaded = m.Map.LoadOrStore(key, value)
	if loaded {
		return v.(V), true
	}
	return value, false
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(func(key any, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.Map.Store(key, value)
}

func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	var v any
	v, loaded = m.Map.Swap(key, value)
	if loaded {
		return v.(V), true
	}
	var zero V
	return zero, false
}
