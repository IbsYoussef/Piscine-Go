package main

import (
	"github.com/01-edu/z01"
)

// Program Breakdown:
//
// Goal: Create a struct, modify it via pointer, and print its values
//
// Key Concepts:
// 1. Struct definition: type point struct { x int; y int }
// 2. Pointer usage: *point to modify original struct
// 3. Manual printing: Cannot use fmt, must use z01.PrintRune
// 4. Number printing: Convert int to individual digit runes
//
// Algorithm:
// 1. Define point struct with x and y fields
// 2. Create setPoint function that modifies via pointer
// 3. Create helper functions for printing strings and numbers
// 4. Print formatted output: "x = 42, y = 21\n"
//
// Printing numbers without fmt:
// - Extract digits using modulo and division
// - Convert digits to runes ('0' + digit)
// - Print each rune

// Status: Required

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Extract digits in reverse order
	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{'0' + rune(digit)}, digits...)
		n /= 10
	}

	// Print digits
	for _, r := range digits {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = ")
	printNumber(points.x)
	printStr(", y = ")
	printNumber(points.y)
	z01.PrintRune('\n')
}
