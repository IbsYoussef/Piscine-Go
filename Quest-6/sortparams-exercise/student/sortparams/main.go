package main

// TODO: Write your code here
// Print command-line arguments in ASCII sorted order
//
// Hints:
// - Get arguments: os.Args[1:]
// - Sort the arguments using bubble sort
// - Bubble sort structure:
//   for i := 0; i < len(args)-1; i++ {
//       for j := 0; j < len(args)-i-1; j++ {
//           if args[j] > args[j+1] {
//               // Swap
//               args[j], args[j+1] = args[j+1], args[j]
//           }
//       }
//   }
// - After sorting, print each argument (same as printparams)
//
// ASCII order reminder:
// - Numbers (0-9) come first
// - Uppercase letters (A-Z) come next
// - Lowercase letters (a-z) come last
//
// Example:
//   ./sortparams 1 a 2 A 3 b 4 C
//   Output:
//   1
//   2
//   3
//   4
//   A
//   C
//   a
//   b

func main() {
	// Your code here
}
