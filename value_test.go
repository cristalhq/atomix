package atomix

import "testing"

func TestValue(t *testing.T) {
	a := NewValue()
	Equal(t, nil, a.Load(), "Load wrong value")

	a = NewValue()
	a.Store("Hello")
	Equal(t, "Hello", a.Load(), "Load wrong value")

	a.Store("World")
	Equal(t, "World", a.Load(), "Store wrong value")
}
