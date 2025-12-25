package solutions

// Rot14 rotates each letter in the string by 14 positions in the alphabet
// Non-letter characters remain unchanged
// Case is preserved (uppercase stays uppercase, lowercase stays lowercase)
//
// Algorithm:
// 1. Check if character is a letter
// 2. Convert letter to position (0-25)
// 3. Add 14 and wrap using modulo 26
// 4. Convert back to character
//
// Formula: (character - base + 14) % 26 + base
//   - Subtract base ('a' or 'A') to get position (0-25)
//   - Add 14 for rotation
//   - Modulo 26 to wrap around
//   - Add base back to convert to ASCII character
func Rot14(s string) string {
	result := ""

	for _, r := range s {
		var newChar rune

		if r >= 'a' && r <= 'z' {
			// Lowercase letter rotation
			// Convert to position (0-25), rotate by 14, wrap, convert back
			newChar = (r-'a'+14)%26 + 'a'
		} else if r >= 'A' && r <= 'Z' {
			// Uppercase letter rotation
			// Same logic, different base
			newChar = (r-'A'+14)%26 + 'A'
		} else {
			// Not a letter - keep as is
			newChar = r
		}

		result += string(newChar)
	}

	return result
}
