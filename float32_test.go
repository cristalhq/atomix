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
