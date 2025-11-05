package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	baseLen := len(base)
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// Create a slice for the result
	var result []rune

	// Conversion loop
	for nbr > 0 {
		remainder := nbr % baseLen
		result = append(result, rune(base[remainder]))
		nbr = nbr / baseLen
	}

	// Print the result in reverse order
	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
