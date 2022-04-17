package atomix

import (
	"testing"
)

func TestInt64(t *testing.T) {
	a := NewInt64(10)

	mustEqual(t, a.String(), "10")
	mustEqual(t, a.Load(), int64(10))

	mustEqual(t, a.Add(5), int64(15))
	mustEqual(t, a.Sub(3), int64(12))

	mustEqual(t, a.Inc(), int64(13))
	mustEqual(t, a.Dec(), int64(12))

	mustEqual(t, a.CAS(12, 0), true)
	mustEqual(t, a.Load(), int64(0))
	mustEqual(t, a.CAS(13, 0), false)

	mustEqual(t, a.Swap(1), int64(0))
	mustEqual(t, a.Load(), int64(1))

	a.Store(15)
	mustEqual(t, a.Load(), int64(15))
}

func TestInt64Compare(t *testing.T) {
	a := NewInt64(42)

	old, ok := a.SwapGreater(80)
	mustEqual(t, old, int64(42))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), int64(80))

	old, ok = a.SwapGreater(40)
	mustEqual(t, old, int64(80))
	mustEqual(t, ok, false)

	old, ok = a.SwapLess(-80)
	mustEqual(t, old, int64(80))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), int64(-80))

	old, ok = a.SwapLess(-40)
	mustEqual(t, old, int64(-80))
	mustEqual(t, ok, false)
}
