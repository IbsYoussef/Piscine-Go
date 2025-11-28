# toupper

**Status:** Required

## Problem Statement

Write a function that capitalizes each letter in a `string`.

The function should convert all lowercase letters (a-z) to uppercase (A-Z), while leaving all other characters unchanged.

## Expected Function Signature
```go
func ToUpper(s string) string
```

## Expected Output
```console
HELLO! HOW ARE YOU?
```

## Requirements

- Create a file named `toupper.go`
- Define package as `package piscine`
- Implement the `ToUpper` function
- Convert all lowercase letters to uppercase
- Leave uppercase letters, numbers, symbols unchanged
- Return the modified string

## Submission Structure
```
toupper.go
```

## How to Work on This

1. Write your solution in `student/toupper.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/toupper.go` if you need help

## Hints

- Convert string to rune slice
- Loop through each character
- Check if character is lowercase (a-z)
- Convert lowercase to uppercase by subtracting 32
- ASCII: 'a' = 97, 'A' = 65, difference = 32
- Convert rune slice back to string

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ToUpper("Hello! How are you?"))
}
```

Output:
```
HELLO! HOW ARE YOU?
```

## Key Concepts

- **Case conversion**: Changing letter case
- **ASCII arithmetic**: Using character code differences
- **Rune manipulation**: Modifying individual characters
- **String immutability**: Create new string from modified runes

## ASCII Values

- Lowercase 'a' = 97, Uppercase 'A' = 65
- Lowercase 'b' = 98, Uppercase 'B' = 66
- Lowercase 'z' = 122, Uppercase 'Z' = 90
- Difference: 32 (97 - 65 = 32)

## Conversion Formula
```
uppercase = lowercase - 32
```

## Notions

- [strings/ToUpper](https://golang.org/pkg/strings/#ToUpper)