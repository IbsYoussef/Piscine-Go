# max

## Problem Statement

Write a function `Max` that will return the maximum value in a slice of integers. If the slice is empty it will return 0.

## Expected Function Signature

```go
func Max(a []int) int {

}
```

## Expected Output

```console
$ go run .
123
```

## Requirements

- Create a file named `max.go`
- Define package as `package piscine`
- Implement the `Max` function
- Return the maximum value in the slice
- Return 0 if slice is empty

## Submission Structure

```
max.go
```

## How to Work on This

1. Write your solution in `student/max.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/max.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := piscine.Max(a)
	fmt.Println(max)
}
```

## Edge Cases

- Empty slice: return 0
- Negative numbers: return the largest (least negative)
- Single element: return that element
