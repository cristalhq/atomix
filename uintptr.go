package atomix

import "sync/atomic"

// Uintptr is an atomic uintptr.
type Uintptr struct {
	noCopy
	value uintptr
}

// NewUintptr creates an Uintptr.
func NewUintptr(value uintptr) *Uintptr {
	return &Uintptr{
		value: value,
	}
}

// Load atomically the value.
func (u *Uintptr) Load() uintptr {
	return atomic.LoadUintptr(&u.value)
}

// Store atomically the passed value.
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

// CAS is an atomic Compare-and-swap.
func (u *Uintptr) CAS(old, new uintptr) bool {
	return atomic.CompareAndSwapUintptr(&u.value, old, new)
}
