# lastrune

**Status:** Required

## Problem Statement

Write a function that returns the last `rune` of a `string`.

## Expected Function Signature

```go
func LastRune(s string) rune
```

## Expected Output

```console
!!!
```

## Requirements

- Create a file named `lastrune.go`
- Define package as `package piscine`
- Implement the `LastRune` function
- Return the last rune (character) of the string

## Submission Structure

```
lastrune.go
```

## How to Work on This

1. Write your solution in `student/lastrune.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/lastrune.go` if you need help

## Hints

- Convert string to rune slice: `[]rune(s)`
- Get length of rune slice: `len(runes)`
- Last element is at index: `len(runes) - 1`
- Remember: arrays/slices are 0-indexed

## Example

```go
package main

import (
	"piscine"
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.LastRune("Hello!"))
	z01.PrintRune(piscine.LastRune("Salut!"))
	z01.PrintRune(piscine.LastRune("Ola!"))
	z01.PrintRune('\n')
}
```

Output:

```
!!!
```

## Key Concepts

- **Runes**: Go's character type (int32)
- **String to rune conversion**: Handling Unicode
- **Array indexing**: Last element = length - 1
- **Length function**: `len()` for slices

## Notions

- [rune-literals](https://golang.org/ref/spec#Rune_literals)
