# any

**Status:** Required

## Problem Statement

Write a function `Any` that returns `true`, for a `string` slice:

- if, when that `string` slice is passed through an `f` function, at least one element returns `true`.

## Expected Function Signature

```go
func Any(f func(string) bool, a []string) bool {

}
```

## Expected Output

```console
$ go run .
false
true
```

## Requirements

- Create a file named `any.go`
- Define package as `package piscine`
- Implement the `Any` function
- Return `true` if at least one element passes the test
- Return `false` if no elements pass the test

## Submission Structure

```
any.go
```

## How to Work on This

1. Write your solution in `student/any.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/any.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := piscine.Any(piscine.IsNumeric, a1)
	result2 := piscine.Any(piscine.IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
```
