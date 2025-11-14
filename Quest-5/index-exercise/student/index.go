package student

// TODO: Write your code here
// Implement the Index function
//
// Hints:
// - Function signature: func Index(s string, toFind string) int
// - Edge case: if toFind == "" { return 0 }
// - Edge case: if len(toFind) > len(s) || s == "" { return -1 }
// - Loop: for i := 0; i <= len(s)-len(toFind); i++
// - Slice substring: s[i:i+len(toFind)]
// - Compare: if s[i:i+len(toFind)] == toFind { return i }
// - After loop: return -1
//
// String slicing:
// - s[start:end] extracts substring from start to end (exclusive)
// - Example: "Hello"[1:4] = "ell"
//
// Examples:
//   Index("Hello!", "l") returns 2
//   Index("Salut!", "alu") returns 1
//   Index("Ola!", "hOl") returns -1

func Index(s string, toFind string) int {
	// Your code here
	return 0
}
