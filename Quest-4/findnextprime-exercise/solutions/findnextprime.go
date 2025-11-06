package solutions

// Function Breakdown:
//
// IsPrime(nb int) bool:
//   - Helper function to check if number is prime
//   - Returns true if prime, false otherwise
//   - Optimized to check only up to √nb
//   - This is the same function from the isprime exercise
//   - Included here for convenience/self-contained exercise
//
// FindNextPrime(nb int) int:
//   - Takes an integer
//   - Returns first prime >= nb
//   - Uses IsPrime helper function to check primality
//   - Handles edge cases
//
// Approach (FindNextPrime):
//   - Handle edge case: nb ≤ 1 → return 2
//   - Check if nb is already prime using IsPrime
//   - If yes, return nb
//   - If no, increment and check again
//   - Continue until prime found
//
// Approach (IsPrime):
//   - Handle edge case: nb ≤ 1 → return false
//   - Calculate square root of nb
//   - Check divisibility from 2 to sqrt-1
//   - If divisible, return false
//   - If no divisors found, return true
//
// Key Concepts:
//   - Code reuse: IsPrime is copied here for self-contained exercise
//   - Helper functions: Breaking down complex problems
//   - Infinite loop with break condition
//   - Sequential search
//   - Edge case handling
//   - Optimization: Only check up to √nb
//
// Examples:
//   IsPrime(5) = true
//   IsPrime(4) = false
//   FindNextPrime(5) = 5 (already prime)
//   FindNextPrime(4) = 5 (next prime after 4)
//   FindNextPrime(14) = 17 (skip 14, 15, 16)
//   FindNextPrime(1) = 2 (first prime)

// Status: Optional

// IsPrime is a helper function included here
// This is the same function from the isprime exercise
func IsPrime(nb int) bool {
	// Numbers less than or equal to 1 are not prime
	if nb <= 1 {
		return false
	}

	// Calculate square root of nb (optimization)
	sqrt := 1
	for sqrt*sqrt <= nb {
		sqrt++
	}

	// Check if divisible by any number from 2 to sqrt-1
	for i := 2; i < sqrt; i++ {
		if nb%i == 0 {
			return false // Found a divisor, not prime
		}
	}

	// No divisors found, number is prime
	return true
}

func FindNextPrime(nb int) int {
	// Edge case: numbers <= 1, return first prime
	if nb <= 1 {
		return 2
	}

	// Check if nb is already prime
	if IsPrime(nb) {
		return nb
	}

	// Keep incrementing until we find a prime
	for {
		nb++
		if IsPrime(nb) {
			return nb
		}
	}
}
