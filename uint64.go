package atomix

import "sync/atomic"

// Uint64 is an atomic wrapper around an uint64.
type Uint64 struct {
	value uint64
}

// NewUint64 creates a Uint64.
func NewUint64(i uint64) *Uint64 {
	return &Uint64{value: i}
}

// Load atomically the value.
func (i *Uint64) Load() uint64 {
	return atomic.LoadUint64(&i.value)
}

// Store atomically the given value.
func (i *Uint64) Store(n uint64) {
	atomic.StoreUint64(&i.value, n)
}

// Swap atomically and return the old value.
func (i *Uint64) Swap(n uint64) uint64 {
	return atomic.SwapUint64(&i.value, n)
}

// Add atomically and return the new value.
func (i *Uint64) Add(n uint64) uint64 {
	return atomic.AddUint64(&i.value, n)
}

// Sub atomically and return the new value.
func (i *Uint64) Sub(n uint64) uint64 {
	return atomic.AddUint64(&i.value, ^(n - 1))
}

// Inc atomically and return the new value.
func (i *Uint64) Inc() uint64 {
	return i.Add(1)
}

// Dec atomically and return the new value.
func (i *Uint64) Dec() uint64 {
	return i.Sub(1)
}

// CAS is an atomic Compare-and-swap.
func (i *Uint64) CAS(old, new uint64) bool {
	return atomic.CompareAndSwapUint64(&i.value, old, new)
}
