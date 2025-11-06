# sqrt

**Status:** Optional

## Problem Statement

Write a function that returns the square root of the `int` passed as parameter, if that square root is a whole number. Otherwise it returns `0`.

## Expected Function Signature

```go
func Sqrt(nb int) int
```

## Expected Output

```console
2
0
```

## Requirements

- Create a file named `sqrt.go`
- Define package as `package piscine`
- Implement the `Sqrt` function
- Return the square root if it's a whole number
- Return 0 if square root is not a whole number
- Return 0 for negative numbers

## Submission Structure

```
sqrt.go
```

## How to Work on This

1. Write your solution in `student/sqrt.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/sqrt.go` if you need help

## Hints

- Negative numbers have no real square root → return 0
- 0 squared is 0 → return 0
- Try each number from 1 to nb
- Check if i × i == nb
- If yes, return i
- If loop completes without finding, return 0

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Sqrt(4))
	fmt.Println(piscine.Sqrt(3))
}
```

Output:

```
2
0
```

## Key Concepts

- **Square root**: Number that when multiplied by itself gives the original
- **Whole numbers**: Integers only
- **Iteration**: Testing values sequentially
- **Perfect squares**: 1, 4, 9, 16, 25, 36, 49, 64, 81, 100...

## Mathematical Background

- √4 = 2 (because 2 × 2 = 4)
- √9 = 3 (because 3 × 3 = 9)
- √3 = not a whole number
- √16 = 4 (because 4 × 4 = 16)
