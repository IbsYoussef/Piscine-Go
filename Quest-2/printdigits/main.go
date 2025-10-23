package main

import "github.com/01-edu/z01"

// Exercise: printdigits
//
// Required:
// Print the decimal digits (0 → 9) in ascending order on one line,
// followed by a newline.
//
// Explanation:
// - Characters '0'..'9' are consecutive runes.
// - Iterate from '0' to '9' and print each as a rune.
// - Print a trailing newline.
//
// Code Breakdown (Step-by-Step):
// 1) `for i := '0'; i <= '9'; i++` — iterate through digit runes.
// 2) `z01.PrintRune(i)` — output the current digit.
// 3) Digits are characters; arithmetic on runes is valid.
// 4) Print '\n' at the end.
//
// Example output:
// $ go run .
// 0123456789
// $
//
// Solution:
func main() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
