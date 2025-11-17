package student

// TODO: Write your code here
// Implement the ToUpper function
//
// Hints:
// - Function signature: func ToUpper(s string) string
// - Convert to runes: runes := []rune(s)
// - Loop: for i := 0; i < len(runes); i++
// - Check if lowercase: if runes[i] >= 'a' && runes[i] <= 'z'
// - Convert to uppercase: runes[i] -= 32
// - Convert back to string: return string(runes)
//
// ASCII conversion:
// - 'a' (97) - 32 = 'A' (65)
// - 'z' (122) - 32 = 'Z' (90)
// - Difference between lowercase and uppercase is always 32
//
// Examples:
//   ToUpper("Hello! How are you?") returns "HELLO! HOW ARE YOU?"
//   ToUpper("abc123") returns "ABC123"
//   ToUpper("test") returns "TEST"

func ToUpper(s string) string {
	// Your code here
	return ""
}
