package solutions

import "github.com/01-edu/z01"

// Function Breakdown:
//
// PrintNbrInOrder(n int):
//   - Extracts digits from integer
//   - Sorts digits in ascending order
//   - Prints the sorted digits
//
// Approach:
//   - Handle special case: n == 0
//   - Extract all digits using modulo and division
//   - Store digits in slice
//   - Sort digits using bubble sort
//   - Print each digit
//
// Key Concepts:
//   - Digit extraction: n % 10 gives last digit, n / 10 removes it
//   - Bubble sort: Compare adjacent elements and swap if needed
//   - Rune conversion: int to printable character
//
// Digit extraction example:
//   n = 321
//   First: 321 % 10 = 1, then n = 321 / 10 = 32
//   Second: 32 % 10 = 2, then n = 32 / 10 = 3
//   Third: 3 % 10 = 3, then n = 3 / 10 = 0
//   Result: [1, 2, 3]
//
// Bubble sort:
//   - Compare adjacent elements
//   - Swap if left > right
//   - Repeat until sorted
//   - Optimization: use swapped flag to stop early
//
// Printing:
//   - Convert int digit to rune: rune(digit) + '0'
//   - Example: 1 → rune(1) + '0' → '1'
//
// Examples:
//   PrintNbrInOrder(321) prints "123"
//   PrintNbrInOrder(0) prints "0"
//   PrintNbrInOrder(54321) prints "12345"

// Status: Required

func PrintNbrInOrder(n int) {
	// Special case: zero
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Extract all digits
	digits := []int{}
	for n != 0 {
		extract := n % 10 // Get last digit
		digits = append(digits, extract)
		n = n / 10 // Remove last digit
	}

	// Bubble sort: sort digits in ascending order
	for i := 0; i < len(digits)-1; i++ {
		swapped := false
		for j := 0; j < len(digits)-1-i; j++ {
			if digits[j] > digits[j+1] {
				// Swap adjacent elements
				digits[j], digits[j+1] = digits[j+1], digits[j]
				swapped = true
			}
		}
		// Early exit if no swaps made (already sorted)
		if !swapped {
			break
		}
	}

	// Print each sorted digit
	for i := 0; i < len(digits); i++ {
		z01.PrintRune(rune(digits[i]) + '0')
	}
}
