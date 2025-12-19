package solutions

// IsSorted checks if a slice is sorted (ascending or descending)
// Parameters:
//
//	f: comparison function that returns:
//	   positive int if a > b
//	   0 if a == b
//	   negative int if a < b
//	a: slice of integers to check
//
// Returns:
//
//	true if array is sorted (ascending or descending), false otherwise
func IsSorted(f func(a, b int) int, a []int) bool {
	// Arrays with 0 or 1 elements are always sorted
	if len(a) < 2 {
		return true
	}

	// STEP 1: Determine direction by finding first non-equal pair
	ascending := false
	for i := 0; i < len(a)-1; i++ {
		result := f(a[i], a[i+1])
		if result > 0 {
			// Current > Next → descending order
			ascending = false
			break
		} else if result < 0 {
			// Current < Next → ascending order
			ascending = true
			break
		}
		// result == 0 → equal elements, keep looking
	}

	// STEP 2: Verify all pairs follow the same direction
	for i := 0; i < len(a)-1; i++ {
		if ascending {
			// In ascending order, no element should be greater than next
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		} else {
			// In descending order, no element should be less than next
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	}

	// All checks passed - array is sorted
	return true
}
