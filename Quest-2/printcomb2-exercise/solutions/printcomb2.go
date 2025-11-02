package solutions

import "github.com/01-edu/z01"

// Function Breakdown:
//
// PrintComb2():
//   - Prints all combinations of two different 2-digit numbers
//   - Format: 00 01, 00 02, ..., 98 99
//   - First number must be strictly less than second
//
// Approach:
//   - Four nested loops: a, b, c, d
//   - First number: 10*a + b (00-99)
//   - Second number: 10*c + d (00-99)
//   - Ensure first < second by careful loop bounds
//
// Key Concepts:
//   - Four nested loops
//   - Digit to rune conversion: rune('0' + digit)
//   - Complex loop initialization for ordering
//   - Conditional comma printing
//
// Output:
//   00 01, 00 02, ..., 98 99

func PrintComb2() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := a; c <= 9; c++ {
				dStart := b + 1
				if c > a {
					dStart = 0 // reset when tens digit increases
				}
				for d := dStart; d <= 9; d++ {
					z01.PrintRune(rune('0' + a))
					z01.PrintRune(rune('0' + b))
					z01.PrintRune(' ')
					z01.PrintRune(rune('0' + c))
					z01.PrintRune(rune('0' + d))
					// Don't print comma after last combination (98 99)
					if !(a == 9 && b == 8 && c == 9 && d == 9) {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
