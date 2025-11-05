package piscine

// IterativeFactorial is a function that calculates the factorial of a number using an iterative approach
// Status: Required

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0 // Return 0 for negative inputs
	}
	result := 1
	for i := 1; i <= nb; i++ {
		if result > (1<<63-1)/i { // Check for overflow
			return 0 // Return 0 for overflow
		}
		result *= i
	}
	return result
}
