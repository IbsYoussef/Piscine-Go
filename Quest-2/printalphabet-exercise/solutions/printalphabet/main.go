package main

import "github.com/01-edu/z01"

// Function Breakdown:
//
// main():
//   - Uses a for loop to iterate through runes 'a' to 'z'
//   - Prints each character individually using z01.PrintRune()
//   - Adds a newline at the end
//
// Approach:
//   - In Go, rune literals ('a', 'z') are character values
//   - Can be used directly in loop conditions
//   - Incrementing a rune moves to the next character
//
// Key Concepts:
//   - Runes: Unicode code points (characters)
//   - For loop with character range
//   - z01.PrintRune(): Prints a single rune to stdout
//
// Output:
//   abcdefghijklmnopqrstuvwxyz

func main() {
	// Loop from 'a' to 'z' (inclusive)
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i) // Print current letter
	}
	z01.PrintRune('\n') // Print newline
}
