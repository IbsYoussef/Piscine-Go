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
	fullPath := os.Args[0]
	lastSlashIndex := -1

	// Find the last slash in the path
	for i, r := range fullPath {
		if r == '/' {
			lastSlashIndex = i
		}
	}

	// Extract program name (everything after last slash)
	var programName string
	if lastSlashIndex != -1 {
		// Has slash: get substring after it
		programName = fullPath[lastSlashIndex+1:]
	} else {
		// No slash: use entire path
		programName = fullPath
	}

	// Print each character of the program name
	for _, v := range programName {
		z01.PrintRune(v)
	}

	// Print newline
	z01.PrintRune('\n')
}
