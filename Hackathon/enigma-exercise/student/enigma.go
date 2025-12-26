package student

// TODO: Implement Enigma function
// Swap values through multiple levels of pointers
//
// Swaps to perform:
// - a → c (a's value goes into c)
// - c → d (c's value goes into d)
// - d → b (d's value goes into b)
// - b → a (b's value goes into a)
//
// Pointer levels:
// - a ***int → dereference 3 times to get value: ***a
// - b *int → dereference 1 time to get value: *b
// - c *******int → dereference 7 times to get value: *******c
// - d ****int → dereference 4 times to get value: ****d
//
// Algorithm:
// 1. Save all original values in temp variables
//    tempA := ***a
//    tempB := *b
//    tempC := *******c
//    tempD := ****d
// 2. Assign new values (following the swap pattern)
//    ***a = tempB (a gets b)
//    *b = tempD (b gets d)
//    *******c = tempA (c gets a)
//    ****d = tempC (d gets c)

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Your code here
}
