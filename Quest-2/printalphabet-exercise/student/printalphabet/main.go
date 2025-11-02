package main

import "github.com/01-edu/z01"

// TODO: Write your code here
// Print the lowercase alphabet (a-z) on one line
//
// Hints:
// - Use a for loop to iterate from 'a' to 'z'
// - In Go, 'a' and 'z' are rune literals (characters)
// - Loop structure: for i := 'a'; i <= 'z'; i++ { ... }
// - Print each character with z01.PrintRune(i)
// - Don't forget to print a newline at the end: z01.PrintRune('\n')
//
// Expected output:
// abcdefghijklmnopqrstuvwxyz

func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
