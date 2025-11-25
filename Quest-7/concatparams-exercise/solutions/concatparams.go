package solutions

// Function Breakdown:
//
// ConcatParams(args []string) string:
//   - Concatenates all strings with newlines between them
//   - No newline after last element
//
// Approach:
//   - Initialize empty result string
//   - Loop through each argument
//   - Add argument to result
//   - Add newline if not last element
//   - Return result
//
// Key Concepts:
//   - String concatenation: result += value
//   - Index checking: i < len(args)-1
//   - Newline character: \n
//
// Examples:
//   ConcatParams(["Hello", "how", "are", "you?"]) = "Hello\nhow\nare\nyou?"
//   ConcatParams(["one"]) = "one"
//   ConcatParams([]) = ""

// Status: Required

func ConcatParams(args []string) string {
	result := ""

	// Loop through each argument with index
	for i, v := range args {
		// Add the argument
		result += v

		// Add newline if not last element
		if i < len(args)-1 {
			result += "\n"
		}
	}

	return result
}
