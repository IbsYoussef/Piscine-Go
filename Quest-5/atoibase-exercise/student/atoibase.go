package student

// TODO: Write your code here
// Implement the AtoiBase function
//
// This is the reverse of PrintNbrBase!
//
// Hints:
// - Validate base first (same rules as PrintNbrBase)
// - Check if all characters in s exist in base
// - Use Horner's method for conversion:
//   result = 0
//   for each character in s:
//       find index of character in base (= digit value)
//       result = result * len(base) + digit_value
//
// Example walkthrough:
//   AtoiBase("125", "0123456789")
//   '1' is at index 1: result = 0*10 + 1 = 1
//   '2' is at index 2: result = 1*10 + 2 = 12
//   '5' is at index 5: result = 12*10 + 5 = 125
//
// Helper functions you might need:
// - isValidBase(base string) bool
// - indexOf(base string, char rune) int
// - contains(base string, char rune) bool
//
// Examples:
//   AtoiBase("125", "0123456789") returns 125
//   AtoiBase("1111101", "01") returns 125 (binary)
//   AtoiBase("7D", "0123456789ABCDEF") returns 125 (hex)
//   AtoiBase("bbbbbab", "-ab") returns 0 (invalid base)

func AtoiBase(s string, base string) int {
	// Your code here
	return 0
}
