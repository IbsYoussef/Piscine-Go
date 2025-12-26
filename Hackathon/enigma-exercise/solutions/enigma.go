package solutions

// Enigma swaps values through multiple levels of pointer indirection
//
// Swaps:
// - a's value → c
// - c's value → d
// - d's value → b
// - b's value → a
//
// Pointer levels:
// - a ***int: 3 levels of indirection
// - b *int: 1 level of indirection
// - c *******int: 7 levels of indirection
// - d ****int: 4 levels of indirection
//
// Algorithm:
// 1. Save all original values using full dereferencing
//    - ***a gets the int value from a
//    - *b gets the int value from b
//    - *******c gets the int value from c
//    - ****d gets the int value from d
// 2. Assign new values to each pointer
//    - a gets b's value
//    - b gets d's value
//    - c gets a's value
//    - d gets c's value
//
// Key: The number of * in the type tells you how many times to
// dereference to reach the actual int value

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Save all original values
	tempA := ***a     // a's value (5)
	tempB := *b       // b's value (2)
	tempC := *******c // c's value (7)
	tempD := ****d    // d's value (6)

	// Perform the swaps:
	// a gets b's value
	***a = tempB

	// b gets d's value
	*b = tempD

	// c gets a's value
	*******c = tempA

	// d gets c's value
	****d = tempC
}
