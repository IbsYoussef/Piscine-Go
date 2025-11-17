package solutions

// Function Breakdown:
//
// Capitalize(s string) string:
//   - Capitalizes first letter of each word
//   - Lowercases all other letters
//   - Words are alphanumeric sequences
//   - Returns the modified string
//
// Approach:
//   - Use boolean flag to track if inside a word
//   - Convert string to rune slice
//   - Loop through each character
//   - First letter of word: capitalize
//   - Subsequent letters: lowercase
//   - Digits: part of word (no case change)
//   - Non-alphanumeric: word boundary
//
// Key Concepts:
//   - State machine: Track word/non-word state
//   - Case conversion: Use ASCII arithmetic
//   - Word detection: Alphanumeric characters
//   - Boundary detection: Non-alphanumeric characters
//
// State logic:
//   - inWord = false: Looking for start of word
//   - inWord = true: Inside a word
//   - First letter found: Capitalize and set inWord = true
//   - Non-letter within word: Lowercase if uppercase
//   - Non-alphanumeric: Reset to inWord = false (or check)
//
// Edge cases:
//   - Numbers are part of words: "4you" → "4you" (no capitalize)
//   - Multiple spaces/symbols: Each word treated separately
//   - Already capitalized: Lowercase non-first letters
//
// Examples:
//   Capitalize("Hello! How are you?") = "Hello! How Are You?"
//   Capitalize("how+are+things+4you?") = "How+Are+Things+4you?"
//   Capitalize("hello") = "Hello"
//   Capitalize("HELLO WORLD") = "Hello World"

// Status: Required

func Capitalize(s string) string {
	inWord := false
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' && !inWord {
			// Capitalize the first lowercase letter of a word
			runes[i] -= 32
			inWord = true
		} else if runes[i] >= 'A' && runes[i] <= 'Z' && inWord {
			// Lowercase subsequent uppercase letters within a word
			runes[i] += 32
		} else if runes[i] >= '0' && runes[i] <= '9' {
			// Digits are part of words
			inWord = true
		} else {
			// Check if we're still in a word (alphanumeric)
			// or if this is a word boundary
			inWord = (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z')
		}
	}
	return string(runes)
}
