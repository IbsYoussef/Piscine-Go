package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}

	digits := []int{}

	for n != 0 {
		extract := n % 10
		digits = append(digits, extract)
		n = n / 10
	}

	for i := 0; i < len(digits)-1; i++ {
		swapped := false
		for j := 0; j < len(digits)-1-i; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	for i := 0; i < len(digits); i++ {
		z01.PrintRune(rune(digits[i]) + '0')
	}
}
