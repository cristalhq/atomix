package atomix

import (
	"reflect"
	"testing"
)

func mustEqual(tb testing.TB, got, want interface{}) {
	tb.Helper()

	if !reflect.DeepEqual(got, want) {
		tb.Fatalf("got: %v, want: %v", got, want)
	}
}
