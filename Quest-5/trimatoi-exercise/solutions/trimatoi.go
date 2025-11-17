package solutions

// Function Breakdown:
//
// TrimAtoi(s string) int:
//   - Extracts all digits from a string
//   - Combines them into an integer
//   - Handles negative sign if it appears before any digit
//   - Returns 0 if no digits found
//
// Approach:
//   - Initialize result = 0, sign = 1
//   - Track whether we've found a digit yet (foundSign)
//   - Loop through each character
//   - If '-' and no digit found yet: set sign = -1
//   - If digit: extract and build number
//   - Return result * sign
//
// Key Concepts:
//   - Number building: result = result * 10 + digit
//   - Digit extraction: int(rune - '0')
//   - Sign handling: Only first '-' before digits matters
//   - State tracking: foundSign prevents later '-' from affecting sign
//
// Number building example:
//   "str123ing45"
//   Start: result = 0
//   Find '1': result = 0*10 + 1 = 1
//   Find '2': result = 1*10 + 2 = 12
//   Find '3': result = 12*10 + 3 = 123
//   Find '4': result = 123*10 + 4 = 1234
//   Find '5': result = 1234*10 + 5 = 12345
//
// Sign handling examples:
//   "sd-x1fa2W3s4" → '-' before digits, so negative: -1234
//   "sdx1-fa2W3s4" → '-' after digit '1', ignored: 1234
//
// Examples:
//   TrimAtoi("12345") = 12345
//   TrimAtoi("str123ing45") = 12345
//   TrimAtoi("012 345") = 12345
//   TrimAtoi("Hello World!") = 0 (no digits)
//   TrimAtoi("sd-x1fa2W3s4") = -1234 (negative)
//   TrimAtoi("sdx1-fa2W3s4") = 1234 (- after digit)

// Status: Required

func TrimAtoi(s string) int {
	result := 0
	runes := []rune(s)
	sign := 1
	foundSign := false

	for i := 0; i < len(runes); i++ {
		// Check for minus sign before any digit
		if runes[i] == '-' && !foundSign {
			sign = -1
		}
		// Check if character is a digit
		if runes[i] >= '0' && runes[i] <= '9' {
			// Extract digit value
			digit := int(runes[i] - '0')
			// Build number: shift left and add new digit
			result = result*10 + digit
			// Mark that we've found at least one digit
			foundSign = true
		}
	}

	return result * sign
}
