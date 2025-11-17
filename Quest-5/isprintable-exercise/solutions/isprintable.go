package solutions

// Function Breakdown:
//
// IsPrintable(s string) bool:
//   - Checks if ALL characters are printable
//   - Printable = ASCII 32 to 126
//   - Returns true for empty string
//   - Returns false if any character is not printable
//
// Approach:
//   - Convert string to rune slice
//   - Loop through each rune
//   - Check if rune is NOT in printable range
//   - If not printable, return false immediately
//   - If loop completes, all are printable → return true
//
// Key Concepts:
//   - Printable characters: Visible characters (ASCII 32-126)
//   - Control characters: Non-printable (ASCII 0-31, 127+)
//   - Range check: 32 (space) to 126 (tilde ~)
//   - Early return: Stop on first non-printable
//
// ASCII ranges:
//   - 0-31: Control characters (not printable)
//   - 32: Space (first printable)
//   - 33-126: Visible characters (printable)
//   - 127+: Extended ASCII / Control (not printable)
//
// Common non-printable characters:
//   - \n (10): Newline
//   - \t (9): Tab
//   - \r (13): Carriage return
//   - \0 (0): Null character
//
// Examples:
//   IsPrintable("Hello") = true (all printable)
//   IsPrintable("Hello\n") = false (newline not printable)
//   IsPrintable("abc 123") = true (space is printable)
//   IsPrintable("test\t") = false (tab not printable)
//   IsPrintable("") = true (empty string)

// Status: Required

func IsPrintable(s string) bool {
	// Convert string to rune slice
	runes := []rune(s)

	// Check each character
	for _, r := range runes {
		// If NOT in printable range, return false
		if !(r >= 32 && r <= 126) {
			return false
		}
	}

	// All characters are printable (or string is empty)
	return true
}
