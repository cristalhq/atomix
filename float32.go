package atomix

import (
	"math"
	"strconv"
	"sync/atomic"
)

// Float32 is an atomic wrapper around float32.
type Float32 struct {
	atomicType
	value uint32
}

// NewFloat32 creates a Float32.
func NewFloat32(f float32) *Float32 {
	return &Float32{value: math.Float32bits(f)}
}

func (f *Float32) String() string {
	return strconv.FormatFloat(float64(f.Load()), 'g', -1, 64)
}

// Load atomically the value.
func (f *Float32) Load() float32 {
	return math.Float32frombits(atomic.LoadUint32(&f.value))
}

// Store atomically the given value.
func (f *Float32) Store(s float32) {
	atomic.StoreUint32(&f.value, math.Float32bits(s))
}

// Add atomically and return the new value.
func (f *Float32) Add(s float32) float32 {
	for {
		old := f.Load()
		new := old + s
		if f.CAS(old, new) {
			return new
		}
	}
}

// Sub atomically and return the new value.
func (f *Float32) Sub(s float32) float32 {
	return f.Add(-s)
}

// CAS is an atomic Compare-And-Swap operation.
func (f *Float32) CAS(old, new float32) bool {
	return atomic.CompareAndSwapUint32(&f.value, math.Float32bits(old), math.Float32bits(new))
}

// SwapGreater value atomically, returns old and swap result.
func (f *Float32) SwapGreater(new float32) (old float32, swapped bool) {
	for {
		old := f.Load()
		if new <= old {
			return old, false
		}
		if f.CAS(old, new) {
			return old, true
		}
	}
}

// SwapLess value atomically, returns old and swap result.
func (f *Float32) SwapLess(new float32) (old float32, swapped bool) {
	for {
		old := f.Load()
		if new >= old {
			return old, false
		}
		if f.CAS(old, new) {
			return old, true
		}
	}
}
