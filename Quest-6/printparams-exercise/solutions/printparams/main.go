package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program Breakdown:
//
// Goal: Print all command-line arguments (except program name)
//
// Approach:
//   - Get arguments from os.Args[1:] (skip program name)
//   - Loop through each argument
//   - For each argument, print each character
//   - Print newline after each argument
//
// Key Concepts:
//   - os.Args[0]: Program name (skip this)
//   - os.Args[1:]: All arguments after program name
//   - Nested loops: Outer for arguments, inner for characters
//
// Examples:
//   Input: ./printparams hello world
//   Output:
//   hello
//   world
//
//   Input: ./printparams one two three
//   Output:
//   one
//   two
//   three

// Status: Required

func main() {
	// Get all arguments except program name
	args := os.Args[1:]

	// Loop through each argument in the slice
	for _, arguements := range args {
		// Print each character of the argument
		for _, r := range arguements {
			z01.PrintRune(r)
		}
		// Print newline after each argument
		z01.PrintRune('\n')
	}
}
