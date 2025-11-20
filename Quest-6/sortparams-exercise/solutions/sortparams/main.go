package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program Breakdown:
//
// Goal: Print command-line arguments in ASCII sorted order
//
// Approach:
//   - Get arguments from os.Args[1:]
//   - Sort arguments using bubble sort
//   - Print each sorted argument
//
// Key Concepts:
//   - Bubble sort: Compare adjacent elements and swap
//   - String comparison: Go can compare strings directly
//   - ASCII order: numbers < uppercase < lowercase
//
// Bubble Sort Algorithm:
//   - Outer loop: Number of passes
//   - Inner loop: Compare adjacent pairs
//   - Swap if out of order
//   - Repeat until sorted
//
// Examples:
//   Input: ./sortparams b a c
//   After sort: [a, b, c]
//   Output:
//   a
//   b
//   c
//
//   Input: ./sortparams 1 a 2 A
//   After sort: [1, 2, A, a]
//   Output:
//   1
//   2
//   A
//   a

// Status: Optional

func main() {
	// Get all arguments except program name
	args := os.Args[1:]

	// Bubble sort: sort arguments in ASCII order
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-i-1; j++ {
			// Compare adjacent strings
			if args[j] > args[j+1] {
				// Swap if out of order
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Print each sorted argument
	for _, str := range args {
		// Print each character
		for _, v := range str {
			z01.PrintRune(v)
		}
		// Print newline after each argument
		z01.PrintRune('\n')
	}
}
