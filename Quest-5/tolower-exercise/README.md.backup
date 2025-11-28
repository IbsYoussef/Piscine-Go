# tolower

**Status:** Required

## Problem Statement

Write a function that lower cases each letter in a `string`.

The function should convert all uppercase letters (A-Z) to lowercase (a-z), while leaving all other characters unchanged.

## Expected Function Signature
```go
func ToLower(s string) string
```

## Expected Output
```console
hello! how are you?
```

## Requirements

- Create a file named `tolower.go`
- Define package as `package piscine`
- Implement the `ToLower` function
- Convert all uppercase letters to lowercase
- Leave lowercase letters, numbers, symbols unchanged
- Return the modified string

## Submission Structure
```
tolower.go
```

## How to Work on This

1. Write your solution in `student/tolower.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/tolower.go` if you need help

## Hints

- Convert string to rune slice
- Loop through each character
- Check if character is uppercase (A-Z)
- Convert uppercase to lowercase by adding 32
- ASCII: 'A' = 65, 'a' = 97, difference = 32
- Convert rune slice back to string

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ToLower("Hello! How are you?"))
}
```

Output:
```
hello! how are you?
```

## Key Concepts

- **Case conversion**: Changing letter case
- **ASCII arithmetic**: Using character code differences
- **Rune manipulation**: Modifying individual characters
- **String immutability**: Create new string from modified runes

## ASCII Values

- Uppercase 'A' = 65, Lowercase 'a' = 97
- Uppercase 'B' = 66, Lowercase 'b' = 98
- Uppercase 'Z' = 90, Lowercase 'z' = 122
- Difference: 32 (97 - 65 = 32)

## Conversion Formula
```
lowercase = uppercase + 32
```

## Notions

- [strings/ToLower](https://golang.org/pkg/strings/#ToLower)