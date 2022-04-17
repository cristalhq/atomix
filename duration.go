package atomix

import (
	"sync/atomic"
	"time"
)

// Duration is an atomic wrapper around an time.Duration.
type Duration struct {
	atomicType
	value int64
}

// NewDuration creates a Duration.
func NewDuration(d time.Duration) *Duration {
	return &Duration{value: int64(d)}
}

func (d *Duration) String() string {
	return d.Load().String()
}

// Load atomically the value.
func (d *Duration) Load() time.Duration {
	return time.Duration(atomic.LoadInt64(&d.value))
}

// Store atomically the given value.
func (d *Duration) Store(dur time.Duration) {
	atomic.StoreInt64(&d.value, int64(dur))
}

// Swap atomically and return the old value.
func (d *Duration) Swap(dur time.Duration) time.Duration {
	return time.Duration(atomic.SwapInt64(&d.value, int64(dur)))
}

// Add atomically and return the new value.
func (d *Duration) Add(dur time.Duration) time.Duration {
	return time.Duration(atomic.AddInt64(&d.value, int64(dur)))
}

// Sub atomically and return the new value.
func (d *Duration) Sub(dur time.Duration) time.Duration {
	return time.Duration(atomic.AddInt64(&d.value, -int64(dur)))
}

// CAS is an atomic Compare-And-Swap operation.
func (d *Duration) CAS(old, new time.Duration) bool {
	return atomic.CompareAndSwapInt64(&d.value, int64(old), int64(new))
}

// SwapGreater value atomically, returns old and swap result.
func (d *Duration) SwapGreater(new time.Duration) (old time.Duration, swapped bool) {
	for {
		old := d.Load()
		if new <= old {
			return old, false
		}
		if d.CAS(old, new) {
			return old, true
		}
	}
}

// SwapLess value atomically, returns old and swap result.
func (d *Duration) SwapLess(new time.Duration) (old time.Duration, swapped bool) {
	for {
		old := d.Load()
		if new >= old {
			return old, false
		}
		if d.CAS(old, new) {
			return old, true
		}
	}
}
