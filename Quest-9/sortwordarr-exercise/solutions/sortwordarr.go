package solutions

// SortWordArr sorts a string slice in ascending ASCII order using bubble sort
// Parameters:
//
//	a: slice of strings to sort in-place
//
// The function modifies the original slice (no return value)
//
// ASCII ordering:
//
//	Numbers (0-9) come before uppercase letters (A-Z)
//	Uppercase letters come before lowercase letters (a-z)
//	Example: "1" < "A" < "a"

func SortWordArr(a []string) {
	// Outer loop: controls number of passes
	// Need len(a)-1 passes to ensure all elements are sorted
	for i := 0; i < len(a)-1; i++ {
		// Inner loop: compares adjacent elements and swaps if needed
		// len(a)-1-i because after each pass, one more element is in final position
		for j := 0; j < len(a)-1-i; j++ {
			// Compare adjacent strings
			// Go's > operator compares strings by ASCII values
			if a[j] > a[j+1] {
				// Swap elements
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
