package solutions

// Status: Required

func IterativePower(nb int, power int) int {
	// Negative power returns 0
	if power < 0 {
		return 0
	}

	// Any number to power 0 is 1
	if power == 0 {
		return 1
	}

	// Calculate nb^power iteratively
	result := 1
	for i := 1; i <= power; i++ {
		result = result * nb
	}

	return result
}
