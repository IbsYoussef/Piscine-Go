package student

// TODO: Write your code here
// Implement the PrintCombN function
//
// Hints:
// - Function signature: func PrintCombN(n int)
// - Validate: return if n <= 0 or n >= 10
// - Use recursion to generate combinations
// - Helper function signature: func helper(current string, start int, n int)
// - Base case: if len(current) == n, print it
// - Check if not max combo before printing ", "
// - Max combo for n: last n digits (e.g., "789" for n=3)
// - Recursive case: for i := start; i < 10; i++
//   - Recurse with current + digit(i) and start = i+1
// - Start recursion with: helper("", 0, n)
// - Print final newline with fmt.Println()
//
// Expected output for n=3:
// 012, 013, 014, ..., 789

func PrintCombN(n int) {
	// Your code here
}
