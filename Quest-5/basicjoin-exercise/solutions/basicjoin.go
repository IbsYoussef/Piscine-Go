package solutions

// Function Breakdown:
//
// BasicJoin(elems []string) string:
//   - Takes a slice of strings
//   - Joins them together without separator
//   - Returns the concatenated result
//
// Approach:
//   - Initialize empty result string
//   - Loop through each element
//   - Concatenate each to result
//   - Return final string
//
// Key Concepts:
//   - String concatenation: Using += operator
//   - Range loop: Iterate over slice
//   - Accumulator: Build result progressively
//   - Simple and straightforward
//
// How it works:
//   - Start with result = ""
//   - For each string in slice: result += string
//   - Example: "" + "Hello!" = "Hello!"
//   - Then: "Hello!" + " How" = "Hello! How"
//   - Continue until all elements processed
//
// Examples:
//   BasicJoin([]string{"Hello!", " How", " are", " you?"}) = "Hello! How are you?"
//   BasicJoin([]string{"Go", "Lang"}) = "GoLang"
//   BasicJoin([]string{"a", "b", "c"}) = "abc"
//   BasicJoin([]string{}) = ""
//   BasicJoin([]string{"single"}) = "single"

// Status: Optional

func BasicJoin(elems []string) string {
	result := ""

	// Loop through each element in the slice
	for _, v := range elems {
		// Concatenate to result
		result += v
	}

	return result
}
