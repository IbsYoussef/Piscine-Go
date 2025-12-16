package main

import (
	"os"

	"github.com/01-edu/z01"
)

// TODO: Fix the code below to make it work correctly
//
// This code compiles but doesn't work properly!
// You need to fix the logic to match the expected behavior.
//
// Issues to fix:
// 1. isEven() always returns false - fix the logic
// 2. The argument counting is wrong - should count args excluding program name
// 3. The messages are wrong - fix the strings
// 4. The even() function doesn't exist - implement it
//
// Expected behavior:
// - Count command-line arguments (excluding program name)
// - Check if count is even or odd
// - Print: "I have an even number of arguments" or "I have an odd number of arguments"

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	// TODO: Fix this function - it always returns false!
	// Hint: Use the even() function (which you need to create)
	return false
}

// TODO: Create the even() function here
// It should return true if nbr is even, false if odd
// Hint: Use modulo operator (%)

func main() {
	// TODO: Fix the argument counting - this is wrong!
	argCount := len(os.Args) // Wrong! Should exclude program name

	if isEven(argCount) {
		// TODO: Fix these messages!
		printStr("Wrong message")
	} else {
		printStr("Wrong message")
	}
}
