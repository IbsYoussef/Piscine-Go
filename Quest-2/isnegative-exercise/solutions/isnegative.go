package solutions

import "github.com/01-edu/z01"

// Function Breakdown:
//
// IsNegative(nb int):
//   - Checks if the given integer is negative
//   - Prints 'T' if negative, 'F' otherwise
//   - Always prints a newline after
//
// Approach:
//   - Simple conditional: if nb < 0
//   - Use z01.PrintRune() for output
//   - No return value needed
//
// Key Concepts:
//   - Conditional logic (if/else)
//   - Function with parameter, no return
//   - Character output with z01.PrintRune()
//
// Examples:
//   IsNegative(-1)  → T
//   IsNegative(0)   → F
//   IsNegative(5)   → F

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T') // Negative: print T
	} else {
		z01.PrintRune('F') // Zero or positive: print F
	}
	z01.PrintRune('\n') // Always print newline
}
