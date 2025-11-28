# islower

**Status:** Required

## Problem Statement

Write a function that returns `true` if the `string` passed as the parameter contains only lowercase characters, otherwise, returns `false`.

## Expected Function Signature
```go
func IsLower(s string) bool
```

## Expected Output
```console
true
false
```

## Requirements

- Create a file named `islower.go`
- Define package as `package piscine`
- Implement the `IsLower` function
- Return true if ALL characters are lowercase letters (a-z)
- Return false if ANY character is not lowercase
- Return false for empty strings
- Return false if string contains numbers, symbols, or spaces

## Submission Structure
```
islower.go
```

## How to Work on This

1. Write your solution in `student/islower.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/islower.go` if you need help

## Hints

- Handle edge case: empty string returns false
- Convert string to rune slice
- Loop through each character
- Check if character is NOT lowercase: return false immediately
- If loop completes without finding non-lowercase, return true
- Lowercase check: `rune >= 'a' && rune <= 'z'`

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsLower("hello"))   // true (all lowercase)
	fmt.Println(piscine.IsLower("hello!"))  // false (has '!')
}
```

Output:
```
true
false
```

## Key Concepts

- **Character validation**: Checking if all characters match criteria
- **ASCII range**: Lowercase letters a-z (97-122)
- **Early return**: Return false on first non-lowercase
- **Logical negation**: Using NOT operator

## Valid (returns true)

- "hello"
- "abc"
- "z"

## Invalid (returns false)

- "Hello" (has uppercase)
- "hello!" (has symbol)
- "hello123" (has numbers)
- "hello " (has space)
- "" (empty string)