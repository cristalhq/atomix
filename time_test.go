package atomix

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	a := NewTime(now)
	Equal(t, now, a.Load(), "Load wrong value")

	now2 := now.Add(time.Hour)
	a.Store(now2)
	Equal(t, now2, a.Load(), "Store wrong value")
}
