package gwrap

import "sync/atomic"

// Atomic value is a wrapper around [atomic.Value].
// AtomicValues should not be copied after first use
type AtomicValue[T comparable] struct {
	atomic.Value
}

// CompareAndSwap exchanges values if the current value matches the passed
// in old value.
func (av *AtomicValue[T]) CompareAndSwap(old T, new T) (swapped bool) {
	return av.Value.CompareAndSwap(old, new)
}

// Load fetches a value
func (av *AtomicValue[T]) Load() T {
	v := av.Value.Load()
	if v == nil {
		var zero T
		return zero
	}
	return v.(T)
}

// Store stores a value
func (av *AtomicValue[T]) Store(v T) {
	av.Value.Store(v)
}

// Swap exchanges a value for the current value. If there is
// no current value, the zero value of T is returned. Unlike
// [atomic.Value_Swap], the nil value is allowed if T is a
// pointer type.
func (av *AtomicValue[T]) Swap(new T) T {
	v := av.Value.Swap(new)
	if v == nil {
		var zero T
		return zero
	}
	return v.(T)
}
