# recursivepower

**Status:** Required

## Problem Statement

Write a **recursive** function that returns the value of `nb` to the power of `power`.

Negative powers will return `0`. Overflows do **not** have to be dealt with.

`for` is **forbidden** for this exercise.

## Expected Function Signature

```go
func RecursivePower(nb int, power int) int
```

## Expected Output

```console
64
```

## Requirements

- Create a file named `recursivepower.go`
- Define package as `package piscine`
- Implement the `RecursivePower` function
- Use recursion, **NOT loops** (`for` is forbidden)
- Negative power returns 0
- Any number to power 0 is 1
- Calculate: nb^power recursively

## Submission Structure

```
recursivepower.go
```

## How to Work on This

1. Write your solution in `student/recursivepower.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/recursivepower.go` if you need help

## Hints

- Base case 1: power < 0 → return 0
- Base case 2: power == 0 → return 1
- Base case 3: power == 1 → return nb
- Recursive case: nb × RecursivePower(nb, power-1)
- Each recursive call reduces power by 1

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RecursivePower(4, 3))
}
```

Output:

```
64
```

## Key Concepts

- **Recursion**: Function calling itself
- **Base cases**: Stop conditions
- **Exponentiation**: Repeated multiplication
- **Power reduction**: Each call reduces power by 1

## Recursion Visualization

```
RecursivePower(4, 3)
→ 4 × RecursivePower(4, 2)
  → 4 × RecursivePower(4, 1)
    → 4 (base case)
  → 4 × 4 = 16
→ 4 × 16 = 64
```
