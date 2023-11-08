package atomix

import (
	"strconv"
)

// Int is an atomic wrapper around an int.
type Int struct {
	atomicType
	value int
}

// NewInt creates an Int.
func NewInt(i int) *Int {
	return &Int{value: i}
}

func (i *Int) String() string {
	return strconv.FormatInt(int64(i.Load()), 10)
}

// Load atomically the value.
func (i *Int) Load() int {
	return LoadInt(&i.value)
}

// Store atomically the given value.
func (i *Int) Store(n int) {
	StoreInt(&i.value, n)
}

// Swap atomically and return the old value.
func (i *Int) Swap(n int) int {
	return SwapInt(&i.value, n)
}

// Add atomically and return the new value.
func (i *Int) Add(n int) int {
	return AddInt(&i.value, n)
}

// Sub atomically and return the new value.
// Deprecated: use Add(-n).
func (i *Int) Sub(n int) int {
	return AddInt(&i.value, -n)
}

// Inc atomically and return the new value.
// Deprecated: use Add(1).
func (i *Int) Inc() int {
	return i.Add(1)
}

// Dec atomically and return the new value.
// Deprecated: use Add(-1).
func (i *Int) Dec() int {
	return i.Sub(1)
}

// CAS is an atomic Compare-And-Swap operation.
func (i *Int) CAS(old, new int) bool {
	return CompareAndSwapInt(&i.value, old, new)
}

// SwapGreater value atomically, returns old and swap result.
func (i *Int) SwapGreater(new int) (old int, swapped bool) {
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
func (i *Int) SwapLess(new int) (old int, swapped bool) {
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
