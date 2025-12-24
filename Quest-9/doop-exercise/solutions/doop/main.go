package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

const (
	MaxInt        = 9223372036854775807
	MinInt        = -9223372036854775808
	invalidDivide = "No division by 0"
	invalidModulo = "No modulo by 0"
)

// printNumber converts an integer to string and prints it digit by digit
func printNumber(n int) {
	// Handle negative numbers
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// Handle zero
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Extract digits in reverse order
	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}

	// Print digits
	for _, r := range digits {
		z01.PrintRune(r)
	}
}

// printString prints a string character by character
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	// Check exactly 3 arguments
	if len(args) != 3 {
		return
	}

	// Convert to integers
	val1, err1 := strconv.Atoi(args[0])
	operator := args[1]
	val2, err2 := strconv.Atoi(args[2])

	// Check valid numbers
	if err1 != nil || err2 != nil {
		return
	}

	// Check valid operator
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" && operator != "%" {
		return
	}

	result := 0

	switch operator {
	case "+":
		// Check for addition overflow
		if (val2 > 0 && val1 > MaxInt-val2) || (val2 < 0 && val1 < MinInt-val2) {
			return
		}
		result = val1 + val2

	case "-":
		// Check for subtraction overflow
		if (val2 < 0 && val1 > MaxInt+val2) || (val2 > 0 && val1 < MinInt+val2) {
			return
		}
		result = val1 - val2

	case "*":
		// Check for multiplication overflow
		if val1 != 0 && val2 != 0 {
			// Positive * Positive
			if val1 > 0 && val2 > 0 && val1 > MaxInt/val2 {
				return
			}
			// Positive * Negative
			if val1 > 0 && val2 < 0 && val2 < MinInt/val1 {
				return
			}
			// Negative * Positive
			if val1 < 0 && val2 > 0 && val1 < MinInt/val2 {
				return
			}
			// Negative * Negative
			if val1 < 0 && val2 < 0 && val1 < MaxInt/val2 {
				return
			}
		}
		result = val1 * val2

	case "/":
		// Check division by zero
		if val2 == 0 {
			printString(invalidDivide)
			z01.PrintRune('\n')
			return
		}
		// Check for division overflow (MinInt / -1 = overflow)
		if val1 == MinInt && val2 == -1 {
			return
		}
		result = val1 / val2

	case "%":
		// Check modulo by zero
		if val2 == 0 {
			printString(invalidModulo)
			z01.PrintRune('\n')
			return
		}
		result = val1 % val2
	}

	// Print result
	printNumber(result)
	z01.PrintRune('\n')
}
