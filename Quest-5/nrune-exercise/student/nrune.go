package student

// TODO: Write your code here
// Implement the NRune function
//
// Hints:
// - Function signature: func NRune(s string, n int) rune
// - Convert string to rune slice: runes := []rune(s)
// - Validate n: if n <= 0 || n > len(runes) { return rune(0) }
// - Convert 1-indexed to 0-indexed: index = n - 1
// - Return: runes[n-1]
//
// Important:
// - n is 1-indexed (first rune is n=1)
// - Arrays are 0-indexed (first element is index 0)
// - Return rune(0) for invalid positions
//
// Examples:
//   NRune("Hello!", 3) returns 'l'
//   NRune("Salut!", 2) returns 'a'
//   NRune("Bye!", -1) returns 0
//   NRune("Bye!", 5) returns 0

func NRune(s string, n int) rune {
	// Your code here
	runes := []rune(s)

	// Validate n: must be positive and within bounds
	if n <= 0 || n > len(runes) {
		return rune(0)
	}

	// Return nth rune (convert 1-indexed to 0-indexed)
	return runes[n-1]
}
