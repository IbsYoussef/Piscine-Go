package solutions

// Function Breakdown:
//
// NRune(s string, n int) rune:
//   - Takes a string and position n (1-indexed)
//   - Returns the nth rune of the string
//   - Returns 0 if n is invalid
//
// Approach:
//   - Convert string to rune slice
//   - Validate n (must be > 0 and <= length)
//   - If invalid, return rune(0)
//   - If valid, return rune at index n-1
//
// Key Concepts:
//   - 1-indexed input: First rune is n=1
//   - 0-indexed array: First element is index 0
//   - Conversion: n (1-indexed) → n-1 (0-indexed)
//   - Validation: Prevent out-of-bounds access
//
// Validation logic:
//   - n <= 0: Invalid (no negative or zero positions)
//   - n > len(runes): Out of bounds
//   - Both cases return rune(0)
//
// Index conversion:
//   - n=1 → index 0 (first rune)
//   - n=2 → index 1 (second rune)
//   - n=3 → index 2 (third rune)
//   - Formula: index = n - 1
//
// Examples:
//   NRune("Hello!", 3) returns 'l' (index 2)
//   NRune("Salut!", 2) returns 'a' (index 1)
//   NRune("Bye!", -1) returns 0 (invalid)
//   NRune("Bye!", 5) returns 0 (out of bounds)
//   NRune("Ola!", 4) returns '!' (index 3)

// Status: Required

func NRune(s string, n int) rune {
	// Convert string to rune slice
	runes := []rune(s)

	// Validate n: must be positive and within bounds
	if n <= 0 || n > len(runes) {
		return rune(0)
	}

	// Return nth rune (convert 1-indexed to 0-indexed)
	return runes[n-1]
}
