# foreach

**Status:** Required

## Problem Statement

Write a function `ForEach` that, for an `int` slice, applies a function on each element of that slice.

## Expected Function Signature

```go
func ForEach(f func(int), a []int) {

}
```

## Expected Output

```console
$ go run .
123456
```

## Requirements

- Create a file named `foreach.go`
- Define package as `package piscine`
- Implement the `ForEach` function
- Apply the given function to each element in the slice

## Submission Structure

```
foreach.go
```

## How to Work on This

1. Write your solution in `student/foreach.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/foreach.go` if you need help

## Usage Example

```go
package main

import "piscine"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(piscine.PrintNbr, a)
}
```
