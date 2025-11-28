package solutions

import "github.com/01-edu/z01"

// Function Breakdown:
//
// PrintWordsTables(a []string):
//   - Prints each string in the slice on a separate line
//   - Uses z01.PrintRune for character-by-character printing
//
// Approach:
//   - Loop through each string in the slice
//   - For each string, loop through each character
//   - Print each character
//   - Print newline after each string
//
// Key Concepts:
//   - Nested loops: Outer for strings, inner for characters
//   - z01.PrintRune: Print individual characters
//   - Newline: '\n' after each word
//
// Examples:
//   PrintWordsTables(["Hello", "world"])
//   Prints:
//   Hello
//   world

// Status: Optional

func PrintWordsTables(a []string) {
	// Loop through each word in the slice
	for _, word := range a {
		// Print each character of the word
		for _, char := range word {
			z01.PrintRune(char)
		}
		// Print newline after each word
		z01.PrintRune('\n')
	}
}
