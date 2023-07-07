package main

import (
	"fmt"
	"go_noodling/udemy/examples"
)

func main() {
	// Emacs emoji picker: `M-x emoji-list`, then `<return>` to pick.
	fmt.Println("Hello World 👋")
	fmt.Println(`
Raw string literal

Backticks are used to do raw strings.
    This keeps the formatting (including tabs/spaces).`)

	// Formatting strings. See: https://pkg.go.dev/fmt#hdr-Printing.
	s := []int{0, 42, 999}
	for _, x := range s {
		fmt.Printf("Value: %v, Binary: %b, Hex: %x (%#x).\n", x, x, x, x)
	}
	fmt.Printf("Boolean - Value: %v, Text: %t.\n", true, true)

	// Bitwise shifting.
	v := 1 << 10
	fmt.Printf("Bitwise shift: `1 << 10`. Equals: %v.\n", v)

	// iota module examples.
	examples.Iota()
}
