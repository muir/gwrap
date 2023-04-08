package gwrap

import "sync"

// SyncMap is a wrapper of sync.Map. It uses generics so
// that casting is not required. All methods directly
// correspond to sync.Map methods. Do not copy a SyncMap
// after you start using it.
type SyncMap[K comparable, V any] struct {
	sync.Map
}

// Delete removes an item from the map
func (m *SyncMap[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

// Load looks up an item in the map. The bool is
// true if the value was found in the map.
func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	var zero V
	var v any
	v, ok = m.Map.Load(key)
	if ok {
		return v.(V), true
	}
	return zero, false
}

// LoadAndDelete looks up an item in the map, removes it from
// the map, and returns it. The boolean is true if the item was
// present in the map.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	var v any
	v, loaded = m.Map.LoadAndDelete(key)
	if loaded {
		return v.(V), true
	}
	var zero V
	return zero, false
}

// LoadOrStore looks up an item in the map. If present, it returns it and
// the returned boolean is true. If not present, the provided value is
// stored in the map and also returned.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	var v any
	v, loaded = m.Map.LoadOrStore(key, value)
	if loaded {
		return v.(V), true
	}
	return value, false
}

// Range iterates over the values in the map. Iteration stops if
// the provided function returns false. The provided value is not
// safe to use because it can be already overwritten at the time
// it is received. (Or at least that's my reading of the sync.Map
// documentation).
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(func(key any, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Store puts a value into the map
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.Map.Store(key, value)
}
