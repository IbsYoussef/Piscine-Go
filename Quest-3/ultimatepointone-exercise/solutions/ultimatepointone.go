package solutions

// Function Breakdown:
//
// UltimatePointOne(n ***int):
//   - Takes a triple pointer to int
//   - Sets the final int value to 1
//   - Uses triple dereferencing
//
// Approach:
//   - n is a pointer to a pointer to a pointer to int
//   - ***n dereferences three times to reach the int
//   - Assign 1 to that int
//
// Key Concepts:
//   - Triple pointer: ***int
//   - Dereferencing: * follows a pointer
//   - Multiple levels: each * goes one level deeper
//   - Direct assignment: ***n = 1
//
// Example:
//   a := 0
//   b := &a    // b points to a
//   c := &b    // c points to b
//   UltimatePointOne(&c)
//   // a is now 1
//
// Visual:
//   &c → c → b → a
//   (***int) (**int) (*int) (int)

func UltimatePointOne(n ***int) {
	***n = 1
}
