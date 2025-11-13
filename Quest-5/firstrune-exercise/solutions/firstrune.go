package solutions

// Function Breakdown:
//
// FirstRune(s string) rune:
//   - Takes a string as parameter
//   - Returns the first rune (character) of the string
//   - Handles UTF-8 encoded strings properly
//
// Approach:
//   - Convert string to rune slice
//   - Return the first element (index 0)
//
// Key Concepts:
//   - Rune: Go's type for Unicode code points (int32)
//   - String to rune conversion: []rune(s)
//   - Indexing: Accessing first element with [0]
//   - UTF-8: Proper handling of multi-byte characters
//
// Why convert to rune slice?
//   - Strings in Go are UTF-8 encoded bytes
//   - Direct indexing (s[0]) gives a byte, not a character
//   - Converting to []rune ensures correct character handling
//   - Important for non-ASCII characters (é, 中, etc.)
//
// Alternative approach:
//   - Use for range loop: for _, r := range s { return r }
//   - Range automatically handles rune conversion
//   - Returns first rune directly
//
// Examples:
//   FirstRune("Hello!") returns 'H'
//   FirstRune("Salut!") returns 'S'
//   FirstRune("中文") returns '中' (Chinese character)

// Status: Required

func FirstRune(s string) rune {
	// Convert string to rune slice
	runes := []rune(s)

	// Return first rune
	return runes[0]
}
