package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program Breakdown:
//
// Goal: Print the program's executable name
//
// Approach:
//   - Get full path from os.Args[0]
//   - Extract just the filename (after last '/')
//   - Print the filename using z01.PrintRune
//   - Add newline at the end
//
// Key Concepts:
//   - os.Args[0]: Contains full path to executable
//   - Path parsing: Find last '/' to get filename
//   - String iteration: Print each character
//
// Algorithm:
//   1. Get full path: os.Args[0]
//   2. Find last slash position
//   3. Extract substring after last slash
//   4. If no slash, use entire path
//   5. Print each character
//   6. Print newline
//
// Examples:
//   Path: "/usr/bin/program" → Print: "program"
//   Path: "./main" → Print: "main"
//   Path: "test" → Print: "test"

// Status: Required

func main() {
	// Get the full path of program
	args := os.Args[0]

	// Find the last slash in the path
	slashPos := -1
	for i, char := range args {
		if char == '/' {
			slashPos = i
		}
	}

	// Extract the program name
	programName := ""
	if slashPos != -1 {
		// If a slash is present get the substring after it
		programName = args[slashPos+1:]
	} else {
		// If no slash is present use the entire path
		programName = args
	}

	// Print each character of the program name
	for _, v := range programName {
		z01.PrintRune(v)
	}

	// Print a newline
	z01.PrintRune('\n')
}
