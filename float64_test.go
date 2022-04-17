package atomix

import "testing"

func TestFloat64(t *testing.T) {
	a := NewFloat64(10.5)

	mustEqual(t, a.String(), "10.5")
	mustEqual(t, a.Load(), float64(10.5))

	mustEqual(t, a.Add(0.3), float64(10.8))
	mustEqual(t, a.Sub(0.5), float64(10.3))

	mustEqual(t, a.CAS(10.3, 0.5), true)
	mustEqual(t, a.Load(), float64(0.5))
	mustEqual(t, a.CAS(0.0, 1.5), false)

	a.Store(42.0)
	mustEqual(t, a.Load(), float64(42.0))
}

func TestFloat64Compare(t *testing.T) {
	a := NewFloat64(42)

	old, ok := a.SwapGreater(80)
	mustEqual(t, old, float64(42.0))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), float64(80))

	old, ok = a.SwapGreater(40)
	mustEqual(t, old, float64(80.0))
	mustEqual(t, ok, false)

	old, ok = a.SwapLess(-80)
	mustEqual(t, old, float64(80))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), float64(-80))

	old, ok = a.SwapLess(-40)
	mustEqual(t, old, float64(-80.0))
	mustEqual(t, ok, false)
}
