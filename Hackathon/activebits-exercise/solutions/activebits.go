package solutions

// ActiveBits counts the number of 1-bits in the binary representation of n
//
// Algorithm:
// 1. Check the last bit using n & 1
//    - If last bit is 1, add to count
//    - If last bit is 0, add nothing
// 2. Shift n right by 1 bit (n >>= 1)
//    - Discards the last bit
//    - Makes the next bit the "last bit"
// 3. Repeat until n becomes 0
//
// Bitwise operations:
// - n & 1: Checks if last bit is 1 or 0
// - n >>= 1: Shifts all bits right by 1 position
//
// Example: n = 7 (binary: 0111)
//   Loop 1: 0111 & 1 = 1, count = 1, shift → 0011
//   Loop 2: 0011 & 1 = 1, count = 2, shift → 0001
//   Loop 3: 0001 & 1 = 1, count = 3, shift → 0000
//   Result: 3 active bits

func ActiveBits(n int) int {
	count := 0

	// Process each bit until n becomes 0
	for n != 0 {
		// Check if last bit is 1
		count += n & 1

		// Shift right to process next bit
		n >>= 1
	}

	return count
}
