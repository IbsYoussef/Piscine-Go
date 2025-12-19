package solutions

// ForEach applies a function to each element in a slice
// Parameters:
//
//	f: function that takes an int
//	a: slice of integers
func ForEach(f func(int), a []int) {
	// Loop through each element in the slice
	for _, val := range a {
		// Apply the function to the current element
		f(val)
	}
}
