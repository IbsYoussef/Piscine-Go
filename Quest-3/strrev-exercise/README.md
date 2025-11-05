# strrev

**Status:** Required

## Problem Statement

Write a function that reverses a `string`. This function will **return** the reversed `string`.

## Expected Function Signature

```go
func StrRev(s string) string
```

## Expected Output

```console
!dlroW olleH
```

## Requirements

- Create a file named `strrev.go`
- Define package as `package piscine`
- Implement the `StrRev` function
- Reverse the order of characters in the string
- Return the reversed string

## Submission Structure

```
strrev.go
```

## How to Work on This

1. Write your solution in `student/strrev.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/strrev.go` if you need help

## Hints

- Create an empty result string
- Iterate through the input string
- Prepend each character to the result (add to the front)
- Alternative: `result = string(c) + result`
- Return the final reversed string

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
```

Output:
```
!dlroW olleH
```

## Key Concepts

- **String reversal**: Common algorithm
- **String concatenation**: Building new strings
- **Rune to string conversion**: `string(rune)`
- **Order manipulation**: Reversing sequences