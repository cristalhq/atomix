package atomix

import "sync/atomic"

// Value is a wrapper for atomically accessed consistently typed values.
type Value struct {
	value atomic.Value
}

// Load atomically the value.
func (v *Value) Load() interface{} {
	return v.value.Load()
}

// Store atomically the given value.
func (v *Value) Store(new interface{}) {
	v.value.Store(new)
}
