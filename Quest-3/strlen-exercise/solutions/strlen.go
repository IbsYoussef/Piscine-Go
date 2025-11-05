package solutions

// Function Breakdown:
//
// StrLen(s string) int:
//   - Takes a string as parameter
//   - Counts the number of runes (characters)
//   - Returns the count as an int
//
// Approach:
//   - Initialize counter to 0
//   - Use for range to iterate through string
//   - Increment counter for each rune
//   - Return the final count
//
// Key Concepts:
//   - Runes: Go's character type (Unicode code points)
//   - Range iteration: Automatically handles UTF-8
//   - Accumulator pattern: Building up a result
//
// Example:
//   StrLen("Hello") returns 5
//   StrLen("") returns 0
//   StrLen("Hello World!") returns 12

// StrLen is a function that returns the length of a string.
// Status: Optional

func StrLen(s string) int {
	count := 0
	// Iterate through each rune in the string
	// We don't need the rune value, so we use blank identifier
	for range s {
		count++
	}
	return count
}
