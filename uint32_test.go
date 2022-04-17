package atomix

import (
	"testing"
)

func TestUint32(t *testing.T) {
	a := NewUint32(10)

	mustEqual(t, a.String(), "10")
	mustEqual(t, a.Load(), uint32(10))

	mustEqual(t, a.Add(5), uint32(15))
	mustEqual(t, a.Sub(3), uint32(12))

	mustEqual(t, a.Inc(), uint32(13))
	mustEqual(t, a.Dec(), uint32(12))

	mustEqual(t, a.CAS(12, 0), true)
	mustEqual(t, a.Load(), uint32(0))
	mustEqual(t, a.CAS(13, 0), false)

	mustEqual(t, a.Swap(1), uint32(0))
	mustEqual(t, a.Load(), uint32(1))

	a.Store(15)
	mustEqual(t, a.Load(), uint32(15))
}

func TestUint32Compare(t *testing.T) {
	a := NewUint32(42)

	old, ok := a.SwapGreater(80)
	mustEqual(t, old, uint32(42))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), uint32(80))

	old, ok = a.SwapGreater(40)
	mustEqual(t, old, uint32(80))
	mustEqual(t, ok, false)

	old, ok = a.SwapLess(8)
	mustEqual(t, old, uint32(80))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), uint32(8))

	old, ok = a.SwapLess(40)
	mustEqual(t, old, uint32(8))
	mustEqual(t, ok, false)
}
