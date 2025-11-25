package solutions

// Function Breakdown:
//
// AppendRange(min, max int) []int:
//   - Creates slice with values from min to max-1
//   - Returns nil if min >= max
//   - Uses append (make not allowed)
//
// Approach:
//   - Check edge case: min >= max
//   - Initialize empty slice
//   - Loop from min to max-1
//   - Append each value to slice
//   - Return result
//
// Key Concepts:
//   - Slice initialization: []int{}
//   - Nil slice: []int(nil)
//   - Append: Dynamic array growth
//   - Range: [min, max) - includes min, excludes max
//
// Examples:
//   AppendRange(5, 10) = [5, 6, 7, 8, 9]
//   AppendRange(10, 5) = [] (nil)
//   AppendRange(1, 2) = [1]
//   AppendRange(5, 5) = [] (nil)

// Status: Required

func AppendRange(min, max int) []int {
	result := []int{}

	// Edge case: min >= max returns nil
	if min >= max {
		return []int(nil)
	}

	// Loop from min to max-1 (max is excluded)
	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}
