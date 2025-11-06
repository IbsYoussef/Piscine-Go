# recursivefactorial

**Status:** Required

## Problem Statement

Write a **recursive** function that returns the factorial of the `int` passed as parameter.

Errors (non possible values or overflows) will return `0`.

`for` is **forbidden** for this exercise.

## Expected Function Signature

```go
func RecursiveFactorial(nb int) int
```

## Expected Output

```console
24
```

## Requirements

- Create a file named `recursivefactorial.go`
- Define package as `package piscine`
- Implement the `RecursiveFactorial` function
- Use recursion, **NOT loops** (`for` is forbidden)
- Handle negative numbers (return 0)
- Handle overflow (return 0)
- Calculate factorial recursively

## Submission Structure

```
recursivefactorial.go
```

## How to Work on This

1. Write your solution in `student/recursivefactorial.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/recursivefactorial.go` if you need help

## Hints

- Base cases: nb < 0 → 0, nb > 20 → 0 (overflow)
- Base cases: nb == 0 or nb == 1 → 1
- Recursive case: nb × RecursiveFactorial(nb-1)
- No loops allowed - must use recursion
- Each call reduces the problem size by 1

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg))
}
```

Output:

```
24
```

## Key Concepts

- **Recursion**: Function calling itself
- **Base case**: Condition to stop recursion
- **Recursive case**: Breaking problem into smaller parts
- **Call stack**: Each recursive call adds to the stack

## Recursion Visualization

```
RecursiveFactorial(4)
→ 4 × RecursiveFactorial(3)
  → 3 × RecursiveFactorial(2)
    → 2 × RecursiveFactorial(1)
      → 1 (base case)
    → 2 × 1 = 2
  → 3 × 2 = 6
→ 4 × 6 = 24
```
