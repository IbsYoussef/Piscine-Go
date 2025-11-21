package main

// TODO: Write your code here
// Convert numbers to alphabet letters
//
// Hints:
// - Create a map or use ASCII arithmetic: 1='a', 2='b', etc.
// - Formula: letter = 'a' + (n - 1)
//
// - Check if first arg is "--upper", set flag
// - Process remaining arguments
//
// - For each argument:
//   1. Convert string to int (manually - loop through digits)
//   2. Validate: is it 1-26?
//   3. If valid: convert to letter
//   4. If --upper flag: convert to uppercase (letter - 'a' + 'A')
//   5. If invalid: print space
//
// - Print newline at end
//
// String to int conversion:
//   n := 0
//   for _, r := range arg {
//       if r < '0' || r > '9' { invalid }
//       n = n*10 + int(r-'0')
//   }
//
// Examples:
//   ./nbrconvertalpha 8 5 12 12 15
//   Output: hello
//
//   ./nbrconvertalpha --upper 8 5 25
//   Output: HEY

func main() {
	// Your code here
}
