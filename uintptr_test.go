package atomix

import (
	"testing"
)

func TestUintptr(t *testing.T) {
	a := NewUintptr(10)

	mustEqual(t, a.String(), "10")
	mustEqual(t, a.Load(), uintptr(10))

	mustEqual(t, a.Add(5), uintptr(15))
	mustEqual(t, a.Sub(3), uintptr(12))

	mustEqual(t, a.CAS(12, 0), true)
	mustEqual(t, a.Load(), uintptr(0))
	mustEqual(t, a.CAS(13, 0), false)

	mustEqual(t, a.Swap(1), uintptr(0))
	mustEqual(t, a.Load(), uintptr(1))

	a.Store(15)
	mustEqual(t, a.Load(), uintptr(15))
}
