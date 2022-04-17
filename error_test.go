package atomix

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	a := NewError(nil)

	mustEqual(t, a.String(), "<nil>")
	mustEqual(t, a.HasError(), false)
	mustEqual(t, a.Load(), nil)

	err := errors.New("ouch")
	a = NewError(err)
	mustEqual(t, a.HasError(), true)
	mustEqual(t, a.Load(), err)
	mustEqual(t, a.String(), "ouch")

	err2 := errors.New("very-ouch")
	a.Store(err2)
	mustEqual(t, a.Load(), err2)

	a.Store(nil)
	mustEqual(t, a.Load(), err2)
}
