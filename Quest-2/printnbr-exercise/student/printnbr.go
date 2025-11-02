package student

import "github.com/01-edu/z01"

// TODO: Write your code here
// Implement the PrintNbr function
//
// Hints:
// - Function signature: func PrintNbr(n int)
// - Handle zero: if n == 0, print '0' and return
// - Handle negative: if n < 0, print '-'
// - Extract digits using modulo: digit := n % 10
// - Store digits in an array (max 19 for int range)
// - Increment n: n /= 10
// - Handle negative digits: if digit < 0, make it positive
// - After extracting all digits, print in reverse order
// - Convert to rune: rune('0' + digit)
// - Loop backwards through array to print
//
// Examples:
//   PrintNbr(-123) prints: -123
//   PrintNbr(0) prints: 0
//   PrintNbr(456) prints: 456

func PrintNbr(n int) {
	// Your code here
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	var digits [19]int // Max digits for int
	var i int
	for n != 0 {
		d := n % 10
		if d < 0 {
			d = -d // Make positive for printing
		}
		digits[i] = d
		i++
		n /= 10
	}
	// Print in reverse order
	for j := i - 1; j >= 0; j-- {
		z01.PrintRune(rune('0' + digits[j]))
	}
}
