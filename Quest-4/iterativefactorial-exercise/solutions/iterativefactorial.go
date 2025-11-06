package solutions

// Function Breakdown:
//
// IterativeFactorial(nb int) int:
//   - Takes an integer as parameter
//   - Calculates factorial using iteration
//   - Returns 0 for errors (negative, overflow)
//   - Returns factorial value for valid input
//
// Approach:
//   - Handle edge case: negative numbers return 0
//   - Initialize result to 1
//   - Loop from 1 to nb, multiplying each value
//   - Check for overflow before each multiplication
//   - Return the accumulated result
//
// Key Concepts:
//   - Factorial: n! = n × (n-1) × ... × 2 × 1
//   - Iterative: Using loops, not recursion
//   - Overflow detection: Check before multiplying
//   - Edge cases: 0! = 1, 1! = 1, negative → 0
//
// Overflow check:
//   Before: result *= i
//   Check: result > (1<<63-1)/i
//   If true, multiplication would overflow
//
// Examples:
//   IterativeFactorial(4) = 24
//   IterativeFactorial(0) = 1
//   IterativeFactorial(-5) = 0
//   IterativeFactorial(21) = 0 (overflow)

// IterativeFactorial is a function that calculates the factorial of a number using an iterative approach
// Status: Required

func IterativeFactorial(nb int) int {
	// Handle negative inputs
	if nb < 0 {
		return 0 // Return 0 for negative inputs
	}

	// Initialize result (0! and 1! both equal 1)
	result := 1

	// Multiply all numbers from 1 to nb
	for i := 1; i <= nb; i++ {
		// Check for overflow before multiplying
		if result > (1<<63-1)/i {
			return 0 // Return 0 for overflow
		}
		result *= i
	}

	return result
}
