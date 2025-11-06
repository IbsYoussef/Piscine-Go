# iterativefactorial

**Status:** Required

## Problem Statement

Write an **iterative** function that returns the factorial of the `int` passed as parameter.

Errors (non possible values or overflows) will return `0`.

## Expected Function Signature

```go
func IterativeFactorial(nb int) int
```

## Expected Output

```console
24
```

## Requirements

- Create a file named `iterativefactorial.go`
- Define package as `package piscine`
- Implement the `IterativeFactorial` function
- Use iteration (loops), **not recursion**
- Handle negative numbers (return 0)
- Handle overflow (return 0)
- Calculate factorial: n! = n × (n-1) × (n-2) × ... × 1

## Submission Structure

```
iterativefactorial.go
```

## How to Work on This

1. Write your solution in `student/iterativefactorial.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/iterativefactorial.go` if you need help

## Hints

- Factorial of 0 is 1
- Factorial of 1 is 1
- Negative numbers return 0
- Use a loop to multiply: result \*= i
- Check for overflow before multiplying
- Overflow check: `result > (1<<63-1)/i`

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.IterativeFactorial(arg))
}
```

Output:

```
24
```

## Key Concepts

- **Factorial**: Product of all positive integers up to n
- **Iterative approach**: Using loops instead of recursion
- **Overflow detection**: Checking before multiplication
- **Edge cases**: 0!, 1!, negative numbers

## Mathematical Background

- 0! = 1 (by definition)
- 1! = 1
- 4! = 4 × 3 × 2 × 1 = 24
- 5! = 5 × 4 × 3 × 2 × 1 = 120
