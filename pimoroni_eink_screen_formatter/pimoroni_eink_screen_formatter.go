package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	line_length := 5
	fmt.Println("Hello, World!")
	fmt.Println(Truncate("Hello, World!", line_length))
	TruncateFilesInFolder("test_data", 50, 10, filepath.Join("build", "truncated"))
}

// TODO: read/write files.

// Truncate the input string.
func Truncate(line string, line_length int) (string, error) {
	if line_length == 0 {
		line_length = 5
	}
	if len(line) == 0 {
		// Avoid runtime slice error if line is empty. Just return.
		return line, nil
	}
	if line_length > len(line) {
		// Avoid runtime slice error if slicing greater than line length. Just
		// return.
		return line, nil
	}
	return line[:line_length], nil
}

// Return a new multi-line string with each line truncated.
func TruncateLines(text string, line_length int, max_lines int) (string, error) {
	lines, err := splitLines(text)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	truncated_lines := []string{}
	for i, line := range lines {
		if max_lines != 0 && i == max_lines {
			break
		}
		truncated_line, err := Truncate(line, line_length)
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

func TruncateFilesInFolder(dir string, line_length int, max_lines int, output string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := os.MkdirAll(output, 0770); err != nil {
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
		truncated_text, err := TruncateLines(string(content), line_length, max_lines)
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("Truncated content for: %s: %s\n", path, truncated_text)

		// Write truncated text to a file.
		tmpfile, err := os.Create(filepath.Join(output, file.Name()))
		if err != nil {
			log.Fatal(err)
			return err
		}
		tmpfile.WriteString(truncated_text)
		fmt.Printf("Wrote truncated content to: %v\n", tmpfile.Name())
	}

	return nil
}
