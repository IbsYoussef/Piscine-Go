# advancedsortwordarr

**Status:** Bonus

## Problem Statement

Write a function `AdvancedSortWordArr` that sorts a slice of `string`, based on the function `f` passed in parameter.

## Expected Function Signature

```go
func AdvancedSortWordArr(a []string, f func(a, b string) int) {

}
```

## Expected Output

```console
$ go run .
[1 2 3 A B C a b c]
```

## Requirements

- Create a file named `advancedsortwordarr.go`
- Define package as `package piscine`
- Implement the `AdvancedSortWordArr` function
- Use the comparison function `f` to determine sort order
- Sort strings in-place based on the comparison function

## Submission Structure

```
advancedsortwordarr.go
```

## How to Work on This

1. Write your solution in `student/advancedsortwordarr.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/advancedsortwordarr.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.AdvancedSortWordArr(result, piscine.Compare)

	fmt.Println(result)
}
```