package solutions

// Function Breakdown:
//
// ToLower(s string) string:
//   - Converts all uppercase letters to lowercase
//   - Leaves other characters unchanged
//   - Returns the modified string
//
// Approach:
//   - Convert string to rune slice (for modification)
//   - Loop through each rune
//   - If rune is uppercase, convert to lowercase
//   - Convert rune slice back to string
//   - Return the result
//
// Key Concepts:
//   - ASCII arithmetic: uppercase + 32 = lowercase
//   - Rune slice: Allows in-place modification
//   - Case conversion: Only affects letters A-Z
//   - Immutability: Strings can't be modified, so use runes
//
// Conversion logic:
//   - Check if uppercase: rune >= 'A' && rune <= 'Z'
//   - Convert: rune += 32
//   - ASCII 'A' (65) + 32 = 'a' (97)
//   - ASCII 'Z' (90) + 32 = 'z' (122)
//
// Why add 32?
//   - In ASCII table, uppercase and lowercase differ by 32
//   - 'A' = 65, 'a' = 97, difference = 32
//   - Adding 32 converts uppercase to lowercase
//
// Examples:
//   ToLower("Hello! How are you?") = "hello! how are you?"
//   ToLower("ABC123") = "abc123"
//   ToLower("already lower") = "already lower"
//   ToLower("TEST") = "test"
//   ToLower("") = ""

// Status: Required

func ToLower(s string) string {
	// Convert string to rune slice (for modification)
	runes := []rune(s)

	// Loop through each rune
	for i := 0; i < len(runes); i++ {
		// Check if uppercase letter
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			// Convert to lowercase by adding 32
			runes[i] += 32
		}
	}

	// Convert rune slice back to string
	return string(runes)
}
