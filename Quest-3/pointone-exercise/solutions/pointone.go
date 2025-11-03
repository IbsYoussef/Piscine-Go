package solutions

// Function Breakdown:
//
// PointOne(n *int):
//   - Takes a pointer to an int as parameter
//   - Sets the value at that pointer to 1
//   - No return value
//
// Approach:
//   - Use dereference operator (*) to access the value
//   - Assign 1 to the dereferenced pointer
//
// Key Concepts:
//   - Pointers: Store memory addresses
//   - Dereference operator (*): Access value at pointer
//   - Address operator (&): Get address of variable
//
// Example:
//   n := 0
//   PointOne(&n)  // n is now 1

func PointOne(n *int) {
	*n = 1 // Dereference pointer and assign 1
}
