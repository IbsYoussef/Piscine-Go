package solutions

// Function Breakdown:
//
// IsLower(s string) bool:
//   - Checks if ALL characters are lowercase letters
//   - Returns true only if all are a-z
//   - Returns false for empty string or any non-lowercase
//
// Approach:
//   - Handle edge case: empty string → return false
//   - Convert string to rune slice
//   - Loop through each rune
//   - If ANY rune is not lowercase, return false immediately
//   - If loop completes, all are lowercase → return true
//
// Key Concepts:
//   - All-or-nothing validation: One failure = false
//   - Early return: Stop on first non-lowercase
//   - Lowercase range: 'a' (97) to 'z' (122)
//   - Negation logic: Check for NOT lowercase
//
// Validation logic:
//   - Use negation: if !(rune >= 'a' && rune <= 'z')
//   - This catches:
//     * Uppercase letters
//     * Numbers
//     * Symbols
//     * Spaces
//     * Any non-lowercase character
//
// Examples:
//   IsLower("hello") = true (all lowercase)
//   IsLower("hello!") = false ('!' is not lowercase)
//   IsLower("Hello") = false ('H' is uppercase)
//   IsLower("abc123") = false (numbers not lowercase)
//   IsLower("") = false (empty string)
//   IsLower("a") = true (single lowercase)

// Status: Required

func IsLower(s string) bool {
	// Edge case: empty string returns false
	if s == "" {
		return false
	}

	// Convert string to rune slice
	runes := []rune(s)

	// Check each character
	for i := 0; i < len(runes); i++ {
		// If NOT lowercase, return false immediately
		if !(runes[i] >= 'a' && runes[i] <= 'z') {
			return false
		}
	}

	// All characters are lowercase
	return true
}
