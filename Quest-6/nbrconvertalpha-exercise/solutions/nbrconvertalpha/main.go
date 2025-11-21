package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program Breakdown:
//
// Goal: Convert numbers to alphabet letters
//
// Approach:
//   - Check for --upper flag
//   - For each argument:
//     * Convert string to number (manually)
//     * Validate number (1-26)
//     * Convert to letter (1='a', 2='b', etc.)
//     * Apply uppercase if flag set
//     * Print letter or space
//   - Print newline
//
// Key Concepts:
//   - Flag parsing: Check first argument
//   - String to int: Manual conversion
//   - Number to letter: ASCII arithmetic
//   - Case conversion: 'a' → 'A' (subtract 32)
//
// Conversion formulas:
//   - Number to lowercase: 'a' + (n-1)
//   - Lowercase to uppercase: letter - 'a' + 'A'
//
// Examples:
//   Input: 8 5 12 12 15
//   Output: hello
//
//   Input: --upper 8 5 25
//   Output: HEY
//
//   Input: 32 86 h
//   Output: "   " (3 spaces)

// Status: Optional

func main() {
	// Create alphabet map (1='a', 2='b', ..., 26='z')
	alphabet := make(map[int]rune)
	for i := 1; i <= 26; i++ {
		alphabet[i] = rune('a' + i - 1)
	}

	isUpper := false
	args := os.Args[1:]

	// No arguments: exit without printing
	if len(args) == 0 {
		return
	}

	// Check for --upper flag
	if args[0] == "--upper" {
		isUpper = true
		args = args[1:] // Skip flag
	}

	// Process each argument
	for _, arg := range args {
		n := 0
		valid := true

		// Manual string to int conversion
		for _, r := range arg {
			if r < '0' || r > '9' {
				valid = false
				break
			}
			n = n*10 + int(r-'0')
		}

		// Validate: must be 1-26
		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ') // Invalid: print space
		} else {
			letter := alphabet[n]
			if isUpper {
				// Convert to uppercase
				letter = letter - 'a' + 'A'
			}
			z01.PrintRune(letter)
		}
	}

	z01.PrintRune('\n')
}
