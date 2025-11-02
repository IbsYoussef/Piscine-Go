package main

import "github.com/01-edu/z01"

// Function Breakdown:
//
// main():
//   - Uses a for loop to iterate through runes 'z' to 'a' (reverse)
//   - Prints each character individually using z01.PrintRune()
//   - Adds a newline at the end
//
// Approach:
//   - Start from 'z' and decrement until 'a'
//   - Use i-- to count backwards
//   - Condition: i >= 'a' (greater than or equal)
//
// Key Concepts:
//   - Runes: Unicode code points (characters)
//   - For loop with reverse iteration
//   - Decrement operator (i--)
//   - z01.PrintRune(): Prints a single rune to stdout
//
// Output:
//   zyxwvutsrqponmlkjihgfedcba

func main() {
	// Loop from 'z' to 'a' (inclusive, reverse order)
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i) // Print current letter
	}
	z01.PrintRune('\n') // Print newline
}
