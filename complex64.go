package atomix

import (
	"math"
	"sync/atomic"
)

// Complex64 is an atomic wrapper around float32.
type Complex64 struct {
	ri uint64
}

// NewComplex64 creates a float32.
func NewComplex64(c complex64) *Complex64 {
	return &Complex64{ri: complex64ToUint64(c)}
}

// Load atomically the value.
func (c *Complex64) Load() complex64 {
	return uint64ToComplex64(atomic.LoadUint64(&c.ri))
}

// Store atomically the given value.
func (c *Complex64) Store(s complex64) {
	atomic.StoreUint64(&c.ri, complex64ToUint64(s))
}

// Add atomically and return the new value.
func (c *Complex64) Add(s complex64) complex64 {
	for {
		oc := c.Load()
		nc := oc + s
		if c.CAS(oc, nc) {
			return nc
		}
	}
}

// Sub atomically and return the new value.
func (c *Complex64) Sub(s complex64) complex64 {
	return c.Add(-s)
}

// CAS is an atomic Compare-and-swap.
func (c *Complex64) CAS(oc, nc complex64) bool {
	return atomic.CompareAndSwapUint64(&c.ri, complex64ToUint64(oc), complex64ToUint64(nc))
}

func uint64ToComplex64(u uint64) complex64 {
	return complex(math.Float32frombits(uint32(u>>32)), math.Float32frombits(uint32((u<<32)>>32)))
}

func complex64ToUint64(c complex64) uint64 {
	return uint64(math.Float32bits(real(c)))<<32 | uint64(math.Float32bits(imag(c)))
}
