package atomix

import "testing"

func TestComplex64(t *testing.T) {
	a := NewComplex64(complex(float32(10.5), float32(12.3)))

	Equal(t, "(10.5+12.3i)", a.String(), "Wrong String")

	Equal(t, complex(float32(10.5), 12.3), a.Load(), "Load wrong value")

	Equal(t, complex(float32(10.8), float32(12.6)), a.Add(complex(float32(0.3), float32(0.3))), "Add wrong value")
	Equal(t, complex(float32(10.3), float32(12.1)), a.Sub(complex(float32(0.5), float32(0.5))), "Sub wrong value")

	OK(t, a.CAS(complex(float32(10.3), float32(12.1)), complex(float32(0.5), float32(0.5))), "CAS should swap")
	Equal(t, complex(float32(0.5), float32(0.5)), a.Load(), "CAS wrong value")
	NotOK(t, a.CAS(complex(float32(0.0), float32(0.0)), complex(float32(1.5), float32(1.5))), "CAS shouldn't swap")

	a.Store(complex(float32(42.0), float32(42.0)))
	Equal(t, complex(float32(42.0), float32(42.0)), a.Load(), "Store wrong value")
}
