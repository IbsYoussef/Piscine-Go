package piscine

// Recursive Factorial is a function that find the factorial of a given a value using recursion
// Status: Required

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}

	if nb == 0 || nb == 1 {
		return 1
	}

	return nb * RecursiveFactorial(nb-1)
}
