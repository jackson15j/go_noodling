package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Truncate("Hello, World!"))
}

// Truncate the input string.
func Truncate(line string) (string, error) {
	return line[:5], nil
}
