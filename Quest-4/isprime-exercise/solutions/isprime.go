package solutions

// Function Breakdown:
//
// IsPrime(nb int) bool:
//   - Takes an integer
//   - Returns true if prime, false otherwise
//   - Optimized to check only up to √nb
//   - 1 and numbers ≤ 1 are not prime
//
// Approach:
//   - Handle edge case: nb ≤ 1 → return false
//   - Calculate square root of nb
//   - Check divisibility from 2 to sqrt-1
//   - If divisible, return false
//   - If no divisors found, return true
//
// Optimization:
//   - Only check up to √nb
//   - If nb has divisor > √nb, it must also have divisor < √nb
//   - This reduces time complexity significantly
//
// Key Concepts:
//   - Prime: Only divisible by 1 and itself
//   - Composite: Has divisors other than 1 and itself
//   - Square root optimization: Major performance improvement
//   - Early return: Stop as soon as divisor found
//
// Examples:
//   IsPrime(5) = true (no divisors)
//   IsPrime(4) = false (divisible by 2)
//   IsPrime(1) = false (by definition)
//   IsPrime(17) = true (prime)

// Status: Optional

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
