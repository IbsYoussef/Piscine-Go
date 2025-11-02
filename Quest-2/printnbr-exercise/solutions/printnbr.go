package solutions

import "github.com/01-edu/z01"

// Function Breakdown:
//
// PrintNbr(n int):
//   - Prints an integer using only z01.PrintRune()
//   - Handles negative, zero, and positive values
//   - No conversion to int64
//
// Approach:
//   - Special case: n == 0, print '0' and return
//   - Negative: print '-', then extract digits
//   - Extract digits using n % 10 and n / 10
//   - Store in array (need to print in reverse)
//   - Handle negative digits (make positive)
//
// Key Concepts:
//   - Digit extraction: modulo and division
//   - Array for digit storage
//   - Reverse printing
//   - Handling negative numbers
//   - Rune conversion: '0' + digit
//
// Output:
//   PrintNbr(-123) → -123
//   PrintNbr(0) → 0
//   PrintNbr(456) → 456

func PrintNbr(n int) {
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
