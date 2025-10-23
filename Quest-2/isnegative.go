package piscine

import "github.com/01-edu/z01"

// Exercise: isnegative
//
// Required:
// Write a function that prints 'T' if the given int is negative,
// otherwise print 'F'. Each result is followed by a newline.
//
// Expected signature:
// func IsNegative(nb int) {}
//
// Explanation:
// - If nb < 0 → print 'T', else 'F'.
// - Always print a newline at the end.
// - Use `z01.PrintRune` for character output.
//
// Code Breakdown (Step-by-Step):
// 1) `if nb < 0` — branch on negativity.
// 2) Print the corresponding rune: 'T' or 'F'.
// 3) Print '\n' after the character.
// 4) No return value; output is the requirement.
//
// Example usage program:
// package main
// import "piscine"
//
//	func main() {
//	  piscine.IsNegative(1)
//	  piscine.IsNegative(0)
//	  piscine.IsNegative(-1)
//	}
//
// Example output:
// F
// F
// T
//
// Solution:
func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
