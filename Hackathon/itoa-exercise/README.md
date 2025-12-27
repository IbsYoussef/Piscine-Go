# itoa

## Problem Statement

Write a function that simulates the behavior of the `Itoa` function in Go. `Itoa` transforms a number represented as an `int` in a number represented as a `string`.

- For this exercise the handling of the signs + or - **does have** to be taken into account.

## Expected Function Signature

```go
func Itoa(n int) string {

}
```

## Expected Output

```console
$ go run .
12345
0
-1234
987654321
```

## Requirements

- Create a file named `itoa.go`
- Define package as `package piscine`
- Implement the `Itoa` function
- Convert integers to strings
- Handle negative numbers
- Handle zero

## Submission Structure

```
itoa.go
```

## How to Work on This

1. Write your solution in `student/itoa.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/itoa.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Itoa(12345))
	fmt.Println(piscine.Itoa(0))
	fmt.Println(piscine.Itoa(-1234))
	fmt.Println(piscine.Itoa(987654321))
}
```

## Notions

- [strconv/Itoa](https://pkg.go.dev/strconv#Itoa)
