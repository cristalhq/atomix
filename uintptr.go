package atomix

import (
	"strconv"
	"sync/atomic"
)

// Uintptr is an atomic uintptr.
type Uintptr struct {
	atomicType
	value uintptr
}

// NewUintptr creates an Uintptr.
func NewUintptr(ptr uintptr) *Uintptr {
	return &Uintptr{value: ptr}
}

func (u *Uintptr) String() string {
	return strconv.FormatUint(uint64(u.Load()), 10)
}

// Load atomically the value.
func (u *Uintptr) Load() uintptr {
	return atomic.LoadUintptr(&u.value)
}

// Store atomically the given value.
func (u *Uintptr) Store(n uintptr) {
	atomic.StoreUintptr(&u.value, n)
}

// Swap atomically and return the old value.
func (u *Uintptr) Swap(n uintptr) uintptr {
	return atomic.SwapUintptr(&u.value, n)
}

// Add atomically and return the new value.
func (u *Uintptr) Add(n uintptr) uintptr {
	return atomic.AddUintptr(&u.value, n)
}

// Sub atomically and return the new value.
func (u *Uintptr) Sub(n uintptr) uintptr {
	return u.Add(^(n - 1))
}

// CAS is an atomic Compare-And-Swap operation.
func (u *Uintptr) CAS(old, new uintptr) bool {
	return atomic.CompareAndSwapUintptr(&u.value, old, new)
}
