package main

import "github.com/01-edu/z01"

// Exercise: printalphabet
//
// Required:
// Write a program that prints the lowercase Latin alphabet (a–z)
// on a single line, followed by a newline character.
//
// Explanation:
// - The Latin alphabet consists of the characters 'a' through 'z'.
// - Each letter is printed sequentially using a for loop.
// - `z01.PrintRune()` is used to print runes (Unicode characters) one by one.
// - Finally, `z01.PrintRune('\n')` prints a newline to end the line properly.
//
// Code Breakdown (Step-by-Step):
// 1) `for i := 'a'; i <= 'z'; i++` — loop from rune 'a' to 'z' inclusive.
// 2) `z01.PrintRune(i)` — print the current rune (a letter).
// 3) Runeliterals like 'a'/'z' are characters; Go can iterate over them directly.
// 4) After the loop, print '\n' to terminate the line.
//
// Example output:
// $ go run .
// abcdefghijklmnopqrstuvwxyz
// $
//
// Solution:
func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
