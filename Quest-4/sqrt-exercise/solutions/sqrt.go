package solutions

// Function Breakdown:
//
// Sqrt(nb int) int:
//   - Takes an integer
//   - Returns square root if it's a whole number
//   - Returns 0 if not a perfect square
//   - Returns 0 for negative numbers
//
// Approach:
//   - Handle negative numbers: return 0
//   - Handle zero: return 0
//   - Loop from 1 to nb
//   - Check if i × i equals nb
//   - If match found, return i
//   - If no match, return 0
//
// Key Concepts:
//   - Square root: √n × √n = n
//   - Perfect squares: 1, 4, 9, 16, 25, 36...
//   - Linear search: Testing each value
//   - Early return: Stop when found
//
// Examples:
//   Sqrt(4) = 2 (2 × 2 = 4)
//   Sqrt(9) = 3 (3 × 3 = 9)
//   Sqrt(3) = 0 (not a perfect square)
//   Sqrt(-4) = 0 (negative)
//   Sqrt(0) = 0

// Status: Optional

func Sqrt(nb int) int {
	// Negative numbers have no real square root
	if nb < 0 {
		return 0
	}

	// Square root of 0 is 0
	if nb == 0 {
		return 0
	}

	// Try each number from 1 to nb
	for i := 1; i <= nb; i++ {
		// Check if i squared equals nb
		if i*i == nb {
			return i
		}
	}

	// No perfect square found
	return 0
}
