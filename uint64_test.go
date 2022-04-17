package atomix

import (
	"testing"
)

func TestUint64(t *testing.T) {
	a := NewUint64(10)

	Equal(t, "10", a.String(), "Wrong String")
	Equal(t, uint64(10), a.Load(), "Load wrong value")

	Equal(t, uint64(15), a.Add(5), "Add wrong value")
	Equal(t, uint64(12), a.Sub(3), "Sub wrong value")

	Equal(t, uint64(13), a.Inc(), "Inc wrong value")
	Equal(t, uint64(12), a.Dec(), "Dec wrong value")

	OK(t, a.CAS(12, 0), "CAS should swap")
	Equal(t, uint64(0), a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(13, 0), "CAS should not swap")

	Equal(t, uint64(0), a.Swap(1), "Swap wrong value")
	Equal(t, uint64(1), a.Load(), "Swap wrong value")

	a.Store(15)
	Equal(t, uint64(15), a.Load(), "Store wrong value")
}

func TestUint64Compare(t *testing.T) {
	a := NewUint64(42)

	old, ok := a.SwapGreater(80)
	Equal(t, old, uint64(42), "Store wrong value")
	Equal(t, true, ok, "Store wrong value")
	Equal(t, a.Load(), uint64(80), "Store wrong value")

	old, ok = a.SwapGreater(40)
	Equal(t, old, uint64(80), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")

	old, ok = a.SwapLess(8)
	Equal(t, old, uint64(80), "Store wrong value")
	Equal(t, ok, true, "Store wrong value")
	Equal(t, a.Load(), uint64(8), "Store wrong value")

	old, ok = a.SwapLess(40)
	Equal(t, old, uint64(8), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")
}
