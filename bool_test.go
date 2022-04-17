package atomix

import (
	"testing"
)

func TestBool(t *testing.T) {
	a := NewBool(false)

	mustEqual(t, a.String(), "false")
	mustEqual(t, a.Toggle(), false)
	mustEqual(t, a.String(), "true")
	mustEqual(t, a.Load(), true)

	mustEqual(t, a.CAS(true, true), true)
	mustEqual(t, a.Load(), true)
	mustEqual(t, a.CAS(true, false), true)
	mustEqual(t, a.Load(), false)
	mustEqual(t, a.CAS(true, false), false)
	mustEqual(t, a.Load(), false)

	a.Store(false)
	mustEqual(t, a.Load(), false)

	mustEqual(t, a.Swap(false), false)
	mustEqual(t, a.Swap(true), false)
}
