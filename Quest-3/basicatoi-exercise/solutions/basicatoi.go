package solutions

// Function Breakdown:
//
// BasicAtoi(s string) int:
//   - Takes a string of digits
//   - Converts to an integer
//   - Returns the integer value
//
// Approach:
//   - Initialize result to 0
//   - Iterate through each character
//   - Multiply result by 10 (shift left)
//   - Add current digit value
//   - Return final result
//
// Key Concepts:
//   - ASCII arithmetic: char - '0' = digit
//   - Or: int(char) - 48 = digit
//   - Building numbers: result * 10 + digit
//   - Base 10 system
//
// Example:
//   BasicAtoi("123")
//   // Returns: 123
//
// Step by step for "42":
//   - Start: result = 0
//   - '4': result = 0*10 + 4 = 4
//   - '2': result = 4*10 + 2 = 42
//   - Return: 42

func BasicAtoi(s string) int {
	var result int
	for _, v := range s {
		result = result*10 + int(v) - 48
	}
	return result
}
