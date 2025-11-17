package solutions

// Function Breakdown:
//
// IsNumeric(s string) bool:
//   - Checks if ALL characters are numeric digits
//   - Returns true only if all are 0-9
//   - Returns false for empty string or any non-digit
//
// Approach:
//   - Handle edge case: empty string → return false
//   - Convert string to rune slice
//   - Loop through each rune
//   - If ANY rune is not a digit, return false immediately
//   - If loop completes, all are digits → return true
//
// Key Concepts:
//   - All-or-nothing validation: One failure = false
//   - Early return: Stop on first non-digit
//   - Digit range: '0' (48) to '9' (57)
//   - Negation logic: Check for NOT digit
//
// Validation logic:
//   - Use negation: if !(rune >= '0' && rune <= '9')
//   - This catches:
//     * Letters (uppercase/lowercase)
//     * Symbols (commas, periods, etc.)
//     * Spaces
//     * Any non-digit character
//
// Examples:
//   IsNumeric("010203") = true (only digits)
//   IsNumeric("01,02,03") = false (commas not digits)
//   IsNumeric("123abc") = false (letters not digits)
//   IsNumeric("12 34") = false (space not digit)
//   IsNumeric("") = false (empty string)
//   IsNumeric("0") = true (single digit)

// Status: Required

func IsNumeric(s string) bool {
	// Edge case: empty string returns false
	if s == "" {
		return false
	}

	// Convert string to rune slice
	runes := []rune(s)

	// Check each character
	for i := 0; i < len(runes); i++ {
		// If NOT a digit, return false immediately
		if !(runes[i] >= '0' && runes[i] <= '9') {
			return false
		}
	}

	// All characters are digits
	return true
}
