package solutions

// Function Breakdown:
//
// Join(strs []string, sep string) string:
//   - Takes a slice of strings and a separator
//   - Joins all strings with separator between them
//   - Returns the concatenated result
//
// Approach:
//   - Handle edge case: empty slice → return ""
//   - Initialize empty result string
//   - Loop through elements with index
//   - Add separator before each element (except first)
//   - Concatenate element to result
//   - Return final string
//
// Key Concepts:
//   - Edge case handling: Empty and single-element slices
//   - Index-based logic: Skip separator for first element
//   - String concatenation: Building result progressively
//
// How it works:
//   Example: []string{"A", "B", "C"} with separator ":"
//   - i=0: result = "" + "A" = "A"
//   - i=1: result = "A" + ":" + "B" = "A:B"
//   - i=2: result = "A:B" + ":" + "C" = "A:B:C"
//
// Edge cases:
//   - Empty slice: return ""
//   - Single element: return just that element
//   - No elements need separator
//
// Examples:
//   Join([]string{"Hello!", " How", " are", " you?"}, ":") = "Hello!: How: are: you?"
//   Join([]string{"Go", "Lang"}, "-") = "Go-Lang"
//   Join([]string{}, ",") = ""
//   Join([]string{"single"}, ",") = "single"
//   Join([]string{"a", "b", "c"}, "") = "abc"

// Status: Optional

func Join(strs []string, sep string) string {
	result := ""

	// Edge case: empty slice
	if len(strs) == 0 {
		return ""
	}

	// Loop through each element with index
	for i, v := range strs {
		// Add separator before each element (except first)
		if i > 0 {
			result += sep
		}
		// Add the element
		result += v
	}

	return result
}
