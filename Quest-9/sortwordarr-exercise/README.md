# sortwordarr

**Status:** Optional

## Problem Statement

Write a function `SortWordArr` that sorts by `ascii` (in ascending order) a `string` slice.

## Expected Function Signature
```go
func SortWordArr(a []string) {

}
```

## Expected Output
```console
$ go run .
[1 2 3 A B C a b c]
```

## Requirements

- Create a file named `sortwordarr.go`
- Define package as `package piscine`
- Implement the `SortWordArr` function
- Sort strings in ascending ASCII order
- Modify the slice in-place (no return value)

## Submission Structure
```
sortwordarr.go
```

## How to Work on This

1. Write your solution in `student/sortwordarr.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/sortwordarr.go` if you need help

## Usage Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.SortWordArr(result)

	fmt.Println(result)
}
```