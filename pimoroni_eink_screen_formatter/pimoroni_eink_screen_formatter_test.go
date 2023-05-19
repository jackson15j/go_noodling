package main

import (
	"testing"
)

func TestTruncate(t *testing.T) {
	line := "A really long line."
	exp := line[:5]
	msg, err := Truncate(line)
	if msg != exp || err != nil {
		t.Fatalf(`Truncate(line) = %q, %v, want exp, error`, msg, err)
	}
}
