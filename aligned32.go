package atomix

import (
	"strconv"
	"sync/atomic"
)

// AlignedInt32 is an atomic wrapper around an int32 aligned to a cache line.
type AlignedInt32 struct {
	value int32
	_     [CacheLine - 4]byte // unsafe.Sizeof(int32) == 4
}

// NewAlignedInt32 creates an AlignedInt32.
func NewAlignedInt32(i int32) *AlignedInt32 {
	return &AlignedInt32{value: i}
}

func (a *AlignedInt32) String() string {
	return strconv.FormatInt(int64(a.Load()), 10)
}

// Load atomically the value.
func (a *AlignedInt32) Load() int32 {
	return atomic.LoadInt32(&a.value)
}

// Store atomically the given value.
func (a *AlignedInt32) Store(n int32) {
	atomic.StoreInt32(&a.value, n)
}

// Swap atomically and return the old value.
func (a *AlignedInt32) Swap(n int32) int32 {
	return atomic.SwapInt32(&a.value, n)
}

// Add atomically and return the new value.
func (a *AlignedInt32) Add(n int32) int32 {
	return atomic.AddInt32(&a.value, n)
}

// Sub atomically and return the new value.
func (a *AlignedInt32) Sub(n int32) int32 {
	return atomic.AddInt32(&a.value, -n)
}

// Inc atomically and return the new value.
func (a *AlignedInt32) Inc() int32 {
	return a.Add(1)
}

// Dec atomically and return the new value.
func (a *AlignedInt32) Dec() int32 {
	return a.Sub(1)
}

// CAS is an atomic Compare-And-Swap operation.
func (a *AlignedInt32) CAS(old, new int32) bool {
	return atomic.CompareAndSwapInt32(&a.value, old, new)
}
