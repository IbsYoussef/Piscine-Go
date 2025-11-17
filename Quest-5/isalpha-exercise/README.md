# isalpha

**Status:** Required

## Problem Statement

Write a function that returns `true` if the `string` passed as the parameter only contains alphanumerical characters or is empty, otherwise, returns `false`.

Alphanumerical characters include:
- Uppercase letters: A-Z
- Lowercase letters: a-z
- Digits: 0-9

## Expected Function Signature
```go
func IsAlpha(s string) bool
```

## Expected Output
```console
false
true
false
true
```

## Requirements

- Create a file named `isalpha.go`
- Define package as `package piscine`
- Implement the `IsAlpha` function
- Return true if ALL characters are alphanumeric (A-Z, a-z, 0-9)
- Return true for empty string
- Return false if ANY character is not alphanumeric (spaces, symbols, etc.)

## Submission Structure
```
isalpha.go
```

## How to Work on This

1. Write your solution in `student/isalpha.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/isalpha.go` if you need help

## Hints

- Convert string to rune slice
- Loop through each character
- Check if character is NOT alphanumeric: return false
- Alphanumeric means: (A-Z) OR (a-z) OR (0-9)
- Use negation with multiple conditions
- Empty string returns true

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsAlpha("Hello! How are you?"))  // false (has space and !)
	fmt.Println(piscine.IsAlpha("HelloHowareyou"))       // true (only letters)
	fmt.Println(piscine.IsAlpha("What's this 4?"))       // false (has symbols)
	fmt.Println(piscine.IsAlpha("Whatsthis4"))           // true (letters + digit)
}
```

Output:
```
false
true
false
true
```

## Key Concepts

- **Alphanumeric validation**: Letters and numbers only
- **Multiple conditions**: Checking three ranges
- **Logical operators**: Using NOT and AND
- **Empty string**: Special case (returns true)

## Valid Characters (alphanumeric)

- Uppercase: A-Z
- Lowercase: a-z
- Digits: 0-9

## Invalid Characters (not alphanumeric)

- Spaces
- Symbols: !, @, #, $, %, etc.
- Punctuation: . , ; : ? etc.