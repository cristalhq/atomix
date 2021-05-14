package atomix

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	a := NewDuration(10 * time.Second)

	Equal(t, "10s", a.String(), "Wrong String")
	Equal(t, 10*time.Second, a.Load(), "Load wrong value")

	Equal(t, 15*time.Second, a.Add(5*time.Second), "Add wrong value")
	Equal(t, 12*time.Second, a.Sub(3*time.Second), "Sub wrong value")

	OK(t, a.CAS(12*time.Second, 0), "CAS should swap")
	Equal(t, 0*time.Second, a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(13*time.Second, 0), "CAS should not swap")

	Equal(t, 0*time.Second, a.Swap(1*time.Second), "Swap wrong value")
	Equal(t, 1*time.Second, a.Load(), "Swap wrong value")

	a.Store(15 * time.Second)
	Equal(t, 15*time.Second, a.Load(), "Store wrong value")
}
