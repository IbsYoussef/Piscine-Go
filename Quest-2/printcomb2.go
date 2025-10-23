package piscine

import "github.com/01-edu/z01"

// Exercise: printcomb2
//
// Optional:
// Print all combinations of two different two-digit numbers in ascending order.
// Format: "00 01, 00 02, ..., 98 99" (no trailing comma), then newline.
//
// Expected signature:
// func PrintComb2() {}
//
// Explanation:
// - Numbers range 00..99; first must be strictly less than the second.
// - Print each as two digits, separated by a space; combos separated by ", ".
// - No comma after the last pair (98 99).
//
// Code Breakdown (Step-by-Step):
// 1) Four nested loops represent digits: a b (first number), c d (second).
// 2) `c` starts at `a`; `d` starts at `b+1` (or 0 when c > a) to enforce ordering.
// 3) Print "ab cd" using rune math: '0'+digit.
// 4) Append ", " unless current pair is the last (98 99).
// 5) Print '\n' at the end.
//
// Example output (truncated):
// 00 01, 00 02, ..., 98 99
//
// Solution:
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
