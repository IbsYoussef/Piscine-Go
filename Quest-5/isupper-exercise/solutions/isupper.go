package solutions

// Function Breakdown:
//
// IsUpper(s string) bool:
//   - Checks if ALL characters are uppercase letters
//   - Returns true only if all are A-Z
//   - Returns false for empty string or any non-uppercase
//
// Approach:
//   - Handle edge case: empty string → return false
//   - Convert string to rune slice
//   - Loop through each rune
//   - If ANY rune is not uppercase, return false immediately
//   - If loop completes, all are uppercase → return true
//
// Key Concepts:
//   - All-or-nothing validation: One failure = false
//   - Early return: Stop on first non-uppercase
//   - Uppercase range: 'A' (65) to 'Z' (90)
//   - Negation logic: Check for NOT uppercase
//
// Validation logic:
//   - Use negation: if !(rune >= 'A' && rune <= 'Z')
//   - This catches:
//     * Lowercase letters
//     * Numbers
//     * Symbols
//     * Spaces
//     * Any non-uppercase character
//
// Examples:
//   IsUpper("HELLO") = true (all uppercase)
//   IsUpper("HELLO!") = false ('!' is not uppercase)
//   IsUpper("Hello") = false ('e', 'l', 'o' are lowercase)
//   IsUpper("ABC123") = false (numbers not uppercase)
//   IsUpper("") = false (empty string)
//   IsUpper("A") = true (single uppercase)

// Status: Required

func IsUpper(s string) bool {
	// Edge case: empty string returns false
	if s == "" {
		return false
	}

	// Convert string to rune slice
	runes := []rune(s)

	// Check each character
	for i := 0; i < len(runes); i++ {
		// If NOT uppercase, return false immediately
		if !(runes[i] >= 'A' && runes[i] <= 'Z') {
			return false
		}
	}

	// All characters are uppercase
	return true
}
