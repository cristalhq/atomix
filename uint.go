package atomix

import (
	"strconv"
)

// Uint is an atomic wrapper around an uint.
type Uint struct {
	atomicType
	value uint
}

// NewUint creates an Uint.
func NewUint(i uint) *Uint {
	return &Uint{value: i}
}

func (i *Uint) String() string {
	return strconv.FormatUint(uint64(i.Load()), 10)
}

// Load atomically the value.
func (i *Uint) Load() uint {
	return LoadUint(&i.value)
}

// Store atomically the given value.
func (i *Uint) Store(n uint) {
	StoreUint(&i.value, n)
}

// Swap atomically and return the old value.
func (i *Uint) Swap(n uint) uint {
	return SwapUint(&i.value, n)
}

// Add atomically and return the new value.
func (i *Uint) Add(n uint) uint {
	return AddUint(&i.value, n)
}

// Sub atomically and return the new value.
func (i *Uint) Sub(n uint) uint {
	return AddUint(&i.value, -n)
}

// Inc atomically and return the new value.
func (i *Uint) Inc() uint {
	return i.Add(1)
}

// Dec atomically and return the new value.
func (i *Uint) Dec() uint {
	return i.Sub(1)
}

// CAS is an atomic Compare-And-Swap operation.
func (i *Uint) CAS(old, new uint) bool {
	return CompareAndSwapUint(&i.value, old, new)
}

// SwapGreater value atomically, returns old and swap result.
func (i *Uint) SwapGreater(new uint) (old uint, swapped bool) {
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
func (i *Uint) SwapLess(new uint) (old uint, swapped bool) {
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
