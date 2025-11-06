package solutions

// Function Breakdown:
//
// RecursivePower(nb int, power int) int:
//   - Takes base (nb) and exponent (power)
//   - Calculates nb^power using recursion
//   - Returns 0 for negative powers
//   - Returns 1 for power = 0
//
// Approach:
//   - Base case 1: power < 0 → return 0
//   - Base case 2: power == 0 → return 1
//   - Base case 3: power == 1 → return nb
//   - Recursive case: nb × RecursivePower(nb, power-1)
//   - Each call reduces power by 1
//
// Key Concepts:
//   - Recursion: Function calls itself
//   - Base cases: Stop conditions
//   - Power reduction: Decreasing exponent
//   - Exponentiation: nb^power = nb × nb^(power-1)
//
// Recursion flow for RecursivePower(4, 3):
//   RecursivePower(4, 3)
//   → 4 × RecursivePower(4, 2)
//     → 4 × RecursivePower(4, 1)
//       → 4 (base case)
//     → 4 × 4 = 16
//   → 4 × 16 = 64
//
// Examples:
//   RecursivePower(4, 3) = 64
//   RecursivePower(2, 5) = 32
//   RecursivePower(5, 0) = 1
//   RecursivePower(3, -2) = 0

// Status: Required

func RecursivePower(nb int, power int) int {
	// Base case: negative power returns 0
	if power < 0 {
		return 0
	}

	// Base case: any number to power 0 is 1
	if power == 0 {
		return 1
	}

	// Base case: any number to power 1 is itself
	if power == 1 {
		return nb
	}

	// Recursive case: nb^power = nb × nb^(power-1)
	return nb * RecursivePower(nb, power-1)
}
