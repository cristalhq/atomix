package atomix

import (
	"strconv"
	"sync/atomic"
)

// Uint64 is an atomic wrapper around an uint64.
type Uint64 struct {
	atomicType
	value uint64
}

// NewUint64 creates an Uint64.
func NewUint64(i uint64) *Uint64 {
	return &Uint64{value: i}
}

func (i *Uint64) String() string {
	return strconv.FormatUint(i.Load(), 10)
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

// CAS is an atomic Compare-And-Swap operation.
func (i *Uint64) CAS(old, new uint64) bool {
	return atomic.CompareAndSwapUint64(&i.value, old, new)
}

// SwapGreater value atomically, returns old and swap result.
func (i *Uint64) SwapGreater(new uint64) (old uint64, swapped bool) {
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
func (i *Uint64) SwapLess(new uint64) (old uint64, swapped bool) {
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
