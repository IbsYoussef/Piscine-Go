package student

// TODO: Write your code here
// Implement the Capitalize function
//
// Hints:
// - Function signature: func Capitalize(s string) string
// - Use a boolean flag: inWord := false
// - Convert to runes: runes := []rune(s)
// - Loop through each character
// - If lowercase AND not in word: capitalize (rune -= 32), set inWord = true
// - If uppercase AND in word: lowercase (rune += 32)
// - If digit: set inWord = true (digits are part of words)
// - If non-alphanumeric: check if still in word or reset
//
// Word definition:
// - Alphanumeric characters (A-Z, a-z, 0-9)
// - Separated by non-alphanumeric (spaces, symbols, etc.)
//
// Examples:
//   Capitalize("Hello! How are you?") returns "Hello! How Are You?"
//   Capitalize("how+are+things+4you?") returns "How+Are+Things+4you?"
//   Capitalize("hello") returns "Hello"

func Capitalize(s string) string {
	// Your code here
	return ""
}
