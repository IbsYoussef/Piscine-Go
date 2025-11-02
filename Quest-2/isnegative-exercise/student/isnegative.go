package student

import "github.com/01-edu/z01"

// TODO: Write your code here
// Implement the IsNegative function
//
// Hints:
// - Function signature: func IsNegative(nb int)
// - Check if nb is less than 0
// - If negative: print 'T' using z01.PrintRune('T')
// - Otherwise: print 'F' using z01.PrintRune('F')
// - Always print a newline at the end: z01.PrintRune('\n')
// - No return value needed
//
// Examples:
//   IsNegative(-1)  prints: T
//   IsNegative(0)   prints: F
//   IsNegative(5)   prints: F

func IsNegative(nb int) {
	// Your code here
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
