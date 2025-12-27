package student

// TODO: Implement Unmatch function
// Return the element that doesn't have a pair
//
// Key insight: A number is unpaired if it appears an ODD number of times
//
// Algorithm:
// 1. Create a map to count frequencies: freq := make(map[int]int)
// 2. Count each number: freq[num]++
// 3. Find number with ODD count: count%2 != 0
// 4. Return that number, or -1 if all counts are even
//
// Common mistake:
// ❌ if count == 1  // Only checks for count of 1
// ✅ if count%2 != 0  // Checks for ANY odd count (1, 3, 5, etc.)
//
// Examples:
//   [1, 2, 3, 1, 2, 3, 4] → 4 (appears 1 time)
//   [1, 1, 1] → 1 (appears 3 times, odd!)
//   [5, 5, 7, 7] → -1 (all even counts)

func Unmatch(a []int) int {
	// Your code here
	return -1
}
