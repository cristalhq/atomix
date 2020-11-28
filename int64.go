package atomix

import "sync/atomic"

// Int64 is an atomic wrapper around an int64.
type Int64 struct {
	noCopy
	value int64
}

// NewInt64 creates an Int64.
func NewInt64(i int64) *Int64 {
	return &Int64{
		value: i,
	}
}

// Load atomically the value.
func (i *Int64) Load() int64 {
	return atomic.LoadInt64(&i.value)
}

// Store atomically the passed value.
func (i *Int64) Store(n int64) {
	atomic.StoreInt64(&i.value, n)
}

// Swap atomically and return the old value.
func (i *Int64) Swap(n int64) int64 {
	return atomic.SwapInt64(&i.value, n)
}

// Add atomically and return the new value.
func (i *Int64) Add(n int64) int64 {
	return atomic.AddInt64(&i.value, n)
}

// Sub atomically and return the new value.
func (i *Int64) Sub(n int64) int64 {
	return atomic.AddInt64(&i.value, -n)
}

// Inc atomically and return the new value.
func (i *Int64) Inc() int64 {
	return i.Add(1)
}

// Dec atomically and return the new value.
func (i *Int64) Dec() int64 {
	return i.Sub(1)
}

// CAS is an atomic Compare-and-swap.
func (i *Int64) CAS(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&i.value, old, new)
}
