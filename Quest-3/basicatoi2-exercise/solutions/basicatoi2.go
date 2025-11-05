package solutions

// Function Breakdown:
//
// BasicAtoi2(s string) int:
//   - Takes a string that may contain non-digits
//   - Validates each character
//   - Returns 0 if any non-digit found
//   - Returns converted integer if all valid
//
// Approach:
//   - Initialize result to 0
//   - Iterate through each character
//   - Check if char is digit: char >= '0' && char <= '9'
//   - If not valid, return 0 immediately
//   - If valid, convert and build number
//   - Return final result
//
// Key Concepts:
//   - Input validation
//   - Early return on invalid input
//   - Character range checking
//   - ASCII arithmetic: char - '0'
//
// Example:
//   BasicAtoi2("123")   // Returns: 123
//   BasicAtoi2("12 3")  // Returns: 0 (space is invalid)
//
// Logic explanation:
//   - 'result * 10' shifts current digits left
//   - 'int(char-'0')' converts char to digit
//   - Combined: builds the number digit by digit

func BasicAtoi2(s string) int {
	result := 0

	// Iterate over s string with range loop
	for _, char := range s {
		// Check to see if character is valid, if not we return 0
		if char < '0' || char > '9' {
			return 0
		}

		// Convert the character to its integer value and update the result
		// The logic: 'result*10' shifts the current digits to the left (base 10),
		// making room for the new digit. 'int(char-'0')' converts the character
		// to its integer value (e.g., '3' -> 3) by subtracting the ASCII value of '0'.
		result = result*10 + int(char-'0')
	}

	// Return value
	return result
}
