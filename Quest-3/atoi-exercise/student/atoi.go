package student

// TODO: Write your code here
// Implement the Atoi function
//
// Hints:
// - Function signature: func Atoi(s string) int
// - s may start with + or - sign
// - Initialize: result := 0, sign := 1, startIndex := 0
// - Check if string is not empty: if len(s) > 0
// - Check first character:
//   if s[0] == '-' {
//       sign = -1
//       startIndex = 1
//   } else if s[0] == '+' {
//       startIndex = 1
//   }
// - Loop from startIndex to end of string
// - Validate each character is a digit
// - If invalid, return 0
// - Build the number: result = result*10 + int(char-'0')
// - Apply sign at the end: return result * sign
//
// Example structure:
//   result := 0
//   sign := 1
//   startIndex := 0
//
//   if len(s) > 0 {
//       if s[0] == '-' {
//           sign = -1
//           startIndex = 1
//       } else if s[0] == '+' {
//           startIndex = 1
//       }
//   }
//
//   for i := startIndex; i < len(s); i++ {
//       char := s[i]
//       if char < '0' || char > '9' {
//           return 0
//       }
//       result = result*10 + int(char-'0')
//   }
//
//   return result * sign
//
// Example:
//   Atoi("123")    // Returns: 123
//   Atoi("-123")   // Returns: -123
//   Atoi("++123")  // Returns: 0

func Atoi(s string) int {
	// Your code here
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
