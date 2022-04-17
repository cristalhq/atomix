package atomix

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	a := NewDuration(10 * time.Second)

	mustEqual(t, a.String(), "10s")
	mustEqual(t, a.Load(), 10*time.Second)

	mustEqual(t, a.Add(5*time.Second), 15*time.Second)
	mustEqual(t, a.Sub(3*time.Second), 12*time.Second)

	mustEqual(t, a.CAS(12*time.Second, 0), true)
	mustEqual(t, a.Load(), 0*time.Second)
	mustEqual(t, a.CAS(13*time.Second, 0), false)

	mustEqual(t, a.Swap(1*time.Second), 0*time.Second)
	mustEqual(t, a.Load(), 1*time.Second)

	a.Store(15 * time.Second)
	mustEqual(t, a.Load(), 15*time.Second)
}

func TestDurationCompare(t *testing.T) {
	a := NewDuration(42)

	old, ok := a.SwapGreater(80)
	mustEqual(t, old, time.Duration(42.0))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), time.Duration(80))

	old, ok = a.SwapGreater(40)
	mustEqual(t, old, time.Duration(80.0))
	mustEqual(t, ok, false)

	old, ok = a.SwapLess(-80)
	mustEqual(t, old, time.Duration(80))
	mustEqual(t, ok, true)
	mustEqual(t, a.Load(), time.Duration(-80))

	old, ok = a.SwapLess(-40)
	mustEqual(t, old, time.Duration(-80.0))
	mustEqual(t, ok, false)
}
