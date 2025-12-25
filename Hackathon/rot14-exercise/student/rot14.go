package student

// TODO: Implement Rot14 function
// Rotate each letter by 14 positions in the alphabet
//
// Algorithm:
// 1. Loop through each character in string
// 2. Check if it's a letter (a-z or A-Z)
// 3. For letters:
//    - Convert to position: r - 'a' (or r - 'A' for uppercase)
//    - Add rotation: + 14
//    - Wrap around: % 26
//    - Convert back: + 'a' (or + 'A' for uppercase)
// 4. For non-letters: keep unchanged
//
// Formula:
//   Lowercase: (r - 'a' + 14) % 26 + 'a'
//   Uppercase: (r - 'A' + 14) % 26 + 'A'
//
// Examples:
//   'a' → 'o' (position 0 + 14 = 14)
//   'z' → 'n' (position 25 + 14 = 39, 39 % 26 = 13)
//   'H' → 'V' (position 7 + 14 = 21)

func Rot14(s string) string {
	// Your code here
	return ""
}
