package solutions

// Function Breakdown:
//
// Compare(a, b string) int:
//   - Compares two strings lexicographically
//   - Returns 0 if equal
//   - Returns -1 if a < b
//   - Returns 1 if a > b
//
// Approach:
//   - Handle edge cases (empty strings)
//   - Convert both to rune slices
//   - Compare character by character
//   - If characters differ, return based on comparison
//   - If one string is shorter, it's "less"
//   - If identical, return 0
//
// Key Concepts:
//   - Lexicographic order: Dictionary/alphabetical order
//   - Character comparison: Compare Unicode values
//   - Prefix handling: "abc" < "abcd"
//   - Edge cases: Empty strings
//
// Comparison rules:
//   - Empty string < any non-empty string
//   - Two empty strings are equal
//   - Compare rune by rune until difference found
//   - If prefix matches, shorter < longer
//
// Examples:
//   Compare("Hello!", "Hello!") = 0 (identical)
//   Compare("Salut!", "lut!") = -1 ('S' < 'l')
//   Compare("Ola!", "Ol") = 1 (longer when prefix matches)
//   Compare("abc", "abd") = -1 ('c' < 'd')
//   Compare("", "test") = -1 (empty < non-empty)

// Status: Required

func Compare(a, b string) int {
	// Edge case: both empty
	if a == "" && b == "" {
		return 0
	}
	// Edge case: a empty, b not
	if a == "" && b != "" {
		return -1
	}
	// Edge case: a not empty, b empty
	if a != "" && b == "" {
		return 1
	}

	// Convert to rune slices for proper comparison
	aRunes := []rune(a)
	bRunes := []rune(b)

	// Compare character by character
	for i := 0; i < len(aRunes) && i < len(bRunes); i++ {
		if aRunes[i] > bRunes[i] {
			return 1 // a comes after b
		}
		if aRunes[i] < bRunes[i] {
			return -1 // a comes before b
		}
	}

	// All characters matched so far, check lengths
	if len(aRunes) > len(bRunes) {
		return 1 // a is longer (e.g., "abc" > "ab")
	}

	if len(aRunes) < len(bRunes) {
		return -1 // a is shorter (e.g., "ab" < "abc")
	}

	// Completely identical
	return 0
}
