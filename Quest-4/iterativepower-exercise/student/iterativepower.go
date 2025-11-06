package student

// TODO: Write your code here
// Implement the IterativePower function
//
// Hints:
// - Function signature: func IterativePower(nb int, power int) int
// - Handle negative power: if power < 0 { return 0 }
// - Handle power == 0: if power == 0 { return 1 }
// - Initialize: result := 1
// - Loop: for i := 1; i <= power; i++
// - Multiply: result = result * nb
// - Return result

func IterativePower(nb int, power int) int {
	// Your code here
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
