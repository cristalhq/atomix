package atomix

import (
	"testing"
)

func TestInt32(t *testing.T) {
	a := NewInt32(10)

	mustEqual(t, a.String(), "10")
	mustEqual(t, a.Load(), int32(10))

	mustEqual(t, a.Add(5), int32(15))
	mustEqual(t, a.Sub(3), int32(12))

	mustEqual(t, a.Inc(), int32(13))
	mustEqual(t, a.Dec(), int32(12))

	mustEqual(t, a.CAS(12, 0), true)
	mustEqual(t, a.Load(), int32(0))
	mustEqual(t, a.CAS(13, 0), false)

	mustEqual(t, a.Swap(1), int32(0))
	mustEqual(t, a.Load(), int32(1))

	a.Store(15)
	mustEqual(t, a.Load(), int32(15))
}

func TestInt32Compare(t *testing.T) {
	a := NewInt32(42)

	old, ok := a.SwapGreater(80)
	mustEqual(t, old, int32(42))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), int32(80))

	old, ok = a.SwapGreater(40)
	mustEqual(t, old, int32(80))
	mustEqual(t, ok, false)

	old, ok = a.SwapLess(-80)
	mustEqual(t, old, int32(80))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), int32(-80))

	old, ok = a.SwapLess(-40)
	mustEqual(t, old, int32(-80))
	mustEqual(t, ok, false)
}
