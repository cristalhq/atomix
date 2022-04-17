package atomix

import (
	"testing"
)

func TestUint32(t *testing.T) {
	a := NewUint32(10)

	Equal(t, "10", a.String(), "Wrong String")
	Equal(t, uint32(10), a.Load(), "Load wrong value")

	Equal(t, uint32(15), a.Add(5), "Add wrong value")
	Equal(t, uint32(12), a.Sub(3), "Sub wrong value")

	Equal(t, uint32(13), a.Inc(), "Inc wrong value")
	Equal(t, uint32(12), a.Dec(), "Dec wrong value")

	OK(t, a.CAS(12, 0), "CAS should swap")
	Equal(t, uint32(0), a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(13, 0), "CAS should not swap")

	Equal(t, uint32(0), a.Swap(1), "Swap wrong value")
	Equal(t, uint32(1), a.Load(), "Swap wrong value")

	a.Store(15)
	Equal(t, uint32(15), a.Load(), "Store wrong value")
}

func TestUint32Compare(t *testing.T) {
	a := NewUint32(42)

	old, ok := a.SwapGreater(80)
	Equal(t, old, uint32(42), "Store wrong value")
	Equal(t, true, ok, "Store wrong value")
	Equal(t, a.Load(), uint32(80), "Store wrong value")

	old, ok = a.SwapGreater(40)
	Equal(t, old, uint32(80), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")

	old, ok = a.SwapLess(8)
	Equal(t, old, uint32(80), "Store wrong value")
	Equal(t, ok, true, "Store wrong value")
	Equal(t, a.Load(), uint32(8), "Store wrong value")

	old, ok = a.SwapLess(40)
	Equal(t, old, uint32(8), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")
}
