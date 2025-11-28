package solutions

// Function Breakdown:
//
// ConvertBase(nbr, baseFrom, baseTo string) string:
//   - Converts number from one base to another
//   - Two-step process: baseFrom → decimal → baseTo
//
// Approach:
//   - Step 1: Convert nbr from baseFrom to decimal (like AtoiBase)
//   - Step 2: Convert decimal to baseTo (like PrintNbrBase)
//   - Return result as string
//
// Key Concepts:
//   - AtoiBase: String in base → decimal integer
//   - PrintNbrBase: Decimal integer → string in base
//   - Combining both: Complete base conversion
//
// Algorithm:
//   1. Validate bases (optional, problem says valid bases only)
//   2. Convert nbr from baseFrom to decimal integer
//   3. Convert decimal integer to baseTo
//   4. Return result
//
// Examples:
//   ConvertBase("101011", "01", "0123456789")
//     Step 1: "101011" in binary = 43 in decimal
//     Step 2: 43 in decimal = "43" in decimal
//     Result: "43"
//
//   ConvertBase("2A", "0123456789ABCDEF", "01")
//     Step 1: "2A" in hex = 42 in decimal
//     Step 2: 42 in decimal = "101010" in binary
//     Result: "101010"

// Status: Bonus

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Step 1: Convert from baseFrom to decimal
	decimal := atoiBase(nbr, baseFrom)

	// Step 2: Convert from decimal to baseTo
	result := printNbrBase(decimal, baseTo)

	return result
}

// Helper: Convert string in given base to decimal integer
// Same as AtoiBase from Quest 5
func atoiBase(s, base string) int {
	result := 0
	baseLen := len(base)

	// Convert each character
	for _, char := range s {
		// Find position of character in base (= digit value)
		digitValue := indexOf(base, char)

		// Build number using Horner's method
		result = result*baseLen + digitValue
	}

	return result
}

// Helper: Convert decimal integer to string in given base
// Same as PrintNbrBase from Quest 5 (but returns string instead of printing)
func printNbrBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}

	result := ""
	baseLen := len(base)

	// Extract digits in reverse order
	for n > 0 {
		remainder := n % baseLen
		result = string(base[remainder]) + result
		n = n / baseLen
	}

	return result
}

// Helper: Find index of character in string
func indexOf(s string, char rune) int {
	for i, c := range s {
		if c == char {
			return i
		}
	}
	return -1
}
