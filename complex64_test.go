package atomix

import (
	"testing"
)

func TestComplex64(t *testing.T) {
	a := NewComplex64(complex(10.5, 10.5))
	var c complex64 = complex(10.5, 10.5)

	Equal(t, c, a.Load(), "Load wrong value")

	var c1 complex64 = complex(10.8, 10.8)
	var c2 complex64 = complex(10.3, 10.3)
	var c3 complex64 = complex(0.5, 0.5)
	var c4 complex64 = complex(0, 0)
	var c5 complex64 = complex(1.5, 1.5)

	Equal(t, c1, a.Add(complex(0.3, 0.3)), "Add wrong value")
	Equal(t, c2, a.Sub(complex(0.5, 0.5)), "Sub wrong value")

	OK(t, a.CAS(c2, complex(0.5, 0.5)), "CAS should swap")
	Equal(t, c3, a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(c4, c5), "CAS shouldn't swap")

	a.Store(complex(42.0, 47.0))
	var d complex64 = complex(42.0, 47.0)
	Equal(t, d, a.Load(), "Store wrong value")
}
