package student

// TODO: Write your code here
// Implement the IsAlpha function
//
// Hints:
// - Function signature: func IsAlpha(s string) bool
// - Convert to runes: runes := []rune(s)
// - Loop: for i := 0; i < len(runes); i++
// - Check if NOT alphanumeric:
//   if !(runes[i] >= 'A' && runes[i] <= 'Z') &&
//      !(runes[i] >= 'a' && runes[i] <= 'z') &&
//      !(runes[i] >= '0' && runes[i] <= '9')
// - If not alphanumeric: return false
// - After loop: return true
//
// Alphanumeric means:
// - Uppercase: A-Z (65-90)
// - Lowercase: a-z (97-122)
// - Digits: 0-9 (48-57)
//
// Examples:
//   IsAlpha("HelloHowareyou") returns true
//   IsAlpha("Hello! How are you?") returns false
//   IsAlpha("Whatsthis4") returns true
//   IsAlpha("") returns true

func IsAlpha(s string) bool {
	// Your code here
	return false
}
