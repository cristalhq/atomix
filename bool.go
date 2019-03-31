package atomix

import "sync/atomic"

// Bool is an atomic boolean.
type Bool struct {
	value uint32
}

// NewBool creates a Bool.
func NewBool(value bool) *Bool {
	return &Bool{
		value: b2i(value),
	}
}

// Load atomically loads the boolean.
func (b *Bool) Load() bool {
	return isTrue(atomic.LoadUint32(&b.value))
}

// Store atomically stores the passed value.
func (b *Bool) Store(new bool) {
	atomic.StoreUint32(&b.value, b2i(new))
}

// Swap sets the given value and returns the previous value.
func (b *Bool) Swap(new bool) bool {
	return isTrue(atomic.SwapUint32(&b.value, b2i(new)))
}

// Toggle atomically negates the boolean and returns the previous value.
func (b *Bool) Toggle() bool {
	return isTrue(atomic.AddUint32(&b.value, 1) - 1)
}

// CAS is an atomic compare-and-swap.
func (b *Bool) CAS(old, new bool) bool {
	return atomic.CompareAndSwapUint32(&b.value, b2i(old), b2i(new))
}

func isTrue(n uint32) bool {
	return (n & 1) == 1
}

func b2i(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}
