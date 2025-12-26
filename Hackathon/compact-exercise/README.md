# compact

## Problem Statement

Write a function `Compact` that takes a pointer to a slice of `string`s as the argument.

This function must:

- Return the number of elements with [non-zero value](https://tour.golang.org/basics/12).
- Compact, i.e., delete the elements with zero-values in the slice.

## Expected Function Signature

```go
func Compact(ptr *[]string) int {

}
```

## Expected Output

```console
$ go run .
a

b

c

Size after compacting: 3
a
b
c
```

## Requirements

- Create a file named `compact.go`
- Define package as `package piscine`
- Implement the `Compact` function
- Remove empty strings from the slice
- Return count of non-empty elements
- Modify the original slice in-place

## Submission Structure

```
compact.go
```

## How to Work on This

1. Write your solution in `student/compact.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/compact.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", piscine.Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
```
