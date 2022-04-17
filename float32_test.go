package atomix

import "testing"

func TestFloat32(t *testing.T) {
	a := NewFloat32(10.5)

	Equal(t, "10.5", a.String(), "Wrong String")
	Equal(t, float32(10.5), a.Load(), "Load wrong value")

	Equal(t, float32(10.8), a.Add(0.3), "Add wrong value")
	Equal(t, float32(10.3), a.Sub(0.5), "Sub wrong value")

	OK(t, a.CAS(10.3, 0.5), "CAS should swap")
	Equal(t, float32(0.5), a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(0.0, 1.5), "CAS shouldn't swap")

	a.Store(42.0)
	Equal(t, float32(42.0), a.Load(), "Store wrong value")
}

func TestFloat32Compare(t *testing.T) {
	a := NewFloat32(42)

	old, ok := a.SwapGreater(80)
	Equal(t, old, float32(42.0), "Store wrong value")
	Equal(t, true, ok, "Store wrong value")
	Equal(t, a.Load(), float32(80), "Store wrong value")

	old, ok = a.SwapGreater(40)
	Equal(t, old, float32(80.0), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")

	old, ok = a.SwapLess(-80)
	Equal(t, old, float32(80), "Store wrong value")
	Equal(t, ok, true, "Store wrong value")
	Equal(t, a.Load(), float32(-80), "Store wrong value")

	old, ok = a.SwapLess(-40)
	Equal(t, old, float32(-80.0), "Store wrong value")
	Equal(t, ok, false, "Store wrong value")
}
