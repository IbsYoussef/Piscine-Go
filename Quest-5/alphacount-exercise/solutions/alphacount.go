package solutions

// Function Breakdown:
//
// AlphaCount(s string) int:
//   - Takes a string as parameter
//   - Counts only alphabetic characters (A-Z, a-z)
//   - Returns the total count
//
// Approach:
//   - Initialize counter to 0
//   - Convert string to rune slice
//   - Loop through each rune
//   - Check if rune is alphabetic
//   - Increment counter if valid
//   - Return the count
//
// Key Concepts:
//   - ASCII ranges: 'A'-'Z' (65-90), 'a'-'z' (97-122)
//   - Character validation: Using range checks
//   - Filtering: Ignoring non-letters
//   - Accumulator: Building up count
//
// Validation logic:
//   - Uppercase check: rune >= 'A' && rune <= 'Z'
//   - Lowercase check: rune >= 'a' && rune <= 'z'
//   - Combined with OR: (uppercase) || (lowercase)
//
// Examples:
//   AlphaCount("Hello 78 World!    4455 /") = 10
//   AlphaCount("abc123") = 3
//   AlphaCount("123!@#") = 0
//   AlphaCount("") = 0
//   AlphaCount("GoLang") = 6

// Status: Required

func AlphaCount(s string) int {
	count := 0

	// Convert string to rune slice
	runes := []rune(s)

	// Loop through each rune
	for i := 0; i < len(runes); i++ {
		// Check if rune is uppercase or lowercase letter
		if (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z') {
			count++
		}
	}

	return count
}
