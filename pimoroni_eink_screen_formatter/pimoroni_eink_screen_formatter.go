package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Truncate("Hello, World!"))
}

// TODO: read/write files.

// Truncate the input string.
func Truncate(line string) (string, error) {
	return line[:5], nil
}

// Return a new multi-line string with each line truncated.
func TruncateLines(text string) (string, error) {
	lines, err := splitLines(text)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	truncated_lines := []string{}
	for i, line := range lines {
		fmt.Printf("index: %d, line: %s\n", i, line)
		truncated_line, err := Truncate(line)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		truncated_lines = append(truncated_lines, truncated_line)
	}
	truncated_text, err := joinLines(truncated_lines)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return truncated_text, nil
}

func splitLines(text string) ([]string, error) { return strings.Split(text, "\n"), nil }
func joinLines(lines []string) (string, error) { return strings.Join(lines[:], "\n"), nil }
