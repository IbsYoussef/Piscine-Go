package student

// TODO: Write your code here
// Implement the IsPrintable function
//
// Hints:
// - Function signature: func IsPrintable(s string) bool
// - Convert to runes: runes := []rune(s)
// - Loop: for _, r := range runes
// - Check NOT printable: if !(r >= 32 && r <= 126)
// - If not printable: return false
// - After loop: return true (all were printable)
//
// Printable range: ASCII 32 (space) to 126 (tilde ~)
//
// Non-printable examples:
// - \n (newline) = ASCII 10
// - \t (tab) = ASCII 9
// - \r (carriage return) = ASCII 13
//
// Examples:
//   IsPrintable("Hello") returns true
//   IsPrintable("Hello\n") returns false
//   IsPrintable("abc 123") returns true
//   IsPrintable("") returns true

func IsPrintable(s string) bool {
	// Your code here
	return false
}
