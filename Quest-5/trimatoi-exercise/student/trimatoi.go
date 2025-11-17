package student

// TODO: Write your code here
// Implement the TrimAtoi function
//
// Hints:
// - Function signature: func TrimAtoi(s string) int
// - Initialize: result := 0, sign := 1, foundSign := false
// - Convert to runes: runes := []rune(s)
// - Loop: for i := 0; i < len(runes); i++
// - Check for '-' before digits: if runes[i] == '-' && !foundSign { sign = -1 }
// - Check for digit: if runes[i] >= '0' && runes[i] <= '9'
// - Extract digit: digit := int(runes[i] - '0')
// - Build number: result = result*10 + digit
// - Mark digit found: foundSign = true
// - Return: result * sign
//
// Number building:
// - Start with 0
// - For each digit: multiply current by 10, add new digit
// - Example: 0 → 1 → 12 → 123
//
// Examples:
//   TrimAtoi("12345") returns 12345
//   TrimAtoi("str123ing45") returns 12345
//   TrimAtoi("sd-x1fa2W3s4") returns -1234

func TrimAtoi(s string) int {
	// Your code here
	return 0
}
