package solutions

// Map applies a function to each element and returns a slice of results
// Parameters:
//
//	f: function that takes an int and returns a bool
//	a: slice of integers
//
// Returns:
//
//	slice of booleans (results of applying f to each element)
func Map(f func(int) bool, a []int) []bool {
	// Initialize empty result slice
	result := []bool{}

	// Loop through each element
	for _, val := range a {
		// Apply function and append result to output slice
		result = append(result, f(val))
	}

	return result
}
