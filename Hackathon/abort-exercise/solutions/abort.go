package solutions

// Abort returns the median of five integers
//
// Algorithm:
// 1. Place all values into a slice
// 2. Sort the slice using bubble sort
// 3. Return the middle element (index 2 for 5 elements)
//
// The median is the middle value when sorted:
//
//	Input: 2, 3, 8, 5, 7
//	Sorted: 2, 3, 5, 7, 8
//	Median: 5 (at index 2)
func Abort(a, b, c, d, e int) int {
	// Create slice with all five values
	numbers := []int{a, b, c, d, e}

	// Bubble sort to arrange in ascending order
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				// Swap if out of order
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	// Return median (middle element at index 2)
	return numbers[len(numbers)/2]
}
