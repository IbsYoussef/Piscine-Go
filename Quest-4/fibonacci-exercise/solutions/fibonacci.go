package solutions

// Function Breakdown:
//
// Fibonacci(index int) int:
//   - Takes an index position
//   - Returns fibonacci value at that position
//   - Returns -1 for negative index
//   - Uses recursion to calculate
//
// Approach:
//   - Base case 1: index < 0 → return -1
//   - Base case 2: index == 0 → return 0
//   - Base case 3: index == 1 → return 1
//   - Recursive case: Fibonacci(index-1) + Fibonacci(index-2)
//   - Each call branches into two recursive calls
//
// Key Concepts:
//   - Fibonacci: F(n) = F(n-1) + F(n-2)
//   - Sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21...
//   - Tree recursion: Multiple branches
//   - Base cases prevent infinite recursion
//
// Fibonacci sequence:
//   Index: 0  1  2  3  4  5  6  7  8
//   Value: 0  1  1  2  3  5  8  13 21
//
// Examples:
//   Fibonacci(4) = 3
//   Fibonacci(0) = 0
//   Fibonacci(1) = 1
//   Fibonacci(-1) = -1

// Status: Required

func Fibonacci(index int) int {
	// Base case: negative index returns -1
	if index < 0 {
		return -1
	}

	// Base case: fibonacci(0) = 0
	if index == 0 {
		return 0
	}

	// Base case: fibonacci(1) = 1
	if index == 1 {
		return 1
	}

	// Recursive case: F(n) = F(n-1) + F(n-2)
	return Fibonacci(index-1) + Fibonacci(index-2)
}
