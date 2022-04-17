package atomix

import "testing"

func TestString(t *testing.T) {
	a := NewString("")
	mustEqual(t, a.Load(), "")

	a = NewString("Hello")
	mustEqual(t, a.Load(), "Hello")

	mustEqual(t, a.Load(), a.String())

	a.Store("World")
	mustEqual(t, a.Load(), "World")

	s := String{}
	mustEqual(t, s.Load(), "")
}
