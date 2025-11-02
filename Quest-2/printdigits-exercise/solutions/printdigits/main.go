package main

import "github.com/01-edu/z01"

// Function Breakdown:
//
// main():
//   - Uses a for loop to iterate through runes '0' to '9'
//   - Prints each digit individually using z01.PrintRune()
//   - Adds a newline at the end
//
// Approach:
//   - In Go, rune literals ('0', '9') represent digit characters
//   - Can be used directly in loop conditions
//   - Incrementing a rune moves to the next character
//
// Key Concepts:
//   - Runes: Unicode code points (characters)
//   - For loop with character range
//   - z01.PrintRune(): Prints a single rune to stdout
//
// Output:
//   0123456789

func main() {
	// Loop from '0' to '9' (inclusive)
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i) // Print current digit
	}
	z01.PrintRune('\n') // Print newline
}
