package piscine

import "fmt"

// Exercise: printcombn
//
// Bonus:
// Print all combinations of n different digits (0 < n < 10) in ascending order.
// Combinations are separated by ", " and the last one has no trailing comma.
// End with a newline.
//
// Expected signature:
// func PrintCombN(n int) {}
//
// Explanation:
// - Build ascending sequences of length n from digits 0..9 (no repeats).
// - Use recursion to append the next digit after the current start.
// - Detect the final combination to omit the trailing comma.
//
// Code Breakdown (Step-by-Step):
// 1) Validate bounds: return if n <= 0 or n >= 10.
// 2) `generateMaxCombination(n)` builds the last combo, e.g. for n=3 → "789".
// 3) Recursive `helper(current, start, n)`:
//   - If `len(current)==n`: print it; if not the max combo, print ", ".
//   - Else loop i from `start` to 9 and recurse with `current + digit(i)` and `i+1`.
//
// 4) Kick off with `helper("", 0, n)`; finish with `fmt.Println()`.
//
// Example output (truncated):
// n=1 → 0, 1, ..., 9
// n=3 → 012, 013, ..., 789
// n=9 → 012345678, ..., 123456789
//
// Solution:
func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	generateMaxCombination := func(n int) string {
		max := ""
		for i := 0; i < n; i++ {
			max += string(rune('9' - rune(n-1-i)))
		}
		return max
	}
	var helper func(current string, start int, n int)
	helper = func(current string, start int, n int) {
		if len(current) == n {
			fmt.Print(current)
			if current != generateMaxCombination(n) {
				fmt.Print(", ")
			}
			return
		}
		for i := start; i < 10; i++ {
			helper(current+string(rune('0'+i)), i+1, n)
		}
	}
	helper("", 0, n)
	fmt.Println()
}
