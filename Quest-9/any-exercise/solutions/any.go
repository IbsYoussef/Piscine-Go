package solutions

// Any returns true if at least one element passes the test function
// Parameters:
//
//	f: test function that takes a string and returns a bool
//	a: slice of strings to test
//
// Returns:
//
//	true if at least one element returns true when passed to f
//	false if no elements return true (or slice is empty)
func Any(f func(string) bool, a []string) bool {
	// Loop through each element
	for _, val := range a {
		// If function returns true for this element, return true immediately
		if f(val) {
			return true
		}
	}
	// No element returned true, so return false
	return false
}
