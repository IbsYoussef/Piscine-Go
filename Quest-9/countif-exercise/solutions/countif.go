package solutions

// CountIf returns the count of elements that pass the test function
// Parameters:
//
//	f: test function that takes a string and returns a bool
//	tab: slice of strings to test
//
// Returns:
//
//	number of elements for which f returns true
func CountIf(f func(string) bool, tab []string) int {
	// Initialize counter
	count := 0

	// Loop through each element
	for _, val := range tab {
		// If function returns true, increment counter
		if f(val) {
			count++
		}
	}

	return count
}
