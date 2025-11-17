package student

// TODO: Write your code here
// Implement the ToLower function
//
// Hints:
// - Function signature: func ToLower(s string) string
// - Convert to runes: runes := []rune(s)
// - Loop: for i := 0; i < len(runes); i++
// - Check if uppercase: if runes[i] >= 'A' && runes[i] <= 'Z'
// - Convert to lowercase: runes[i] += 32
// - Convert back to string: return string(runes)
//
// ASCII conversion:
// - 'A' (65) + 32 = 'a' (97)
// - 'Z' (90) + 32 = 'z' (122)
// - Difference between uppercase and lowercase is always 32
//
// Examples:
//   ToLower("Hello! How are you?") returns "hello! how are you?"
//   ToLower("ABC123") returns "abc123"
//   ToLower("TEST") returns "test"

func ToLower(s string) string {
	// Your code here
	return ""
}
