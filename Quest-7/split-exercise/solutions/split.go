package solutions

// Function Breakdown:
//
// Split(s, sep string) []string:
//   - Splits string by custom separator
//   - Returns slice of substrings
//   - Handles multi-character separators
//
// Approach:
//   - Loop through the string
//   - Check if current position matches separator
//   - If match: save current word, skip separator
//   - If no match: add character to current word
//   - After loop: save last word if exists
//
// Key Concepts:
//   - Substring matching: Check if separator appears at current position
//   - String slicing: Extract parts of string
//   - Word building: Accumulate characters until separator
//
// Algorithm:
//   1. Initialize result slice and current word
//   2. Loop through string with index
//   3. Check if remaining string starts with separator
//   4. If yes: save word, skip separator length
//   5. If no: add character to word
//   6. After loop: save final word
//
// The Trick:
//   Use s[i:i+len(sep)] to check if separator appears at position i
//   This allows multi-character separator matching!
//
// Edge Cases:
//   - Empty separator: return whole string
//   - Separator not found: return whole string in slice
//   - Multiple consecutive separators: creates empty strings
//
// Examples:
//   Split("HelloHAhowHAareHAyou?", "HA")
//     → ["Hello", "how", "are", "you?"]
//   Split("a,b,c", ",")
//     → ["a", "b", "c"]
//   Split("hello", "X")
//     → ["hello"]
//   Split("one##two##three", "##")
//     → ["one", "two", "three"]

// Status: Optional

func Split(s, sep string) []string {
	result := []string{}
	word := ""

	// Edge case: empty separator returns whole string
	if len(sep) == 0 {
		return []string{s}
	}

	// Loop through string with manual index control
	for i := 0; i < len(s); {
		// Check if separator appears at current position
		// s[i:i+len(sep)] extracts substring to compare with separator
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			// Found separator: save current word
			result = append(result, word)
			word = "" // Reset for next word
			// Skip past the separator
			i += len(sep)
		} else {
			// No separator: add character to current word
			word += string(s[i])
			i++
		}
	}

	// Add the last word (always exists, even if empty)
	result = append(result, word)

	return result
}
