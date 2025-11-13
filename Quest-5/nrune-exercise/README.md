# nrune

**Status:** Required

## Problem Statement

Write a function that returns the nth `rune` of a `string`. If not possible, it returns `0`.

## Expected Function Signature

```go
func NRune(s string, n int) rune
```

## Expected Output

```console
la!
```

## Requirements

- Create a file named `nrune.go`
- Define package as `package piscine`
- Implement the `NRune` function
- Return the nth rune (character) of the string
- Return `0` if n is invalid (negative, zero, or exceeds length)
- Note: n is 1-indexed (first rune is n=1)

## Submission Structure

```
nrune.go
```

## How to Work on This

1. Write your solution in `student/nrune.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/nrune.go` if you need help

## Hints

- Convert string to rune slice: `[]rune(s)`
- Check if n is valid: n > 0 and n <= len(runes)
- n is 1-indexed, but arrays are 0-indexed
- So rune at position n is at index n-1
- Return `rune(0)` for invalid n

## Example

```go
package main

import (
	"piscine"
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.NRune("Hello!", 3))  // 'l'
	z01.PrintRune(piscine.NRune("Salut!", 2))  // 'a'
	z01.PrintRune(piscine.NRune("Bye!", -1))   // 0 (nothing printed)
	z01.PrintRune(piscine.NRune("Bye!", 5))    // 0 (nothing printed)
	z01.PrintRune(piscine.NRune("Ola!", 4))    // '!'
	z01.PrintRune('\n')
}
```

Output:

```
la!
```

## Key Concepts

- **1-indexed vs 0-indexed**: n starts at 1, arrays start at 0
- **Validation**: Check bounds before accessing
- **Rune conversion**: Proper Unicode handling
- **Return 0**: Special value for invalid cases

## Edge Cases

- n <= 0 → return 0
- n > length → return 0
- n == 1 → first rune (index 0)
- n == length → last rune (index length-1)

## Notions

- [rune-literals](https://golang.org/ref/spec#Rune_literals)
