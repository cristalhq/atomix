package atomix

import (
	"strconv"
	"sync/atomic"
)

// Uint32 is an atomic wrapper around an uint32.
type Uint32 struct {
	atomicType
	value uint32
}

// NewUint32 creates an Uint32.
func NewUint32(i uint32) *Uint32 {
	return &Uint32{value: i}
}

func (u *Uint32) String() string {
	return strconv.FormatUint(uint64(u.Load()), 10)
}

// Load atomically the value.
func (u *Uint32) Load() uint32 {
	return atomic.LoadUint32(&u.value)
}

// Store atomically the given value.
func (u *Uint32) Store(n uint32) {
	atomic.StoreUint32(&u.value, n)
}

// Swap atomically and return the old value.
func (u *Uint32) Swap(n uint32) uint32 {
	return atomic.SwapUint32(&u.value, n)
}

// Add atomically and return the new value.
func (u *Uint32) Add(n uint32) uint32 {
	return atomic.AddUint32(&u.value, n)
}

// Sub atomically and return the new value.
// Deprecated: use Add(-n).
func (u *Uint32) Sub(n uint32) uint32 {
	return atomic.AddUint32(&u.value, ^(n - 1))
}

// Inc atomically and return the new value.
// Deprecated: use Add(1).
func (u *Uint32) Inc() uint32 {
	return u.Add(1)
}

// Dec atomically and return the new value.
// Deprecated: use Add(-1).
func (u *Uint32) Dec() uint32 {
	return u.Sub(1)
}

// CAS is an atomic Compare-And-Swap operation.
func (u *Uint32) CAS(old, new uint32) bool {
	return atomic.CompareAndSwapUint32(&u.value, old, new)
}

// SwapGreater value atomically, returns old and swap result.
func (u *Uint32) SwapGreater(new uint32) (old uint32, swapped bool) {
	for {
		old := u.Load()
		if new <= old {
			return old, false
		}
		if u.CAS(old, new) {
			return old, true
		}
	}
}

// SwapLess value atomically, returns old and swap result.
func (u *Uint32) SwapLess(new uint32) (old uint32, swapped bool) {
	for {
		old := u.Load()
		if new >= old {
			return old, false
		}
		if u.CAS(old, new) {
			return old, true
		}
	}
}
