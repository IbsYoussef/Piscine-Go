# strlen

**Status:** Optional

## Problem Statement

Write a function that counts the `runes` of a `string` and that returns that count.

## Expected Function Signature

```go
func StrLen(s string) int
```

## Expected Output

```console
12
```

## Requirements

- Create a file named `strlen.go`
- Define package as `package piscine`
- Implement the `StrLen` function
- Count all runes (characters) in the string
- Return the count as an int

## Submission Structure

```
strlen.go
```

## How to Work on This

1. Write your solution in `student/strlen.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/strlen.go` if you need help

## Hints

- Use a counter variable initialized to 0
- Use `for range` to iterate through the string
- Increment counter for each character
- Return the final count
- You could also use `for range s { count++ }` without capturing values

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}
```

Output:
```
12
```

## Key Concepts

- **String length**: Counting characters
- **Runes**: Unicode code points
- **Iteration**: Using for range
- **Return values**: Returning computed results