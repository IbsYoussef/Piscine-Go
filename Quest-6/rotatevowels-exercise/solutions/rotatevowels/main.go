package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Program Breakdown:
//
// Goal: Mirror vowel positions across all arguments
//
// Approach:
//   - Check for no arguments → print newline
//   - Collect all vowels from all arguments
//   - Reverse the vowel list
//   - Rebuild arguments by replacing vowels with reversed ones
//   - Print modified arguments with spaces between them
//
// Key Concepts:
//   - Vowel detection: Using map for O(1) lookup
//   - Collection: Gathering vowels into slice
//   - Reversal: Swap elements from ends toward middle
//   - Replacement: Track position in reversed vowels with index
//   - Immutability: Convert strings to rune slices for modification
//
// Algorithm Steps:
//   1. Handle edge case: no arguments
//   2. Create vowel map (a,e,i,o,u uppercase and lowercase)
//   3. Collect all vowels from all arguments
//   4. Reverse the vowel slice (swap from ends)
//   5. Loop through arguments again:
//      - Convert to rune slice
//      - Replace each vowel with next from reversed list
//      - Print modified argument
//   6. Print spaces between arguments
//   7. Print final newline
//
// Examples:
//   "Hello World" → vowels: [e,o,o] → reversed: [o,o,e] → "Hollo Werld"
//   "aEi" "Ou" → vowels: [a,E,i,O,u] → reversed: [u,O,i,E,a] → "uOi Ea"

// Status: Bonus

func main() {
	args := os.Args[1:]

	// Handle no arguments
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Create vowel map for O(1) lookup
	vowels := map[rune]bool{
		'a': true, 'A': true,
		'e': true, 'E': true,
		'i': true, 'I': true,
		'o': true, 'O': true,
		'u': true, 'U': true,
	}

	// Collect all vowels from all arguments
	store := []rune{}
	for _, arg := range args {
		for _, r := range arg {
			if vowels[r] {
				store = append(store, r)
			}
		}
	}

	// Reverse the vowel slice
	// Swap elements from ends toward middle
	for i := 0; i < len(store)/2; i++ {
		store[i], store[len(store)-1-i] = store[len(store)-1-i], store[i]
	}

	// Rebuild arguments with reversed vowels
	vowelIndex := 0
	for argIdx, arg := range args {
		result := []rune(arg) // Convert to rune slice for modification

		// Replace each vowel with next from reversed list
		for i, r := range result {
			if vowels[r] {
				result[i] = store[vowelIndex]
				vowelIndex++
			}
		}

		// Print the modified argument
		for _, char := range result {
			z01.PrintRune(char)
		}

		// Print space between arguments (not after last)
		if argIdx < len(args)-1 {
			z01.PrintRune(' ')
		}
	}

	// Print final newline
	z01.PrintRune('\n')
}
