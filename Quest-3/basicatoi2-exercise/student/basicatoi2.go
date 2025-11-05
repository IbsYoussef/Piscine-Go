package student

// TODO: Write your code here
// Implement the BasicAtoi2 function
//
// Hints:
// - Function signature: func BasicAtoi2(s string) int
// - s may contain non-digit characters
// - Initialize result to 0: result := 0
// - Use a for range loop to iterate over characters
// - For each character, check if it's a digit:
//   if char < '0' || char > '9' {
//       return 0
//   }
// - If valid, convert and build number:
//   result = result * 10 + int(char - '0')
// - Return the result
//
// Example structure:
//   result := 0
//   for _, char := range s {
//       if char < '0' || char > '9' {
//           return 0
//       }
//       result = result*10 + int(char-'0')
//   }
//   return result
//
// Example:
//   BasicAtoi2("123")     // Returns: 123
//   BasicAtoi2("12 3")    // Returns: 0 (space is invalid)
//   BasicAtoi2("Hello")   // Returns: 0 (letters are invalid)

func BasicAtoi2(s string) int {
	// Your code here
	return 0
}
