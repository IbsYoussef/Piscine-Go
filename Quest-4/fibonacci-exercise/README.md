# fibonacci

**Status:** Required

## Problem Statement

Write a **recursive** function that returns the value at the position `index` in the fibonacci sequence.

The first value is at index `0`.

The sequence starts this way: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...

A negative index will return `-1`.

`for` is **forbidden** for this exercise.

## Expected Function Signature

```go
func Fibonacci(index int) int
```

## Expected Output

```console
3
```

## Requirements

- Create a file named `fibonacci.go`
- Define package as `package piscine`
- Implement the `Fibonacci` function
- Use recursion, **NOT loops** (`for` is forbidden)
- Negative index returns -1
- Return value at position index in fibonacci sequence

## Submission Structure

```
fibonacci.go
```

## How to Work on This

1. Write your solution in `student/fibonacci.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/fibonacci.go` if you need help

## Hints

- Base case 1: index < 0 → return -1
- Base case 2: index == 0 → return 0
- Base case 3: index == 1 → return 1
- Recursive case: Fibonacci(index-1) + Fibonacci(index-2)
- Each call branches into two recursive calls

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg1 := 4
	fmt.Println(piscine.Fibonacci(arg1))
}
```

Output:

```
3
```

## Key Concepts

- **Fibonacci sequence**: Each number is sum of previous two
- **Recursion**: Function calling itself
- **Multiple recursive calls**: Branches into two calls
- **Base cases**: Stop conditions

## Fibonacci Sequence

```
Index: 0  1  2  3  4  5  6  7  8  9  10
Value: 0  1  1  2  3  5  8  13 21 34 55
```

## Recursion Visualization

```
Fibonacci(4)
→ Fibonacci(3) + Fibonacci(2)
  → (Fibonacci(2) + Fibonacci(1)) + (Fibonacci(1) + Fibonacci(0))
    → ((Fibonacci(1) + Fibonacci(0)) + 1) + (1 + 0)
      → ((1 + 0) + 1) + (1 + 0)
      → (1 + 1) + 1
      → 2 + 1
      → 3
```
