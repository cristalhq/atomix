package atomix

import "testing"

func TestUint(t *testing.T) {
	a := NewUint(10)

	mustEqual(t, a.String(), "10")
	mustEqual(t, a.Load(), uint(10))

	mustEqual(t, a.Add(5), uint(15))
	mustEqual(t, a.Sub(3), uint(12))

	mustEqual(t, a.Inc(), uint(13))
	mustEqual(t, a.Dec(), uint(12))

	mustEqual(t, a.CAS(12, 0), true)
	mustEqual(t, a.Load(), uint(0))
	mustEqual(t, a.CAS(13, 0), false)

	mustEqual(t, a.Swap(1), uint(0))
	mustEqual(t, a.Load(), uint(1))

	a.Store(15)
	mustEqual(t, a.Load(), uint(15))
}

func TestUintCompare(t *testing.T) {
	a := NewUint(42)

	old, ok := a.SwapGreater(80)
	mustEqual(t, old, uint(42))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), uint(80))

	old, ok = a.SwapGreater(40)
	mustEqual(t, old, uint(80))
	mustEqual(t, ok, false)

	old, ok = a.SwapLess(8)
	mustEqual(t, old, uint(80))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), uint(8))

	old, ok = a.SwapLess(40)
	mustEqual(t, old, uint(8))
	mustEqual(t, ok, false)
}
