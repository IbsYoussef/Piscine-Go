package solutions

// AdvancedSortWordArr sorts a string slice using a custom comparison function
// Parameters:
//
//	a: slice of strings to sort in-place
//	f: comparison function that returns:
//	   positive int if a > b
//	   0 if a == b
//	   negative int if a < b
//
// This is a flexible sorting function that delegates comparison logic
// to the function parameter, allowing different sorting behaviors

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	// Outer loop: controls number of passes
	for i := 0; i < len(a)-1; i++ {
		// Inner loop: compares adjacent elements using comparison function
		for j := 0; j < len(a)-1-i; j++ {
			// Use comparison function instead of direct comparison
			// If f returns positive, a[j] should come after a[j+1], so swap
			if f(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
