package solutions

// Function Breakdown:
//
// LastRune(s string) rune:
//   - Takes a string as parameter
//   - Returns the last rune (character) of the string
//   - Handles UTF-8 encoded strings properly
//
// Approach:
//   - Convert string to rune slice
//   - Get the length of the rune slice
//   - Return the last element (index = length - 1)
//
// Key Concepts:
//   - Rune: Go's type for Unicode code points (int32)
//   - String to rune conversion: []rune(s)
//   - Length: len(runes) gives number of runes
//   - Last index: Always length - 1 (0-indexed)
//
// Why len(runes) - 1?
//   - Arrays and slices are 0-indexed
//   - If length is 5, indices are: 0, 1, 2, 3, 4
//   - Last element is at index 4 (5 - 1)
//
// Example walkthrough:
//   s = "Hello!"
//   runes = ['H', 'e', 'l', 'l', 'o', '!']
//   len(runes) = 6
//   Last index = 6 - 1 = 5
//   runes[5] = '!'
//
// Examples:
//   LastRune("Hello!") returns '!'
//   LastRune("Salut!") returns '!'
//   LastRune("Ola!") returns '!'
//   LastRune("Go") returns 'o'

// Status: Required

func LastRune(s string) rune {
	// Convert string to rune slice
	runes := []rune(s)

	// Return the last rune (length - 1)
	return runes[len(runes)-1]
}
