package atomix

import (
	"sync/atomic"
)

// Error is an atomic wrapper around error.
type Error struct {
	atomicType
	value atomic.Value
}

var _ error = &Error{}

// NewError creates an Error.
// Cannot store nil after first non-nil store.
func NewError(err error) *Error {
	e := &Error{}
	e.Store(err)
	return e
}

func (e *Error) String() string {
	return e.Error()
}

func (e *Error) Error() string {
	if err := e.Load(); err != nil {
		return err.Error()
	}
	return "<nil>"
}

// HasError reports whether non-nil error is stored.
func (e *Error) HasError() bool {
	return e.Load() != nil
}

// Load atomically the value.
func (e *Error) Load() error {
	v := e.value.Load()
	if v == nil {
		return nil
	}
	return v.(error)
}

// Store atomically the given value.
// Doesn't store nil error.
func (e *Error) Store(err error) {
	if err != nil {
		e.value.Store(err)
	}
}
