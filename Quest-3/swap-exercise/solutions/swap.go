package solutions

// Function Breakdown:
//
// Swap(a *int, b *int):
//   - Takes two pointers to int as parameters
//   - Swaps the values at those pointers
//   - No return value
//
// Approach:
//   - Use Go's multiple assignment feature
//   - Dereference both pointers to access values
//   - Assign in one line: *a, *b = *b, *a
//
// Key Concepts:
//   - Pointers: Store memory addresses
//   - Dereference operator (*): Access value at pointer
//   - Multiple assignment: Swap without temp variable
//
// Example:
//   a := 0
//   b := 1
//   Swap(&a, &b)  // a is now 1, b is now 0

func Swap(a *int, b *int) {
	*a, *b = *b, *a // Swap values using multiple assignment
}
