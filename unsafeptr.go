// +build !purego !appengine !js

package atomix

import (
	"sync/atomic"
	"unsafe"
)

// UnsafePointer is an atomic unsafe.Pointer.
type UnsafePointer struct {
	value unsafe.Pointer
}

// NewUnsafePointer creates a UnsafePointer.
func NewUnsafePointer(value unsafe.Pointer) *UnsafePointer {
	return &UnsafePointer{value: value}
}

// Load atomically loads the boolean.
func (p *UnsafePointer) Load() unsafe.Pointer {
	return atomic.LoadPointer(&p.value)
}

// Store atomically stores the given value.
func (p *UnsafePointer) Store(new unsafe.Pointer) {
	atomic.StorePointer(&p.value, new)
}

// Swap sets the given value and returns the previous value.
func (p *UnsafePointer) Swap(new unsafe.Pointer) unsafe.Pointer {
	return atomic.SwapPointer(&p.value, new)
}

// CAS is an atomic compare-and-swap.
func (p *UnsafePointer) CAS(old, new unsafe.Pointer) bool {
	return atomic.CompareAndSwapPointer(&p.value, old, new)
}
