package atomix

import "testing"

func TestComplex64(t *testing.T) {
	a := NewComplex64(complex(float32(10.5), float32(12.3)))

	mustEqual(t, a.String(), "(10.5+12.3i)")

	mustEqual(t, a.Load(), complex(float32(10.5), 12.3))

	mustEqual(t, a.Add(complex(float32(0.3), float32(0.3))), complex(float32(10.8), float32(12.6)))
	mustEqual(t, a.Sub(complex(float32(0.5), float32(0.5))), complex(float32(10.3), float32(12.1)))

	mustEqual(t, a.CAS(complex(float32(10.3), float32(12.1)), complex(float32(0.5), float32(0.5))), true)
	mustEqual(t, a.Load(), complex(float32(0.5), float32(0.5)))
	mustEqual(t, a.CAS(complex(float32(0.0), float32(0.0)), complex(float32(1.5), float32(1.5))), false)

	a.Store(complex(float32(42.0), float32(42.0)))
	mustEqual(t, a.Load(), complex(float32(42.0), float32(42.0)))
}
