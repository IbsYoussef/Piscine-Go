package solutions

// Function Breakdown:
//
// Concat(str1 string, str2 string) string:
//   - Takes two strings as parameters
//   - Joins them together (concatenation)
//   - Returns the combined string
//
// Approach:
//   - Use the + operator to concatenate
//   - Return the result
//
// Key Concepts:
//   - String concatenation: Joining strings
//   - + operator: Works for string concatenation in Go
//   - Immutability: Creates new string, doesn't modify originals
//   - Simple and straightforward!
//
// How it works:
//   - The + operator creates a new string
//   - Copies str1's content
//   - Appends str2's content
//   - Returns the new combined string
//
// Examples:
//   Concat("Hello!", " How are you?") = "Hello! How are you?"
//   Concat("Go", "Lang") = "GoLang"
//   Concat("", "test") = "test"
//   Concat("test", "") = "test"
//   Concat("", "") = ""

// Status: Required

func Concat(str1 string, str2 string) string {
	// Simply concatenate using + operator
	return str1 + str2
}
