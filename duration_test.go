package atomix

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	a := NewDuration(10 * time.Second)

	Equal(t, "10s", a.String(), "Wrong String")
	Equal(t, 10*time.Second, a.Load(), "Load wrong value")

	Equal(t, 15*time.Second, a.Add(5*time.Second), "Add wrong value")
	Equal(t, 12*time.Second, a.Sub(3*time.Second), "Sub wrong value")

	OK(t, a.CAS(12*time.Second, 0), "CAS should swap")
	Equal(t, 0*time.Second, a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(13*time.Second, 0), "CAS should not swap")

	Equal(t, 0*time.Second, a.Swap(1*time.Second), "Swap wrong value")
	Equal(t, 1*time.Second, a.Load(), "Swap wrong value")

	a.Store(15 * time.Second)
	Equal(t, 15*time.Second, a.Load(), "Store wrong value")
}

func TestDurationCompare(t *testing.T) {
	a := NewDuration(42)

	old, ok := a.SwapGreater(80)
	Equal(t, old, time.Duration(42.0), "Store wrong value")
	Equal(t, true, ok, "Store wrong value")
	Equal(t, a.Load(), time.Duration(80), "Store wrong value")

	old, ok = a.SwapGreater(40)
	Equal(t, old, time.Duration(80.0), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")

	old, ok = a.SwapLess(-80)
	Equal(t, old, time.Duration(80), "Store wrong value")
	Equal(t, ok, true, "Store wrong value")
	Equal(t, a.Load(), time.Duration(-80), "Store wrong value")

	old, ok = a.SwapLess(-40)
	Equal(t, old, time.Duration(-80.0), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")
}
