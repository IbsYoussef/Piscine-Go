package solutions

import "github.com/01-edu/z01"

// Function Breakdown:
//
// PrintComb():
//   - Prints all combinations of 3 different digits in ascending order
//   - Format: abc, def, ..., 789 (no trailing comma)
//   - Uses three nested loops to generate combinations
//
// Approach:
//   - Loop a: '0' to '7' (leaves room for b and c)
//   - Loop b: a+1 to '8' (ensures b > a, leaves room for c)
//   - Loop c: b+1 to '9' (ensures c > b)
//   - This guarantees a < b < c
//
// Key Concepts:
//   - Nested loops for combinations
//   - Rune arithmetic ('0' + 1 = '1')
//   - Conditional formatting (comma separator)
//   - Edge case: last combination has no comma
//
// Output:
//   012, 013, 014, ..., 789

func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				// Don't print comma after last combination (789)
				if a != '7' || b != '8' || c != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
