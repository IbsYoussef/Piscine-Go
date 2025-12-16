package main

// TODO: Complete the code to make it work correctly
//
// This code compiles but is incomplete!
// You need to fill in the missing parts.
//
// Tasks:
// 1. The point struct is defined but you need to understand it
// 2. The setPoint() function works - study how it uses pointers
// 3. Complete the printStr() function
// 4. Complete the printNumber() function to convert int to printed digits
// 5. In main(), use your functions to print: "x = 42, y = 21\n"
//
// Note: You CANNOT use fmt package - only z01.PrintRune!
//
// Hints for printNumber():
// - Handle 0 as special case
// - Extract digits using n % 10 and n / 10
// - Convert digit to rune: '0' + rune(digit)
// - Build digits in correct order, then print them

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	// TODO: Loop through string and print each character using z01.PrintRune
}

func printNumber(n int) {
	// TODO: Convert the integer to individual digits and print them
	// Hint: You'll need to handle digits and convert them to runes
	// For now, this does nothing - you need to implement it!
}

func main() {
	points := &point{}

	setPoint(points)

	// TODO: Use printStr() and printNumber() to output: "x = 42, y = 21\n"
	// Example structure:
	// printStr("x = ")
	// printNumber(points.x)
	// printStr(", y = ")
	// printNumber(points.y)
	// z01.PrintRune('\n')
}
