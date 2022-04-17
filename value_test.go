package atomix

import "testing"

func TestValue(t *testing.T) {
	a := NewValue()
	mustEqual(t, a.Load(), nil)

	a = NewValue()
	a.Store("Hello")
	mustEqual(t, a.Load(), "Hello")

	a.Store("World")
	mustEqual(t, a.Load(), "World")
}
