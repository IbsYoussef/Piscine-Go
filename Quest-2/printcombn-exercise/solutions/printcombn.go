package solutions

import "fmt"

// Function Breakdown:
//
// PrintCombN(n int):
//   - Prints all combinations of n different digits in ascending order
//   - Uses recursion to generate combinations
//   - Valid for 0 < n < 10
//
// Approach:
//   - Validate n bounds (return if invalid)
//   - Generate max combination to know when to stop commas
//   - Recursive helper builds combinations digit by digit
//   - Base case: print when length reaches n
//   - Recursive case: try each digit from start to 9
//
// Key Concepts:
//   - Recursion for combination generation
//   - String building with concatenation
//   - Edge detection for formatting
//   - Dynamic algorithm based on n
//
// Output examples:
//   n=1: 0, 1, 2, ..., 9
//   n=3: 012, 013, ..., 789

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	// Generate the maximum combination (e.g., "789" for n=3)
	generateMaxCombination := func(n int) string {
		max := ""
		for i := 0; i < n; i++ {
			max += string(rune('9' - rune(n-1-i)))
		}
		return max
	}

	// Recursive helper to generate combinations
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
