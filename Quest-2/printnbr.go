package piscine

import "github.com/01-edu/z01"

// Exercise: printnbr
//
// Optional:
// Print an integer `n` using only z01.PrintRune. Handle all int values
// (including negative and zero), without converting to int64.
//
// Expected signature:
// func PrintNbr(n int) {}
//
// Explanation:
// - Print '0' immediately for n == 0.
// - If negative, print '-' then operate on its absolute value digit-by-digit.
// - Extract digits with modulo (%) and division (/), store, then print in reverse.
//
// Code Breakdown (Step-by-Step):
// 1) If n == 0 → print '0' and return.
// 2) If n < 0 → print '-' and keep extracting digits; apply `-digit` when needed.
// 3) Loop: `digit := n % 10`; store digits in an array/slice; `n /= 10`.
// 4) Print digits in reverse order using `'0' + digit` to form runes.
// 5) Works for full int range without extra types.
//
// Example output:
// -1230123 (for -123, 0, 123 printed consecutively)
//
// Solution:
func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	var digits [19]int
	var i int
	for n != 0 {
		d := n % 10
		if d < 0 {
			d = -d
		}
		digits[i] = d
		i++
		n /= 10
	}
	for j := i - 1; j >= 0; j-- {
		z01.PrintRune(rune('0' + digits[j]))
	}
}
