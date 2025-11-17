package student

// TODO: Write your code here
// Implement the Join function
//
// Hints:
// - Function signature: func Join(strs []string, sep string) string
// - Handle edge case: if len(strs) == 0 { return "" }
// - Initialize result: result := ""
// - Loop with index: for i, v := range strs
// - Add separator before element (except first):
//   if i > 0 {
//       result += sep
//   }
// - Add element: result += v
// - Return result
//
// Logic:
// - First element: no separator before it
// - Subsequent elements: add separator, then element
//
// Examples:
//   Join([]string{"Hello!", " How", " are", " you?"}, ":") returns "Hello!: How: are: you?"
//   Join([]string{"Go", "Lang"}, "-") returns "Go-Lang"
//   Join([]string{}, ",") returns ""

func Join(strs []string, sep string) string {
	// Your code here
	return ""
}
