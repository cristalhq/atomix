package atomix

import (
	"sync/atomic"
	"time"
)

// Time is an atomic wrapper around an time.Time.
type Time struct {
	value atomic.Value
}

// NewTime creates a Time.
func NewTime(t time.Time) *Time {
	tm := &Time{}
	tm.Store(t)
	return tm
}

// Load atomically the value.
func (tm *Time) Load() time.Time {
	return tm.value.Load().(time.Time)
}

// Store atomically the given value.
func (tm *Time) Store(t time.Time) {
	tm.value.Store(tm)
}
