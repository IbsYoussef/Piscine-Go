# collatzcountdown

## Problem Statement

Write a function, `CollatzCountdown`, that returns the number of steps necessary to reach 1 using the [collatz countdown](https://en.wikipedia.org/wiki/Collatz_conjecture).

- It must return `-1` if `start` is equal to `0` or negative.

## Expected Function Signature

```go
func CollatzCountdown(start int) int {

}
```

## Expected Output

```console
$ go run .
9
```

## Requirements

- Create a file named `collatzcountdown.go`
- Define package as `package piscine`
- Implement the `CollatzCountdown` function
- Apply Collatz rules until reaching 1
- Return -1 for zero or negative input

## Submission Structure

```
collatzcountdown.go
```

## How to Work on This

1. Write your solution in `student/collatzcountdown.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/collatzcountdown.go` if you need help

## Collatz Rules

Starting with a number, apply these rules until reaching 1:

- **If even:** divide by 2
- **If odd:** multiply by 3 and add 1

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	steps := piscine.CollatzCountdown(12)
	fmt.Println(steps)
}
```

## Example Trace

```
Start: 12
12 → 6 → 3 → 10 → 5 → 16 → 8 → 4 → 2 → 1
Steps: 9
```
