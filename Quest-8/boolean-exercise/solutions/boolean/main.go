package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program Breakdown:
//
// Goal: Check if number of command-line arguments is even or odd
//
// Key Fixes from broken code:
// 1. boolean → bool (correct Go type)
// 2. yes/no → true/false (correct boolean values)
// 3. Added even() function to check if number is even
// 4. lengthOfArg → len(os.Args) - 1 (count arguments)
// 5. Added EvenMsg and OddMsg constants
// 6. Removed == 1 comparisons (booleans don't need that)
//
// Algorithm:
// 1. Count arguments: len(os.Args) - 1 (exclude program name)
// 2. Check if count is even using modulo operator
// 3. Print appropriate message
//
// Examples:
//   go run . "not" "odd" → 2 args → even → "I have an even number of arguments"
//   go run . "not even" → 1 arg → odd → "I have an odd number of arguments"

// Status: Required

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	return even(nbr)
}

func even(nbr int) bool {
	return nbr%2 == 0
}

func main() {
	// Count arguments (excluding program name)
	argCount := len(os.Args) - 1

	if isEven(argCount) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
