package atomix

import (
	"testing"
)

func TestUintptr(t *testing.T) {
	a := NewUintptr(10)

	Equal(t, "10", a.String(), "Wrong String")
	Equal(t, uintptr(10), a.Load(), "Load wrong value")

	Equal(t, uintptr(15), a.Add(5), "Add wrong value")
	Equal(t, uintptr(12), a.Sub(3), "Sub wrong value")

	OK(t, a.CAS(12, 0), "CAS should swap")
	Equal(t, uintptr(0), a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(13, 0), "CAS should not swap")

	Equal(t, uintptr(0), a.Swap(1), "Swap wrong value")
	Equal(t, uintptr(1), a.Load(), "Swap wrong value")

	a.Store(15)
	Equal(t, uintptr(15), a.Load(), "Store wrong value")
}
