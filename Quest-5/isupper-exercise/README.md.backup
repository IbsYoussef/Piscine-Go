# isupper

**Status:** Required

## Problem Statement

Write a function that returns `true` if the `string` passed as parameter contains only uppercase characters, otherwise, returns `false`.

## Expected Function Signature
```go
func IsUpper(s string) bool
```

## Expected Output
```console
true
false
```

## Requirements

- Create a file named `isupper.go`
- Define package as `package piscine`
- Implement the `IsUpper` function
- Return true if ALL characters are uppercase letters (A-Z)
- Return false if ANY character is not uppercase
- Return false for empty strings
- Return false if string contains numbers, symbols, or spaces

## Submission Structure
```
isupper.go
```

## How to Work on This

1. Write your solution in `student/isupper.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/isupper.go` if you need help

## Hints

- Handle edge case: empty string returns false
- Convert string to rune slice
- Loop through each character
- Check if character is NOT uppercase: return false immediately
- If loop completes without finding non-uppercase, return true
- Uppercase check: `rune >= 'A' && rune <= 'Z'`

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsUpper("HELLO"))   // true (all uppercase)
	fmt.Println(piscine.IsUpper("HELLO!"))  // false (has '!')
}
```

Output:
```
true
false
```

## Key Concepts

- **Character validation**: Checking if all characters match criteria
- **ASCII range**: Uppercase letters A-Z (65-90)
- **Early return**: Return false on first non-uppercase
- **Logical negation**: Using NOT operator

## Valid (returns true)

- "HELLO"
- "ABC"
- "Z"

## Invalid (returns false)

- "Hello" (has lowercase)
- "HELLO!" (has symbol)
- "HELLO123" (has numbers)
- "HELLO " (has space)
- "" (empty string)