package solutions

// Itoa converts an integer to a string
// This is the reverse of Atoi
//
// Algorithm:
// 1. Handle sign: save "-" if negative, make number positive
// 2. Handle zero edge case
// 3. Extract digits using modulo and division
// 4. Convert each digit to character: '0' + digit
// 5. Build string from right to left
// 6. Prepend sign if needed
//
// Example for n = -1234:
//   Step 1: sign = "-", n = 1234
//   Step 2: Extract digits
//     1234 % 10 = 4 → '4', n = 123
//     123 % 10 = 3 → '3', n = 12
//     12 % 10 = 2 → '2', n = 1
//     1 % 10 = 1 → '1', n = 0
//   Result: "1234"
//   Step 3: Prepend sign → "-1234"

func Itoa(n int) string {
	result := ""
	sign := ""

	// Handle negative numbers
	if n < 0 {
		sign = "-"
		n = -n
	}

	// Handle zero
	if n == 0 {
		return "0"
	}

	// Extract digits and build string
	for n > 0 {
		digit := n % 10                           // Get last digit
		result = string(rune('0'+digit)) + result // Convert and prepend
		n /= 10                                   // Remove last digit
	}

	return sign + result
}
