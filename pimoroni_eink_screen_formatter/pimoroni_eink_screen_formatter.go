package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Truncate("Hello, World!"))
	TruncateFilesInFolder("test_data")
}

// TODO: read/write files.

// Truncate the input string.
func Truncate(line string) (string, error) {
	if len(line) == 0 {
		// Avoid runtime slice error if line is empty. Just return.
		return line, nil
	}
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
	for _, line := range lines {
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

func TruncateFilesInFolder(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		// TODO: collect errors into a slice report at the end.
		fmt.Printf("Truncating: %s...\n", path)
		content, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
			return err
		}
		truncated_text, err := TruncateLines(string(content))
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("Truncated content for: %s: %s\n", path, truncated_text)
		// TODO: Write truncated text to a file.
	}
	return nil
}
