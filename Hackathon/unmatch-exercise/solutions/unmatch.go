package solutions

// Unmatch returns the element that doesn't have a pair
// Returns -1 if all elements are paired
//
// A number is "unpaired" if it appears an ODD number of times
// A number is "paired" if it appears an EVEN number of times
//
// Algorithm:
// 1. Count frequency of each number using a map
// 2. Find the number with odd frequency
// 3. Return that number, or -1 if all are even
//
// Examples:
//   [1, 2, 3, 1, 2, 3, 4]
//   Frequencies: {1:2, 2:2, 3:2, 4:1}
//   4 has odd count (1) → return 4
//
//   [1, 1, 1]
//   Frequencies: {1:3}
//   1 has odd count (3) → return 1
//
//   [5, 5, 7, 7]
//   Frequencies: {5:2, 7:2}
//   All even counts → return -1

func Unmatch(a []int) int {
	// Create frequency map
	freq := make(map[int]int)

	// Count occurrences of each number
	for _, num := range a {
		freq[num]++
	}

	// Find number with odd count
	for num, count := range freq {
		if count%2 != 0 {
			return num
		}
	}

	// All numbers are paired
	return -1
}
