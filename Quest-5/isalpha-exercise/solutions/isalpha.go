package solutions

// Function Breakdown:
//
// IsAlpha(s string) bool:
//   - Checks if ALL characters are alphanumeric
//   - Alphanumeric = letters (A-Z, a-z) or digits (0-9)
//   - Returns true for empty string
//   - Returns false if any character is not alphanumeric
//
// Approach:
//   - Convert string to rune slice
//   - Loop through each rune
//   - Check if rune is NOT alphanumeric
//   - If not alphanumeric, return false
//   - If loop completes, all are alphanumeric → return true
//
// Key Concepts:
//   - Multiple range checks: uppercase, lowercase, digits
//   - Logical negation: Check for NOT alphanumeric
//   - Short-circuit evaluation: Return false on first failure
//   - Empty string: Returns true (no invalid characters)
//
// Validation logic:
//   - NOT uppercase: !(rune >= 'A' && rune <= 'Z')
//   - AND NOT lowercase: !(rune >= 'a' && rune <= 'z')
//   - AND NOT digit: !(rune >= '0' && rune <= '9')
//   - If all three are true, character is invalid
//
// De Morgan's Law application:
//   - NOT (uppercase OR lowercase OR digit)
//   - Equals: (NOT uppercase) AND (NOT lowercase) AND (NOT digit)
//
// Examples:
//   IsAlpha("HelloHowareyou") = true (only letters)
//   IsAlpha("Whatsthis4") = true (letters and digit)
//   IsAlpha("Hello! How are you?") = false (space and !)
//   IsAlpha("What's this 4?") = false (apostrophe, space, ?)
//   IsAlpha("") = true (empty string)
//   IsAlpha("123") = true (only digits)

// Status: Required

func IsAlpha(s string) bool {
	// Convert string to rune slice
	runes := []rune(s)

	// Check each character
	for i := 0; i < len(runes); i++ {
		// Check if NOT uppercase AND NOT lowercase AND NOT digit
		if !(runes[i] >= 'A' && runes[i] <= 'Z') &&
			!(runes[i] >= 'a' && runes[i] <= 'z') &&
			!(runes[i] >= '0' && runes[i] <= '9') {
			return false // Found non-alphanumeric character
		}
	}

	// All characters are alphanumeric (or string is empty)
	return true
}
