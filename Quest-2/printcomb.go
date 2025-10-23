package piscine

import "github.com/01-edu/z01"

// Exercise: printcomb
//
// Required:
// Print all unique combinations of three different digits (a<b<c),
// separated by ", ", in ascending order. No trailing comma on the last combo.
//
// Expected signature:
// func PrintComb() {}
//
// Explanation:
// - Generate ordered triples with three nested loops ensuring a<b<c.
// - Print each triple; add ", " unless it's the final one (789).
// - End with a newline.
//
// Code Breakdown (Step-by-Step):
// 1) Outer loops: a in '0'..'7', b in a+1..'8', c in b+1..'9' → guarantees distinct & ascending.
// 2) Print a, b, c via `z01.PrintRune`.
// 3) Check if it's not the last combo (a,b,c != 7,8,9) to append ", ".
// 4) After all, print '\n'.
//
// Example output (truncated):
// 012, 013, ..., 789
//
// Solution:
func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a != '7' || b != '8' || c != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
