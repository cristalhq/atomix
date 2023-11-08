package atomix

import (
	"strconv"
	"sync/atomic"
)

// Int64 is an atomic wrapper around an int64.
type Int64 struct {
	atomicType
	value int64
}

// NewInt64 creates an Int64.
func NewInt64(i int64) *Int64 {
	return &Int64{value: i}
}

func (i *Int64) String() string {
	return strconv.FormatInt(i.Load(), 10)
}

// Load atomically the value.
func (i *Int64) Load() int64 {
	return atomic.LoadInt64(&i.value)
}

// Store atomically the given value.
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
// Deprecated: use Add(-n).
func (i *Int64) Sub(n int64) int64 {
	return atomic.AddInt64(&i.value, -n)
}

// Inc atomically and return the new value.
// Deprecated: use Add(1).
func (i *Int64) Inc() int64 {
	return i.Add(1)
}

// Dec atomically and return the new value.
// Deprecated: use Add(-1).
func (i *Int64) Dec() int64 {
	return i.Sub(1)
}

// CAS is an atomic Compare-And-Swap operation.
func (i *Int64) CAS(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&i.value, old, new)
}

// SwapGreater value atomically, returns old and swap result.
func (i *Int64) SwapGreater(new int64) (old int64, swapped bool) {
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
func (i *Int64) SwapLess(new int64) (old int64, swapped bool) {
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
