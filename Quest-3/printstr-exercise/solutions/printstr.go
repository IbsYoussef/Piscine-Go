package solutions

import "github.com/01-edu/z01"

// Function Breakdown:
//
// PrintStr(s string):
//   - Takes a string as parameter
//   - Prints each character one by one
//   - Uses z01.PrintRune to output each character
//
// Approach:
//   - Use for range loop to iterate over string
//   - Each iteration gives us a rune (character)
//   - Print each rune using z01.PrintRune
//
// Key Concepts:
//   - String iteration: for _, r := range s
//   - Runes: Go's representation of characters
//   - z01.PrintRune: Custom print function from 01-edu
//   - No return value (void function)
//
// Example:
//   PrintStr("Hello")
//   // Output: Hello

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
