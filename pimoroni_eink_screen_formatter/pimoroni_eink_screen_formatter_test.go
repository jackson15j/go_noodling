package main

import (
	"os"
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

func TestTruncateLinesHandleEmptyLines(t *testing.T) {
	text := "First long line.\n\n\n"
	exp := "First\n\n\n"
	msg, err := TruncateLines(text)
	if msg != exp || err != nil {
		t.Fatalf(`TruncateLines(text) = %q, %v, want exp, error`, msg, err)
	}
}

func TestTruncateLinesFromAFile(t *testing.T) {
	content, err := os.ReadFile("test_data/file1.txt")
	if err != nil {
		t.Fatalf(`Reading test file failed: %v`, err)
	}
	exp := "1st l\n2nd l\n"
	msg, err := TruncateLines(string(content))
	if msg != exp || err != nil {
		t.Fatalf(`TruncateLines(text) = %q, %v, want exp, error`, msg, err)
	}
}
