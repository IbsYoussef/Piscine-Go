package solutions

// Function Breakdown:
//
// RecursiveFactorial(nb int) int:
//   - Takes an integer as parameter
//   - Calculates factorial using recursion
//   - Returns 0 for errors (negative, overflow)
//   - Returns factorial value for valid input
//
// Approach:
//   - Base case 1: negative or overflow (>20) → return 0
//   - Base case 2: 0 or 1 → return 1
//   - Recursive case: nb × RecursiveFactorial(nb-1)
//   - Each call reduces problem size by 1
//
// Key Concepts:
//   - Recursion: Function calls itself
//   - Base cases: Stop conditions
//   - Recursive case: Smaller subproblem
//   - Call stack: Each call waits for the next
//
// Recursion flow for RecursiveFactorial(4):
//   RecursiveFactorial(4)
//   → 4 × RecursiveFactorial(3)
//     → 3 × RecursiveFactorial(2)
//       → 2 × RecursiveFactorial(1)
//         → 1 (base case)
//       → 2 × 1 = 2
//     → 3 × 2 = 6
//   → 4 × 6 = 24
//
// Examples:
//   RecursiveFactorial(4) = 24
//   RecursiveFactorial(0) = 1
//   RecursiveFactorial(-5) = 0
//   RecursiveFactorial(21) = 0 (overflow)

// Recursive Factorial is a function that find the factorial of a given a value using recursion
// Status: Required

func RecursiveFactorial(nb int) int {
	// Base case: negative numbers or overflow (> 20)
	if nb < 0 || nb > 20 {
		return 0
	}

	// Base case: 0! = 1 and 1! = 1
	if nb == 0 || nb == 1 {
		return 1
	}

	// Recursive case: n! = n × (n-1)!
	return nb * RecursiveFactorial(nb-1)
}
