# findnextprime

**Status:** Optional

## Problem Statement

Write a function that returns the first prime number that is equal or superior to the `int` passed as parameter.

The function must be optimized in order to avoid time-outs with the tester.

(We consider that only positive numbers can be prime numbers)

## Expected Function Signature

```go
func FindNextPrime(nb int) int
```

## Expected Output

```console
5
5
```

## Requirements

- Create a file named `findnextprime.go`
- Define package as `package piscine`
- Implement the `FindNextPrime` function
- Return first prime >= nb
- Must be optimized
- Use IsPrime function to check primality

## Submission Structure

```
findnextprime.go
```

## How to Work on This

1. Write your solution in `student/findnextprime.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/findnextprime.go` if you need help

## Hints

- Handle edge case: nb ≤ 1 → return 2 (first prime)
- Check if nb is already prime using IsPrime
- If yes, return nb
- If no, increment nb and check again
- Use infinite loop: for { ... }
- Break when prime found

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FindNextPrime(5))
	fmt.Println(piscine.FindNextPrime(4))
}
```

Output:

```
5
5
```

## Key Concepts

- **Next prime**: First prime >= given number
- **IsPrime function**: Reuse existing primality test
- **Infinite loop**: Continue until prime found
- **Edge cases**: Numbers ≤ 1

## Examples

- FindNextPrime(5) = 5 (5 is prime)
- FindNextPrime(4) = 5 (4 not prime, next is 5)
- FindNextPrime(14) = 17 (14, 15, 16 not prime)
- FindNextPrime(1) = 2 (first prime)
