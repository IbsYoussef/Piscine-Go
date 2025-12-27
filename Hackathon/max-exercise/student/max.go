package student

// TODO: Implement Max function
// Return the maximum value in a slice of integers
//
// Algorithm:
// 1. Check if slice is empty → return 0
// 2. Initialize max to FIRST ELEMENT (not 0!)
//    Why? Consider negative numbers: [-5, -10, -3]
//    If max = 0, it would return 0 instead of -3
// 3. Loop through slice
// 4. If current element > max, update max
// 5. Return max
//
// Common mistake:
// ❌ max := 0  // Fails for negative numbers!
// ✅ max := a[0]  // Works for all cases!

func Max(a []int) int {
	// Your code here
	return 0
}
