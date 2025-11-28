# isnumeric

**Status:** Required

## Problem Statement

Write a function that returns `true` if the `string` passed as a parameter contains only numerical characters, otherwise, returns `false`.

## Expected Function Signature
```go
func IsNumeric(s string) bool
```

## Expected Output
```console
true
false
```

## Requirements

- Create a file named `isnumeric.go`
- Define package as `package piscine`
- Implement the `IsNumeric` function
- Return true if ALL characters are digits (0-9)
- Return false if ANY character is not a digit
- Return false for empty strings
- Return false if string contains letters, symbols, or spaces

## Submission Structure
```
isnumeric.go
```

## How to Work on This

1. Write your solution in `student/isnumeric.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/isnumeric.go` if you need help

## Hints

- Handle edge case: empty string returns false
- Convert string to rune slice
- Loop through each character
- Check if character is NOT a digit: return false immediately
- If loop completes without finding non-digit, return true
- Digit check: `rune >= '0' && rune <= '9'`

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsNumeric("010203"))     // true (only digits)
	fmt.Println(piscine.IsNumeric("01,02,03"))   // false (has commas)
}
```

Output:
```
true
false
```

## Key Concepts

- **Numeric validation**: Checking if all characters are digits
- **ASCII range**: Digits 0-9 (48-57)
- **Early return**: Return false on first non-digit
- **Logical negation**: Using NOT operator

## Valid (returns true)

- "123"
- "010203"
- "0"

## Invalid (returns false)

- "12a3" (has letter)
- "01,02,03" (has commas)
- "12 34" (has space)
- "" (empty string)