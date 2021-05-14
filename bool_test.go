package atomix

import (
	"testing"
)

func TestBool(t *testing.T) {
	atom := NewBool(false)

	Equal(t, "false", atom.String(), "Expected 'false'")
	NotOK(t, atom.Toggle(), "Expected swap to return previous value.")
	Equal(t, "true", atom.String(), "Expected 'true'")
	OK(t, atom.Load(), "Unexpected state after swap.")

	OK(t, atom.CAS(true, true), "CAS should swap when old matches")
	OK(t, atom.Load(), "CAS should have no effect")
	OK(t, atom.CAS(true, false), "CAS should swap when old matches")
	NotOK(t, atom.Load(), "CAS should have modified the value")
	NotOK(t, atom.CAS(true, false), "CAS should fail on old mismatch")
	NotOK(t, atom.Load(), "CAS should not have modified the value")

	atom.Store(false)
	NotOK(t, atom.Load(), "Unexpected state after store.")

	prev := atom.Swap(false)
	NotOK(t, prev, "Expected Swap to return previous value.")

	prev = atom.Swap(true)
	NotOK(t, prev, "Expected Swap to return previous value.")
}
