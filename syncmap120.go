//go:build go1.20

package gwrap

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
