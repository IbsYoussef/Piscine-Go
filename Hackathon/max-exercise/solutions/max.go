package solutions

// Max returns the maximum value in a slice of integers
// Returns 0 if the slice is empty
//
// Algorithm:
// 1. Handle edge case: empty slice returns 0
// 2. Initialize max to first element (important for negative numbers!)
// 3. Loop through remaining elements
// 4. Update max if current element is larger
// 5. Return max
//
// Edge cases:
// - Empty slice: []int{} → 0
// - All negative: []int{-5, -10, -3} → -3
// - Single element: []int{42} → 42
// - All same: []int{5, 5, 5} → 5

func Max(a []int) int {
	// Handle empty slice
	if len(a) == 0 {
		return 0
	}

	// Initialize max to first element
	// This handles negative numbers correctly
	max := a[0]

	// Compare with remaining elements
	for _, num := range a {
		if num > max {
			max = num
		}
	}

	return max
}
