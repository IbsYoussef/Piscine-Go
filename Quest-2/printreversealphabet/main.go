package main

import "github.com/01-edu/z01"

// Exercise: printreversealphabet
//
// Required:
// Print the lowercase Latin alphabet in reverse order ('z' → 'a') on a single line,
// followed by a newline. Casting is not allowed.
//
// Explanation:
// - Start from 'z' and step down to 'a'.
// - Use `z01.PrintRune` to output each rune.
// - Finish with a newline.
//
// Code Breakdown (Step-by-Step):
// 1) `for i := 'z'; i >= 'a'; i--` — countdown loop over runes.
// 2) `z01.PrintRune(i)` — print current letter.
// 3) No casting needed; runes represent characters directly.
// 4) Print '\n' at the end for correct formatting.
//
// Example output:
// $ go run .
// zyxwvutsrqponmlkjihgfedcba
// $
//
// Solution:
func main() {
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
