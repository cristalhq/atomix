package atomix

import (
	"math"
	"sync/atomic"
)

// Complex64 is an atomic wrapper around float32.
type Complex64 struct {
	r uint32
	i uint32
}

// NewComplex65 creates a float32.
func NewComplex64(f complex64) *Complex64 {
	return &Complex64{
		math.Float32bits(real(f)),
		math.Float32bits(imag(f)),
	}
}

// Load atomically the value.
func (f *Complex64) Load() complex64 {
	return complex(math.Float32frombits(atomic.LoadUint32(&f.r)), math.Float32frombits(atomic.LoadUint32(&f.i)))
}

// Store atomically the passed value.
func (f *Complex64) Store(s complex64) {
	atomic.StoreUint32(&f.r, math.Float32bits(real(s)))
	atomic.StoreUint32(&f.i, math.Float32bits(imag(s)))
}

// Add atomically and return the new value.
func (f *Complex64) Add(s complex64) complex64 {
	for {
		old := f.Load()
		new := old + s
		if f.CAS(old, new) {
			return new
		}
	}
}

// Sub atomically and return the new value.
func (f *Complex64) Sub(s complex64) complex64 {
	return f.Add(-s)
}

// CAS is an atomic Compare-and-swap.
func (f *Complex64) CAS(old, new complex64) bool {
	r := atomic.CompareAndSwapUint32(&f.r, math.Float32bits(real(old)), math.Float32bits(real(new)))
	i := atomic.CompareAndSwapUint32(&f.i, math.Float32bits(imag(old)), math.Float32bits(imag(new)))
	return r && i
}
