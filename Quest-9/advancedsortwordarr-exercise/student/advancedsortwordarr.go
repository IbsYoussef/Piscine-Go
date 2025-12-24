package student

// TODO: Implement AdvancedSortWordArr function
// Sort a string slice using a custom comparison function
//
// This is similar to SortWordArr, but instead of comparing directly,
// you use the function f to compare elements
//
// Algorithm:
// 1. Use bubble sort (same as SortWordArr)
// 2. Instead of: if a[j] > a[j+1]
//    Use: if f(a[j], a[j+1]) > 0
//
// The comparison function f returns:
//   positive if first > second
//   0 if equal
//   negative if first < second
//
// Hints:
// - Same nested loop structure as SortWordArr
// - Replace direct comparison with function call

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	// Your code here
}
