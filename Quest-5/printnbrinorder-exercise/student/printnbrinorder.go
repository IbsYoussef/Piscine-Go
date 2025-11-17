package student

// TODO: Write your code here
// Implement the PrintNbrInOrder function
//
// Hints:
// - Function signature: func PrintNbrInOrder(n int)
// - Handle edge case: if n == 0 { z01.PrintRune('0'); return }
// - Extract digits:
//   digits := []int{}
//   for n != 0 {
//       digit := n % 10
//       digits = append(digits, digit)
//       n = n / 10
//   }
// - Sort digits (bubble sort):
//   for i := 0; i < len(digits)-1; i++ {
//       for j := 0; j < len(digits)-1-i; j++ {
//           if digits[j] > digits[j+1] {
//               digits[j], digits[j+1] = digits[j+1], digits[j]
//           }
//       }
//   }
// - Print digits:
//   for i := 0; i < len(digits); i++ {
//       z01.PrintRune(rune(digits[i]) + '0')
//   }
//
// Examples:
//   PrintNbrInOrder(321) prints "123"
//   PrintNbrInOrder(0) prints "0"
//   PrintNbrInOrder(54321) prints "12345"

func PrintNbrInOrder(n int) {
	// Your code here
}
