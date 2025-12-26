package main

import (
	"fmt"
	"os"
)

// TODO: Implement IsPowerOfTwo function
// Check if x is a power of 2
//
// Hint: Use bitwise operation
// Powers of 2 have exactly one bit set in binary
// Formula: x > 0 && (x & (x-1)) == 0
//
// Why this works:
// - 8 in binary:  1000
// - 7 in binary:  0111
// - 8 & 7:        0000 (zero means power of 2!)

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
