# countif

**Status:** Required

## Problem Statement

Write a function `CountIf` that returns the number of elements of a `string` slice for which the `f` function returns `true`.

## Expected Function Signature

```go
func CountIf(f func(string) bool, tab []string) int {

}
```

## Expected Output

```console
$ go run .
0
2
```

## Requirements

- Create a file named `countif.go`
- Define package as `package piscine`
- Implement the `CountIf` function
- Count how many elements pass the test function
- Return the count as an integer

## Submission Structure

```
countif.go
```

## How to Work on This

1. Write your solution in `student/countif.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/countif.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
	answer2 := piscine.CountIf(piscine.IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
```
