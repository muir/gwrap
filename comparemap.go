//go:build go1.20

package gwrap

// CompareMap is only available when compiling with go 1.20 and above
type CompareMap[K comparable, V comparable] struct {
	SyncMap[K, V]
}

func (m *CompareMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.Map.CompareAndDelete(key, old)
}

func (m *CompareMap[K, V]) CompareAndSwap(key K, old V, new V) bool {
	return m.Map.CompareAndSwap(key, old, new)
}

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
