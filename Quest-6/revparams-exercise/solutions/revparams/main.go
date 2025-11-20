package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program Breakdown:
//
// Goal: Print all command-line arguments in REVERSE order
//
// Approach:
//   - Get arguments from os.Args[1:]
//   - Loop backwards through arguments
//   - For each argument, print each character
//   - Print newline after each argument
//
// Key Concepts:
//   - Reverse iteration: Start at end, go to beginning
//   - Index-based loop: for i := len(args)-1; i >= 0; i--
//   - Nested loops: Outer for arguments, inner for characters
//
// Examples:
//   Input: ./revparams hello world
//   Output:
//   world
//   hello
//
//   Input: ./revparams one two three
//   Output:
//   three
//   two
//   one

// Status: Required

func main() {
	// Get all arguments except program name
	args := os.Args[1:]

	// Loop backwards through arguments
	for i := len(args) - 1; i >= 0; i-- {
		// Print each character of the argument
		for _, char := range args[i] {
			z01.PrintRune(char)
		}
		// Print newline after each argument
		z01.PrintRune('\n')
	}
}
