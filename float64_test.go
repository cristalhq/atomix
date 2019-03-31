package atomix

import "testing"

func TestFloat64(t *testing.T) {
	a := NewFloat64(10.5)

	Equal(t, float64(10.5), a.Load(), "Load wrong value")

	Equal(t, float64(10.8), a.Add(0.3), "Add wrong value")
	Equal(t, float64(10.3), a.Sub(0.5), "Sub wrong value")

	OK(t, a.CAS(10.3, 0.5), "CAS should swap")
	Equal(t, float64(0.5), a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(0.0, 1.5), "CAS shouldn't swap")

	a.Store(42.0)
	Equal(t, float64(42.0), a.Load(), "Store wrong value")
}
