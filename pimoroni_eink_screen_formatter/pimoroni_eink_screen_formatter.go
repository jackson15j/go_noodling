package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Truncate("Hello, World!"))
}
// TODO: don't swallow errors!
// TODO: read/write files.


// Truncate the input string.
func Truncate(line string) (string, error) {
	return line[:5], nil
}

// Return a new multi-line string with each line truncated.
func TruncateLines(text string) (string, error) {
	lines, _ := splitLines(text)
	truncated_lines := []string{}
	for i, line := range lines {
		fmt.Printf("index: %d, line: %s\n", i, line)
		truncated_line, _ := Truncate(line)
		truncated_lines = append(truncated_lines, truncated_line)
	}
	truncated_text, _ := joinLines(truncated_lines)
	return truncated_text, nil
}

func splitLines(text string) ([]string, error) { return strings.Split(text, "\n"), nil }
func joinLines(lines []string) (string, error) { return strings.Join(lines[:], "\n"), nil }
