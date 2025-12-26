# abort

## Problem Statement

Write a function that returns the median of five `int` arguments.

## Expected Function Signature

```go
func Abort(a, b, c, d, e int) int {

}
```

## Expected Output

```console
$ go run .
5
```

## Requirements

- Create a file named `abort.go`
- Define package as `package piscine`
- Implement the `Abort` function
- Return the median (middle value) of the five integers

## Submission Structure

```
abort.go
```

## How to Work on This

1. Write your solution in `student/abort.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/abort.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	middle := piscine.Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
```
