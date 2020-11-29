package atomix

import (
	"math"
	"sync/atomic"
)

// Float32 is an atomic wrapper around float32.
type Float32 struct {
	atomicType
	v uint32
}

// NewFloat32 creates a float32.
func NewFloat32(f float32) *Float32 {
	return &Float32{v: math.Float32bits(f)}
}

// Load atomically the value.
func (f *Float32) Load() float32 {
	return math.Float32frombits(atomic.LoadUint32(&f.v))
}

// Store atomically the given value.
func (f *Float32) Store(s float32) {
	atomic.StoreUint32(&f.v, math.Float32bits(s))
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

// CAS is an atomic Compare-and-swap.
func (f *Float32) CAS(old, new float32) bool {
	return atomic.CompareAndSwapUint32(&f.v, math.Float32bits(old), math.Float32bits(new))
}
