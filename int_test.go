package atomix

import "testing"

func TestInt(t *testing.T) {
	a := NewInt(10)

	mustEqual(t, a.String(), "10")
	mustEqual(t, a.Load(), int(10))

	mustEqual(t, a.Add(5), int(15))
	mustEqual(t, a.Sub(3), int(12))

	mustEqual(t, a.Inc(), int(13))
	mustEqual(t, a.Dec(), int(12))

	mustEqual(t, a.CAS(12, 0), true)
	mustEqual(t, a.Load(), int(0))
	mustEqual(t, a.CAS(13, 0), false)

	mustEqual(t, a.Swap(1), int(0))
	mustEqual(t, a.Load(), int(1))

	a.Store(15)
	mustEqual(t, a.Load(), int(15))
}

func TestIntCompare(t *testing.T) {
	a := NewInt(42)

	old, ok := a.SwapGreater(80)
	mustEqual(t, old, int(42))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), int(80))

	old, ok = a.SwapGreater(40)
	mustEqual(t, old, int(80))
	mustEqual(t, ok, false)

	old, ok = a.SwapLess(8)
	mustEqual(t, old, int(80))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), int(8))

	old, ok = a.SwapLess(40)
	mustEqual(t, old, int(8))
	mustEqual(t, ok, false)
}
