package solutions

// CollatzCountdown returns the number of steps to reach 1
// following the Collatz conjecture rules
//
// Rules:
// - If number is even: divide by 2
// - If number is odd: multiply by 3 and add 1
// - Continue until reaching 1
//
// # Returns -1 if start is 0 or negative
//
// Algorithm:
// 1. Check edge case (start <= 0)
// 2. Initialize step counter
// 3. Loop while not at 1:
//   - Check if even or odd
//   - Apply transformation
//   - Increment counter
//
// 4. Return step count
//
// Example: 12 → 6 → 3 → 10 → 5 → 16 → 8 → 4 → 2 → 1 (9 steps)
func CollatzCountdown(start int) int {
	// Handle edge case: zero or negative
	if start <= 0 {
		return -1
	}

	// Initialize step counter
	steps := 0

	// Loop until we reach 1
	for start != 1 {
		if start%2 == 0 {
			// Even: divide by 2
			start /= 2
		} else {
			// Odd: multiply by 3 and add 1
			start = start*3 + 1
		}
		// Increment step counter
		steps++
	}

	return steps
}
