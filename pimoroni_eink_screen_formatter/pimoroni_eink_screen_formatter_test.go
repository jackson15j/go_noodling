package main

import (
	"os"
	"testing"
)

func TestTruncate(t *testing.T) {
	line := "A really long line."
	line_length := 5
	exp := line[:line_length]
	msg, err := Truncate(line, line_length)
	if msg != exp || err != nil {
		t.Fatalf(`Truncate(line) = %q, %v, want exp, error`, msg, err)
	}
}

func TestTruncateLines(t *testing.T) {
	text := "First long line.\nSecond long line."
	exp := "First\nSecon"
	line_length := 5
	msg, err := TruncateLines(text, line_length)
	if msg != exp || err != nil {
		t.Fatalf(`TruncateLines(text) = %q, %v, want exp, error`, msg, err)
	}
}

func TestTruncateLinesHandleEmptyLines(t *testing.T) {
	text := "First long line.\n\n\n"
	exp := "First\n\n\n"
	line_length := 5
	msg, err := TruncateLines(text, line_length)
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
	line_length := 5
	msg, err := TruncateLines(string(content), line_length)
	if msg != exp || err != nil {
		t.Fatalf(`TruncateLines(text) = %q, %v, want exp, error`, msg, err)
	}
}
