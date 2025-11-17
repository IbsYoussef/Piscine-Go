# printnbrinorder

**Status:** Required

## Problem Statement

Write a function which prints the digits of an `int` passed in parameter in ascending order.

All possible values of type `int` have to go through, excluding negative numbers.
Conversion to `int64` is not allowed.

## Expected Function Signature
```go
func PrintNbrInOrder(n int)
```

## Expected Output
```console
1230123
```

## Requirements

- Create a file named `printnbrinorder.go`
- Define package as `package piscine`
- Import `"github.com/01-edu/z01"`
- Implement the `PrintNbrInOrder` function
- Extract digits from the number
- Sort digits in ascending order
- Print the sorted digits
- Handle 0 as special case (print '0')

## Submission Structure
```
printnbrinorder.go
```

## How to Work on This

1. Write your solution in `student/printnbrinorder.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printnbrinorder.go` if you need help

## Hints

- Handle special case: if n == 0, print '0' and return
- Extract digits using modulo and division: n % 10, n / 10
- Store digits in a slice
- Sort the slice (use bubble sort or any sorting algorithm)
- Print each digit using z01.PrintRune
- Convert digit to rune: rune(digit) + '0'

## Example
```go
package main

import "piscine"

func main() {
	piscine.PrintNbrInOrder(321)  // prints: 123
	piscine.PrintNbrInOrder(0)    // prints: 0
	piscine.PrintNbrInOrder(321)  // prints: 123
}
```

Output:
```
1230123
```

## Key Concepts

- **Digit extraction**: Using modulo and division
- **Sorting**: Arranging digits in order
- **Bubble sort**: Simple sorting algorithm
- **Rune conversion**: Converting int to printable character

## Algorithm Steps

1. Handle edge case: n == 0
2. Extract digits: use n % 10 and n / 10
3. Store digits in slice
4. Sort digits (ascending)
5. Print each sorted digit

## Sorting Example

Input: 321
- Extract digits: [1, 2, 3] (reversed from extraction)
- Sort: [1, 2, 3]
- Print: 123

## Notions

- [01-edu/z01](https://github.com/01-edu/z01)
- [rune-literals](https://golang.org/ref/spec#Rune_literals)