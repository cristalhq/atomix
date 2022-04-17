package atomix

import "testing"

func TestFloat32(t *testing.T) {
	a := NewFloat32(10.5)

	mustEqual(t, a.String(), "10.5")
	mustEqual(t, a.Load(), float32(10.5))

	mustEqual(t, a.Add(0.3), float32(10.8))
	mustEqual(t, a.Sub(0.5), float32(10.3))

	mustEqual(t, a.CAS(10.3, 0.5), true)
	mustEqual(t, a.Load(), float32(0.5))
	mustEqual(t, a.CAS(0.0, 1.5), false)

	a.Store(42.0)
	mustEqual(t, a.Load(), float32(42.0))
}

func TestFloat32Compare(t *testing.T) {
	a := NewFloat32(42)

	old, ok := a.SwapGreater(80)
	mustEqual(t, old, float32(42.0))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), float32(80))

	old, ok = a.SwapGreater(40)
	mustEqual(t, old, float32(80.0))
	mustEqual(t, ok, false)

	old, ok = a.SwapLess(-80)
	mustEqual(t, old, float32(80))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), float32(-80))

	old, ok = a.SwapLess(-40)
	mustEqual(t, old, float32(-80.0))
	mustEqual(t, ok, false)
}
