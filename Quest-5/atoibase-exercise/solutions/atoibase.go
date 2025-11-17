package solutions

// Function Breakdown:
//
// AtoiBase(s string, base string) int:
//   - Converts string number in given base to decimal integer
//   - Returns 0 if base is invalid
//   - Does not handle negative numbers
//
// Approach:
//   - Validate base (same rules as PrintNbrBase)
//   - Check if all characters in s exist in base
//   - Use Horner's method to convert
//   - Formula: result = result * baseLen + digitValue
//
// Key Concepts:
//   - Base validation: length >= 2, unique chars, no +/-
//   - Character lookup: Find position of char in base string
//   - Horner's method: Efficient base conversion
//   - Position = digit value
//
// Horner's method example:
//   "125" in base 10:
//   result = 0
//   '1': result = 0*10 + 1 = 1
//   '2': result = 1*10 + 2 = 12
//   '5': result = 12*10 + 5 = 125
//
// Another example:
//   "7D" in base "0123456789ABCDEF" (hex):
//   '7' is at index 7: result = 0*16 + 7 = 7
//   'D' is at index 13: result = 7*16 + 13 = 112 + 13 = 125
//
// Examples:
//   AtoiBase("125", "0123456789") = 125
//   AtoiBase("1111101", "01") = 125 (binary)
//   AtoiBase("7D", "0123456789ABCDEF") = 125 (hex)
//   AtoiBase("uoi", "choumi") = 125 (custom base)
//   AtoiBase("bbbbbab", "-ab") = 0 (invalid base)

// Status: Bonus

func AtoiBase(s string, base string) int {
	// Validate base
	if !isValidBase(base) {
		return 0
	}

	// Check if all characters in s are in base
	for _, char := range s {
		if !contains(base, char) {
			return 0
		}
	}

	result := 0
	baseLen := len(base)

	// Convert using Horner's method
	for _, char := range s {
		// Find position of character in base (this is the digit value)
		digitValue := indexOf(base, char)

		// Horner's method: multiply by base, add digit
		result = result*baseLen + digitValue
	}

	return result
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

// Helper function to check if base contains character
func contains(base string, char rune) bool {
	for _, b := range base {
		if b == char {
			return true
		}
	}
	return false
}

// Helper function to find index of character in base string
func indexOf(base string, char rune) int {
	for i, b := range base {
		if b == char {
			return i
		}
	}
	return -1
}
