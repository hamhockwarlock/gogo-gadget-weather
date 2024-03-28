package test

import "testing"

func Fatal(t *testing.T, want, got interface{}) {
	t.Helper()
	t.Fatalf(`want: %v, got: %v`, want, got)
}
