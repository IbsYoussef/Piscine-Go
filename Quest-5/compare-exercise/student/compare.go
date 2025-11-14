package student

// TODO: Write your code here
// Implement the Compare function
//
// Hints:
// - Function signature: func Compare(a, b string) int
// - Handle edge cases: empty strings
// - Convert to runes: aRunes := []rune(a), bRunes := []rune(b)
// - Loop and compare: for i := 0; i < len(aRunes) && i < len(bRunes); i++
// - If aRunes[i] > bRunes[i] → return 1
// - If aRunes[i] < bRunes[i] → return -1
// - After loop, compare lengths
// - If equal, return 0
//
// Return values:
// - 0: strings are equal
// - -1: a comes before b (a < b)
// - 1: a comes after b (a > b)
//
// Examples:
//   Compare("Hello!", "Hello!") returns 0
//   Compare("Salut!", "lut!") returns -1
//   Compare("Ola!", "Ol") returns 1

func Compare(a, b string) int {
	// Your code here
	return 0
}
