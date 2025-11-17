package solutions

// Function Breakdown:
//
// ToUpper(s string) string:
//   - Converts all lowercase letters to uppercase
//   - Leaves other characters unchanged
//   - Returns the modified string
//
// Approach:
//   - Convert string to rune slice (for modification)
//   - Loop through each rune
//   - If rune is lowercase, convert to uppercase
//   - Convert rune slice back to string
//   - Return the result
//
// Key Concepts:
//   - ASCII arithmetic: lowercase - 32 = uppercase
//   - Rune slice: Allows in-place modification
//   - Case conversion: Only affects letters a-z
//   - Immutability: Strings can't be modified, so use runes
//
// Conversion logic:
//   - Check if lowercase: rune >= 'a' && rune <= 'z'
//   - Convert: rune -= 32
//   - ASCII 'a' (97) - 32 = 'A' (65)
//   - ASCII 'z' (122) - 32 = 'Z' (90)
//
// Why subtract 32?
//   - In ASCII table, uppercase and lowercase differ by 32
//   - 'A' = 65, 'a' = 97, difference = 32
//   - Subtracting 32 converts lowercase to uppercase
//
// Examples:
//   ToUpper("Hello! How are you?") = "HELLO! HOW ARE YOU?"
//   ToUpper("abc123") = "ABC123"
//   ToUpper("ALREADY UPPER") = "ALREADY UPPER"
//   ToUpper("test") = "TEST"
//   ToUpper("") = ""

// Status: Required

func ToUpper(s string) string {
	// Convert string to rune slice (for modification)
	runes := []rune(s)

	// Loop through each rune
	for i := 0; i < len(runes); i++ {
		// Check if lowercase letter
		if runes[i] >= 'a' && runes[i] <= 'z' {
			// Convert to uppercase by subtracting 32
			runes[i] -= 32
		}
	}

	// Convert rune slice back to string
	return string(runes)
}
