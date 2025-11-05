package student

// TODO: Write your code here
// Implement the UltimatePointOne function
//
// Hints:
// - Function signature: func UltimatePointOne(n ***int)
// - n is a pointer to a pointer to a pointer to an int
// - To reach the final int, dereference 3 times: ***n
// - Set that value to 1: ***n = 1
//
// Understanding pointers:
//   - *n dereferences once (gives you **int)
//   - **n dereferences twice (gives you *int)
//   - ***n dereferences three times (gives you int)
//
// Example structure:
//   ***n = 1
//
// Example:
//   a := 0
//   b := &a
//   c := &b
//   UltimatePointOne(&c)
//   fmt.Println(a)  // Output: 1

func UltimatePointOne(n ***int) {
	// Your code here
}
