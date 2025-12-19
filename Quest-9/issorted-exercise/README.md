# issorted

**Status:** Required

## Problem Statement

Write a function `IsSorted()` that returns `true` if the slice of `int` is sorted, otherwise returns `false`.

- The function passed as an argument `func(a, b int)` returns a positive `int` if the first argument is greater than the second argument, it returns `0` if they are equal and it returns a negative `int` otherwise.
- To do your testing you have to write your own `f` function.

## Expected Function Signature

```go
func IsSorted(f func(a, b int) int, a []int) bool {

}
```

## Expected Output

```console
$ go run .
true
false
```

## Requirements

- Create a file named `issorted.go`
- Define package as `package piscine`
- Implement the `IsSorted` function
- Function should detect both ascending and descending order
- Handle arrays with duplicates

## Submission Structure

```
issorted.go
```

## How to Work on This

1. Write your solution in `student/issorted.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/issorted.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
```
