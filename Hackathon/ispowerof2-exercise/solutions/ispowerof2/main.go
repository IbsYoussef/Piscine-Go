package main

import (
	"fmt"
	"os"
)

// IsPowerOfTwo checks if x is a power of 2 using bitwise operation
//
// How it works:
// - Powers of 2 in binary have exactly ONE bit set (e.g., 8 = 1000)
// - Subtracting 1 flips that bit and all bits to the right (8-1 = 0111)
// - ANDing them gives 0 if and only if x is a power of 2
//
// Examples:
//
//	8 (1000) & 7 (0111) = 0 → power of 2
//	6 (0110) & 5 (0101) = 4 → NOT a power of 2
func IsPowerOfTwo(x int) bool {
	return x > 0 && (x&(x-1)) == 0
}

func main() {
	args := os.Args[1:]

	// Check for exactly one argument
	if len(args) != 1 {
		return
	}

	// Convert string to integer manually
	num := args[0]
	result := 0

	for i := 0; i < len(num); i++ {
		result = result*10 + int(num[i]-'0')
	}

	// Check and print result
	if IsPowerOfTwo(result) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
