# map

**Status:** Required

## Problem Statement

Write a function `Map` that, for an `int` slice, applies a function of this type `func(int) bool` on each element of that slice and returns a slice of all the return values.

## Expected Function Signature

```go
func Map(f func(int) bool, a []int) []bool {

}
```

## Expected Output

```console
$ go run .
[false true true false true false]
```

## Requirements

- Create a file named `map.go`
- Define package as `package piscine`
- Implement the `Map` function
- Apply the function to each element
- Return a new slice with all the boolean results

## Submission Structure

```
map.go
```

## How to Work on This

1. Write your solution in `student/map.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/map.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(piscine.IsPrime, a)
	fmt.Println(result)
}
```