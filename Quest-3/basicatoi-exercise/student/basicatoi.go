package student

// TODO: Write your code here
// Implement the BasicAtoi function
//
// Hints:
// - Function signature: func BasicAtoi(s string) int
// - s is a string containing only digits
// - Initialize result to 0: var result int
// - Use a for range loop to iterate over characters
// - For each character, convert to digit: int(v) - 48
// - Or use: int(v - '0')
// - Build the number: result = result * 10 + digit
// - Return the result
//
// Example structure:
//   var result int
//   for _, v := range s {
//       result = result*10 + int(v) - 48
//   }
//   return result
//
// Why subtract 48 (or '0')?
//   - '0' in ASCII is 48, '1' is 49, etc.
//   - Subtracting 48 converts character to digit
//
// Example:
//   BasicAtoi("123")
//   // Returns: 123

func BasicAtoi(s string) int {
	// Your code here
	return 0
}
