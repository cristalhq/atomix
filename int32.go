package atomix

import "sync/atomic"

// Int32 is an atomic wrapper around an int32.
type Int32 struct {
	noCopy
	value int32
}

// NewInt32 creates an Int32.
func NewInt32(i int32) *Int32 {
	return &Int32{
		value: i,
	}
}

// Load atomically the value.
func (i *Int32) Load() int32 {
	return atomic.LoadInt32(&i.value)
}

// Store atomically the passed value.
func (i *Int32) Store(n int32) {
	atomic.StoreInt32(&i.value, n)
}

// Swap atomically and return the old value.
func (i *Int32) Swap(n int32) int32 {
	return atomic.SwapInt32(&i.value, n)
}

// Add atomically and return the new value.
func (i *Int32) Add(n int32) int32 {
	return atomic.AddInt32(&i.value, n)
}

// Sub atomically and return the new value.
func (i *Int32) Sub(n int32) int32 {
	return atomic.AddInt32(&i.value, -n)
}

// Inc atomically and return the new value.
func (i *Int32) Inc() int32 {
	return i.Add(1)
}

// Dec atomically and return the new value.
func (i *Int32) Dec() int32 {
	return i.Sub(1)
}

// CAS is an atomic Compare-and-swap.
func (i *Int32) CAS(old, new int32) bool {
	return atomic.CompareAndSwapInt32(&i.value, old, new)
}
