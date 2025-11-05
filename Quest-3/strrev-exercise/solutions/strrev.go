package solutions

// Function Breakdown:
//
// StrRev(s string) string:
//   - Takes a string as parameter
//   - Reverses the order of characters
//   - Returns the reversed string
//
// Approach:
//   - Initialize empty result string
//   - Iterate through input string
//   - For each character, prepend it to result
//   - This builds the reversed string
//
// Key Concepts:
//   - String building: Creating new strings
//   - Prepending: Adding to the front
//   - String concatenation: result = new + old
//   - Rune conversion: string(c) converts rune to string
//
// Example:
//   StrRev("abc") returns "cba"
//   StrRev("Hello") returns "olleH"
//   StrRev("") returns ""

// StrRev is a function that returns the reverse of a string
// Status: Required

func StrRev(s string) string {
	var result string
	// Iterate through each character in the string
	for _, c := range s {
		// Prepend current character to result
		// This effectively reverses the string
		result = string(c) + result
	}
	return result
}
