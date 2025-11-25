package solutions

// Function Breakdown:
//
// MakeRange(min, max int) []int:
//   - Creates slice with values from min to max-1
//   - Returns nil if min >= max
//   - Uses make (append not allowed)
//
// Approach:
//   - Check edge case: min >= max
//   - Calculate length: max - min
//   - Create slice with make
//   - Fill slice with values using index
//   - Return result
//
// Key Concepts:
//   - make: Pre-allocates slice with fixed size
//   - Direct indexing: result[i] = value
//   - Length calculation: max - min gives number of elements
//   - Offset: min + i gives correct value at each index
//
// Difference from AppendRange:
//   - AppendRange: Uses append (dynamic growth)
//   - MakeRange: Uses make (pre-allocated size)
//
// Examples:
//   MakeRange(5, 10) = [5, 6, 7, 8, 9]
//     length = 10 - 5 = 5
//     [0]=5, [1]=6, [2]=7, [3]=8, [4]=9
//   MakeRange(10, 5) = nil
//   MakeRange(1, 2) = [1]

// Status: Required

func MakeRange(min, max int) []int {
	// Edge case: min >= max returns nil
	if min >= max {
		return nil
	}

	// Calculate length of slice
	length := max - min

	// Create slice with pre-allocated size
	result := make([]int, length)

	// Fill slice with values from min to max-1
	for i := 0; i < length; i++ {
		result[i] = min + i
	}

	return result
}
