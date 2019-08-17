package atomix

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}, msg string) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(`%s:
expected: %v
actual  : %v`, msg, expected, actual)
	}
}

func OK(t *testing.T, cond bool, msg string) {
	t.Helper()

	if !cond {
		t.Errorf("%s: should be true", msg)
	}
}

func NotOK(t *testing.T, cond bool, msg string) {
	t.Helper()

	if cond {
		t.Errorf("%s: should be false", msg)
	}
}
