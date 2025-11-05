package solutions

// Function Breakdown:
//
// Atoi(s string) int:
//   - Takes a string that may have +/- sign and digits
//   - Validates and converts to integer
//   - Returns 0 if invalid
//   - Handles positive and negative numbers
//
// Approach:
//   - Initialize result to 0, sign to 1
//   - Check if string is empty
//   - Check first character for sign
//   - Set sign and start index accordingly
//   - Validate all remaining characters are digits
//   - Build number digit by digit
//   - Apply sign to result
//
// Key Concepts:
//   - Sign handling: +/- at start only
//   - String indexing: s[0], s[i]
//   - Character validation
//   - Start index to skip sign
//   - Negative numbers: multiply by -1
//
// Example:
//   Atoi("123")    // Returns: 123
//   Atoi("-123")   // Returns: -123
//   Atoi("+123")   // Returns: 123
//   Atoi("++123")  // Returns: 0 (invalid)
//   Atoi("12 3")   // Returns: 0 (invalid)

func Atoi(s string) int {
	result := 0
	sign := 1
	startIndex := 0

	// Check if the string is not empty
	if len(s) > 0 {
		// Check for sign at the beginning
		if s[0] == '-' {
			sign = -1
			startIndex = 1
		} else if s[0] == '+' {
			startIndex = 1
		}
	}

	// Iterate over the string starting from startIndex
	for i := startIndex; i < len(s); i++ {
		char := s[i]

		// Check if the character is a digit
		if char < '0' || char > '9' {
			return 0 // Return integer 0 for any non-digit character
		}

		// Convert the character to its integer value and update the result
		result = result*10 + int(char-'0')
	}

	// Apply the sign to the result
	return result * sign
}
