package solutions

// Function Breakdown:
//
// SplitWhiteSpaces(s string) []string:
//   - Splits string into words separated by whitespace
//   - Whitespace: space, tab (\t), newline (\n)
//   - Returns slice of words
//
// Approach:
//   - Initialize empty result slice and word builder
//   - Loop through each character
//   - If whitespace: save current word (if not empty), reset word
//   - If not whitespace: add to current word
//   - After loop: add last word if exists
//   - Return result
//
// Key Concepts:
//   - Word building: Accumulate characters until separator
//   - Whitespace detection: char == ' ' || char == '\t' || char == '\n'
//   - Edge case handling: Multiple consecutive whitespace
//   - Final word: Don't forget to add last word after loop
//
// Algorithm:
//   1. Start with empty word and empty result
//   2. For each character:
//      - If whitespace and word not empty: save word, reset
//      - If not whitespace: add to word
//   3. After loop: save last word if not empty
//
// Examples:
//   "Hello how" → ["Hello", "how"]
//   "a  b" → ["a", "b"] (double space ignored)
//   "  trim  " → ["trim"] (leading/trailing whitespace ignored)

// Status: Required

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""

	// Loop through each character
	for _, char := range s {
		// Check if whitespace (space, tab, or newline)
		if char == ' ' || char == '\t' || char == '\n' {
			// If we have a word, save it
			if word != "" {
				result = append(result, word)
				word = "" // Reset for next word
			}
		} else {
			// Build current word
			word += string(char)
		}
	}

	// Add the last word if there is one
	if word != "" {
		result = append(result, word)
	}

	return result
}
