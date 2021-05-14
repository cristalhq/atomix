package atomix

import "testing"

func TestString(t *testing.T) {
	a := NewString("")
	Equal(t, "", a.Load(), "Load wrong value")

	a = NewString("Hello")
	Equal(t, "Hello", a.Load(), "Load wrong value")

	Equal(t, a.String(), a.Load(), "Load != String")

	a.Store("World")
	Equal(t, "World", a.Load(), "Store wrong value")

	s := String{}
	Equal(t, "", s.Load(), "nil String should be empty")
}
