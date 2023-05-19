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

func TestTruncateLines(t *testing.T) {
	text := "First long line.\nSecond long line."
	exp := "First\nSecon"
	msg, err := TruncateLines(text)
	if msg != exp || err != nil {
		t.Fatalf(`TruncateLines(text) = %q, %v, want exp, error`, msg, err)
	}
}
