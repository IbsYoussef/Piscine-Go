# activebits

## Problem Statement

Write a function, `ActiveBits`, that returns the number of active `bits` (bits with the value 1) in the binary representation of an integer number.

## Expected Function Signature

```go
func ActiveBits(n int) int {

}
```

## Expected Output

```console
$ go run .
3
```

## Requirements

- Create a file named `activebits.go`
- Define package as `package piscine`
- Implement the `ActiveBits` function
- Count the number of 1-bits in binary representation
- Do not use standard library packages

## Submission Structure

```
activebits.go
```

## How to Work on This

1. Write your solution in `student/activebits.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/activebits.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ActiveBits(7))
}
```

## Examples

```
7 in binary:  0111 → 3 active bits
10 in binary: 1010 → 2 active bits
15 in binary: 1111 → 4 active bits
```
