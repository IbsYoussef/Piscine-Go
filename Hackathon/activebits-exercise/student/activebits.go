package student

// TODO: Implement ActiveBits function
// Count the number of 1-bits in binary representation
//
// Algorithm:
// 1. Initialize count = 0
// 2. Loop while n != 0:
//    - Check last bit: n & 1
//    - Add result to count
//    - Shift right: n >>= 1
// 3. Return count
//
// Bitwise operations explained:
// - n & 1: ANDs n with 1 to check last bit
//   Example: 0111 & 0001 = 0001 (last bit is 1)
//            0110 & 0001 = 0000 (last bit is 0)
//
// - n >>= 1: Shifts bits right by 1 position
//   Example: 0111 >> 1 = 0011
//            0011 >> 1 = 0001
//
// Example trace for n = 7:
//   Binary: 0111
//   Step 1: 0111 & 1 = 1, count = 1, shift → 0011
//   Step 2: 0011 & 1 = 1, count = 2, shift → 0001
//   Step 3: 0001 & 1 = 1, count = 3, shift → 0000
//   Result: 3

func ActiveBits(n int) int {
	// Your code here
	return 0
}
