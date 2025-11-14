package solutions

// Function Breakdown:
//
// Index(s string, toFind string) int:
//   - Finds first occurrence of toFind in s
//   - Returns the index position (0-based)
//   - Returns -1 if not found
//   - Returns 0 if toFind is empty
//
// Approach:
//   - Handle edge case: empty toFind → return 0
//   - Handle edge case: toFind longer than s or s empty → return -1
//   - Loop through s up to position where toFind could fit
//   - At each position, check if substring matches toFind
//   - If match, return current index
//   - If loop completes, return -1 (not found)
//
// Key Concepts:
//   - String slicing: s[i:i+len(toFind)] extracts substring
//   - Substring comparison: Direct string equality
//   - Loop bounds: i <= len(s)-len(toFind) ensures valid slicing
//   - Early return: Stop on first match
//
// Loop bounds explanation:
//   - If s = "Hello" (length 5) and toFind = "lo" (length 2)
//   - Last valid position is 3: "Hello"[3:5] = "lo"
//   - So loop: i <= 5-2, i <= 3
//
// Examples:
//   Index("Hello!", "l") = 2 (first 'l' at index 2)
//   Index("Salut!", "alu") = 1 (substring starts at 1)
//   Index("Ola!", "hOl") = -1 (not found, case-sensitive)
//   Index("test", "") = 0 (empty string found at start)
//   Index("", "test") = -1 (empty source string)

// Status: Required

func Index(s string, toFind string) int {
	// Edge case: empty substring always found at position 0
	if toFind == "" {
		return 0
	}

	// Edge case: can't find if toFind is longer than s, or s is empty
	if (len(toFind) > len(s)) || s == "" {
		return -1
	}

	// Loop through s, checking each possible position
	// Stop when there's no room left for toFind
	for i := 0; i <= len(s)-len(toFind); i++ {
		// Check if substring at position i matches toFind
		if s[i:i+len(toFind)] == toFind {
			return i // Found! Return index
		}
	}

	// Not found
	return -1
}
