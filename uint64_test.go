package atomix

import (
	"testing"
)

func TestUint64(t *testing.T) {
	a := NewUint64(10)

	mustEqual(t, a.String(), "10")
	mustEqual(t, a.Load(), uint64(10))

	mustEqual(t, a.Add(5), uint64(15))
	mustEqual(t, a.Sub(3), uint64(12))

	mustEqual(t, a.Inc(), uint64(13))
	mustEqual(t, a.Dec(), uint64(12))

	mustEqual(t, a.CAS(12, 0), true)
	mustEqual(t, a.Load(), uint64(0))
	mustEqual(t, a.CAS(13, 0), false)

	mustEqual(t, a.Swap(1), uint64(0))
	mustEqual(t, a.Load(), uint64(1))

	a.Store(15)
	mustEqual(t, a.Load(), uint64(15))
}

func TestUint64Compare(t *testing.T) {
	a := NewUint64(42)

	old, ok := a.SwapGreater(80)
	mustEqual(t, old, uint64(42))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), uint64(80))

	old, ok = a.SwapGreater(40)
	mustEqual(t, old, uint64(80))
	mustEqual(t, ok, false)

	old, ok = a.SwapLess(8)
	mustEqual(t, old, uint64(80))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), uint64(8))

	old, ok = a.SwapLess(40)
	mustEqual(t, old, uint64(8))
	mustEqual(t, ok, false)
}
