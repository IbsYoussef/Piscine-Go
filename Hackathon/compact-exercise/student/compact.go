package student

// TODO: Implement Compact function
// Remove empty strings and return count of non-empty elements
//
// Algorithm:
// 1. Dereference pointer: slice := *ptr
// 2. Use index variable to track write position
// 3. Loop through slice:
//    - If element is non-empty, write it at index position
//    - Increment index
// 4. Reslice to trim: *ptr = slice[:index]
// 5. Return count: len(*ptr)
//
// Example trace:
//   Start:  ["a", "", "b", "", "c", ""]  index=0
//   After:  ["a", "b", "c", "", "c", ""]  index=3
//   Trim:   ["a", "b", "c"]
//   Return: 3

func Compact(ptr *[]string) int {
	// Your code here
	return 0
}
