package gwrap

import (
	"reflect"
	"sync"
)

// SyncMap is a wrapper of sync.Map. It uses generics so
// that casting is not required. All methods directly
// correspond to sync.Map methods. Do not copy a SyncMap
// after you start using it.
//
// The underlying Map is accessible which creates a problem:
// if it's used to bypass SyncMap's methods and a value of a
// different type is stored in the map, that value will be
// ignored by SyncMap's Load, LoadAndDelete, LoadOrStore,
// and Range methods. There is one gotcha: for LoadOrStore unless both the
// wrong-typed value and the valid value are both comparable,
// there is no safe way to overwrite the invalid value so
// an unsafe operation is done instead: the invalid value is
// overwritten non-transactionally.
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
	if v, ok = m.Map.Load(key); ok {
		if typed, ok := v.(V); ok {
			return typed, true
		}
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
		if typed, ok := v.(V); ok {
			return typed, true
		}
	}
	var zero V
	return zero, false
}

// LoadOrStore looks up an item in the map. If present, it returns it and
// the returned boolean is true. If not present, the provided value is
// stored in the map and also returned.
//
// If the underlying map has a value of the wrong type stored in it (only
// possibly by bypassing SyncMap) and the wrong type value is not
// comparable then LoadOrStore will do a non-transactional overwrite of
// the existing value.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, ok := m.Map.LoadOrStore(key, value)
	if ok {
		for {
			if typed, ok := v.(V); ok {
				return typed, true
			}
			rv := reflect.ValueOf(v)
			if !rv.IsValid() || !rv.Comparable() {
				// We're now in a difficult spot. The value in the map
				// is not of the right type but CompareAndSwap is
				// forbidden because the old value is not comparable
				// so we don't have a good atomic way to
				// fix it. There is no good solution here.
				// We'll simply overwrite the value. That could
				// overwrite a valid value.
				m.Map.Store(key, value)
				break
			}
			//nolint:staticcheck // falsely reports Map could be omitted
			if m.Map.CompareAndSwap(key, v, value) {
				break
			}
			// We infer the wrongly-typed value changed. Re-load it.
			v, ok = m.Map.LoadOrStore(key, value)
			if !ok {
				break
			}
		}
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

// CompareMap is a wrapper of sync.Map. It uses generics so
// that casting is not required. All methods directly
// correspond to sync.Map methods. Do not copy a SyncMap
// after you start using it.
//
// CompareMap is only available when compiling with go 1.20 and above.
//
// CompareMap requires that values be comparable which is not true
// of SyncMap.
type CompareMap[K comparable, V comparable] struct {
	SyncMap[K, V]
}

// CompareAndDelete will delete a value from the map if the
// current value matches the provided value.
func (m *CompareMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.Map.CompareAndDelete(key, old)
}

// CompareAndSwap will replace a value in the map if the
// current value matches the provided "old" value.
func (m *CompareMap[K, V]) CompareAndSwap(key K, old V, new V) bool {
	return m.Map.CompareAndSwap(key, old, new)
}

// Swap will exchange the provided value with the value in the
// map.
//
// Swap is only available with go 1.20 and above
func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	var v any
	v, loaded = m.Map.Swap(key, value)
	if loaded {
		return v.(V), true
	}
	var zero V
	return zero, false
}
