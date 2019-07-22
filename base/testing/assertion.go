package btest

import (
	"testing"

	"github.com/d4l3k/messagediff"
	"github.com/google/go-cmp/cmp"
)

// Check a condition to be true, with an option to print an error message with
// the template passed in.
func Expect(t *testing.T, condition bool, messageTemplate string, args ...interface{}) {
	if !condition {
		t.Helper()
		t.Errorf(messageTemplate, args...)
	}
}

// Check the two values passed in are equal.
func ExpectEq(t *testing.T, got, expected interface{}) {
	if !cmp.Equal(expected, got) {
		t.Helper()
		t.Errorf("Got: %T[%+v]\nExpected: %T[%+v]", got, got, expected, expected)
	}
}

// Check the two structs passed in are equal.
func ExpectEqStruct(t *testing.T, got, expected interface{}) {
	if diff, equal := messagediff.PrettyDiff(expected, got); !equal {
		t.Helper()
		t.Errorf("Got: %T[%+v]\nDiff: [%+v]", got, got, diff)
	}
}

// Assert a condition to be true, with an option to print an error message with
// the template passed in.
func Assert(t *testing.T, condition bool, messageTemplate string, args ...interface{}) {
	if !condition {
		t.Helper()
		t.Fatalf(messageTemplate, args...)
	}
}

// Assert an error is nil.
func AssertOk(t *testing.T, err error) {
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}
