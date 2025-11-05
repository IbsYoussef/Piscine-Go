package student

// TODO: Write your code here
// Implement the StrLen function
//
// Hints:
// - Function signature: func StrLen(s string) int
// - s is the string to count
// - Initialize a counter: count := 0
// - Use for range: for range s { count++ }
// - Return the count
//
// Example:
//   StrLen("Hello World!") returns 12

func StrLen(s string) int {
	// Your code here
	count := 0
	for range s {
		count++
	}
	return count
}
