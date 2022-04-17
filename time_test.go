package atomix

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	a := NewTime(now)
	mustEqual(t, a.Load(), now)

	now2 := now.Add(time.Hour)
	a.Store(now2)
	mustEqual(t, a.Load(), now2)
}
