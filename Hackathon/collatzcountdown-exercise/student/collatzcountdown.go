package student

// TODO: Implement CollatzCountdown function
// Count steps to reach 1 using Collatz rules
//
// Collatz Rules:
// - If even: divide by 2
// - If odd: multiply by 3 and add 1
// - Repeat until reaching 1
//
// Algorithm:
// 1. Handle edge case: if start <= 0, return -1
// 2. Initialize steps = 0
// 3. Loop while start != 1:
//    - Check if even: start % 2 == 0
//    - If even: start = start / 2
//    - If odd: start = start * 3 + 1
//    - Increment steps
// 4. Return steps
//
// Example:
//   Start: 12
//   12 → 6 → 3 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//   Steps: 9

func CollatzCountdown(start int) int {
	// Your code here
	return 0
}
