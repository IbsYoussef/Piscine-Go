package solutions

import "github.com/01-edu/z01"

// Function Breakdown:
//
// PrintNbrBase(nbr int, base string):
//   - Converts integer to specified base
//   - Prints result or "NV" if invalid
//   - Handles negative numbers
//
// Approach:
//   - Validate base using helper function
//   - If invalid, print "NV" and return
//   - Handle negative: print '-', convert absolute value
//   - Handle zero: print first character of base
//   - Convert using modulo and division
//   - Build result in reverse order
//   - Print result forwards
//
// Key Concepts:
//   - Base validation: length, uniqueness, no +/-
//   - Base conversion: repeated division
//   - Reverse accumulation: build backwards, print forwards
//
// Validation rules:
//   - Base length >= 2
//   - All characters unique
//   - No '+' or '-' characters
//
// Conversion algorithm:
//   1. Divide number by base length
//   2. Remainder gives digit position in base
//   3. Quotient becomes new number
//   4. Repeat until quotient is 0
//   5. Read digits in reverse order
//
// Examples:
//   PrintNbrBase(125, "0123456789") prints "125"
//   PrintNbrBase(-125, "01") prints "-1111101"
//   PrintNbrBase(125, "0123456789ABCDEF") prints "7D"
//   PrintNbrBase(-125, "choumi") prints "-uoi"
//   PrintNbrBase(125, "aa") prints "NV"

// Status: Bonus

func PrintNbrBase(nbr int, base string) {
	// Validate base
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	// Handle negative numbers
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	baseLen := len(base)

	// Handle zero
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// Create a slice for the result
	var result []rune

	// Conversion loop: extract digits in reverse order
	for nbr > 0 {
		remainder := nbr % baseLen
		result = append(result, rune(base[remainder]))
		nbr = nbr / baseLen
	}

	// Print the result in reverse order (correct order)
	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}

// Helper function to validate base
func isValidBase(base string) bool {
	// Base must have at least 2 characters
	if len(base) < 2 {
		return false
	}

	// Check for duplicates and invalid characters
	seen := make(map[rune]bool)
	for _, r := range base {
		// Check for '+' or '-'
		if r == '+' || r == '-' {
			return false
		}
		// Check for duplicates
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
