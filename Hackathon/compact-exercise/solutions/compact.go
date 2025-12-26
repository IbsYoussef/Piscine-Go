package solutions

// Compact removes empty strings from a slice and returns the count
// of non-empty elements
//
// Algorithm (two-pointer technique):
// 1. Use index to track where to write non-empty values
// 2. Loop through all elements (read pointer)
// 3. When finding non-empty value, write it at index position
// 4. Increment index after writing
// 5. Reslice to only include [0:index] elements
//
// Example:
//
//	Input:  ["a", "", "b", "", "c", ""]
//	After moving: ["a", "b", "c", "", "c", ""]
//	After reslice: ["a", "b", "c"]
//	Return: 3

func Compact(ptr *[]string) int {
	// Dereference pointer to get actual slice
	slice := *ptr

	// Track where to write non-empty values
	index := 0

	// Loop through all elements
	for _, v := range slice {
		if v != "" {
			// Move non-empty value to front
			slice[index] = v
			index++
		}
	}

	// Trim slice to only include non-empty values
	*ptr = slice[:index]

	// Return count of non-empty elements
	return len(*ptr)
}
