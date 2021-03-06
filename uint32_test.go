package atomix

import (
	"testing"
)

func TestUint32(t *testing.T) {
	a := NewUint32(10)

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
