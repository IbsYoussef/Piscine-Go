package student

// TODO: Implement IsSorted function
// Check if array is sorted (ascending or descending)
//
// Algorithm:
// 1. Handle edge case: arrays with < 2 elements are sorted
// 2. Determine direction: find first non-equal pair
//    - If a[i] > a[i+1] → descending
//    - If a[i] < a[i+1] → ascending
//    - If a[i] == a[i+1] → keep looking
// 3. Verify all pairs follow that direction
//    - If ascending: no pair should have a[i] > a[i+1]
//    - If descending: no pair should have a[i] < a[i+1]
//
// Hints:
// - Use two loops: one to find direction, one to verify
// - Use f() to compare pairs: f(a[i], a[i+1])

func IsSorted(f func(a, b int) int, a []int) bool {
	// Your code here
	return false
}
