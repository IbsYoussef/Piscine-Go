package student

// TODO: Write your code here
// Implement the IsNumeric function
//
// Hints:
// - Function signature: func IsNumeric(s string) bool
// - Handle edge case: if s == "" { return false }
// - Convert to runes: runes := []rune(s)
// - Loop: for i := 0; i < len(runes); i++
// - Check NOT digit: if !(runes[i] >= '0' && runes[i] <= '9')
// - If not digit: return false
// - After loop: return true (all were digits)
//
// Digit range: '0' (48) to '9' (57)
//
// Examples:
//   IsNumeric("010203") returns true
//   IsNumeric("01,02,03") returns false
//   IsNumeric("123abc") returns false
//   IsNumeric("") returns false

func IsNumeric(s string) bool {
	// Your code here
	return false
}
