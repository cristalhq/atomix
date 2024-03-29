package atomix

import (
	"strconv"
	"sync/atomic"
)

// Int32 is an atomic wrapper around an int32.
type Int32 struct {
	atomicType
	value int32
}

// NewInt32 creates an Int32.
func NewInt32(i int32) *Int32 {
	return &Int32{value: i}
}

func (i *Int32) String() string {
	return strconv.FormatInt(int64(i.Load()), 10)
}

// Load atomically the value.
func (i *Int32) Load() int32 {
	return atomic.LoadInt32(&i.value)
}

// Store atomically the given value.
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
// Deprecated: use Add(-n).
func (i *Int32) Sub(n int32) int32 {
	return atomic.AddInt32(&i.value, -n)
}

// Inc atomically and return the new value.
// Deprecated: use Add(1).
func (i *Int32) Inc() int32 {
	return i.Add(1)
}

// Dec atomically and return the new value.
// Deprecated: use Add(-1).
func (i *Int32) Dec() int32 {
	return i.Sub(1)
}

// CAS is an atomic Compare-And-Swap operation.
func (i *Int32) CAS(old, new int32) bool {
	return atomic.CompareAndSwapInt32(&i.value, old, new)
}

// SwapGreater value atomically, returns old and swap result.
func (i *Int32) SwapGreater(new int32) (old int32, swapped bool) {
	for {
		old := i.Load()
		if new <= old {
			return old, false
		}
		if i.CAS(old, new) {
			return old, true
		}
	}
}

// SwapLess value atomically, returns old and swap result.
func (i *Int32) SwapLess(new int32) (old int32, swapped bool) {
	for {
		old := i.Load()
		if new >= old {
			return old, false
		}
		if i.CAS(old, new) {
			return old, true
		}
	}
}
