# isprintable

**Status:** Required

## Problem Statement

Write a function that returns `true` if the `string` passed as a parameter contains only printable characters, otherwise, returns `false`.

Printable characters are visible characters that can be displayed on screen, excluding control characters like newline, tab, etc.

## Expected Function Signature
```go
func IsPrintable(s string) bool
```

## Expected Output
```console
true
false
```

## Requirements

- Create a file named `isprintable.go`
- Define package as `package piscine`
- Implement the `IsPrintable` function
- Return true if ALL characters are printable (ASCII 32-126)
- Return false if ANY character is not printable
- Return true for empty strings
- Printable range: space (32) to tilde (~) (126)

## Submission Structure
```
isprintable.go
```

## How to Work on This

1. Write your solution in `student/isprintable.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/isprintable.go` if you need help

## Hints

- Convert string to rune slice
- Loop through each character
- Check if character is NOT in printable range: return false
- Printable range: 32 (space) to 126 (tilde ~)
- Control characters like \n, \t, \r are NOT printable
- Empty string returns true

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("Hello"))    // true (all printable)
	fmt.Println(piscine.IsPrintable("Hello\n"))  // false (has newline)
}
```

Output:
```
true
false
```

## Key Concepts

- **Printable characters**: Visible characters (ASCII 32-126)
- **Control characters**: Non-printable (ASCII 0-31, 127)
- **ASCII range**: Space to tilde
- **Early return**: Return false on first non-printable

## Printable Characters (ASCII 32-126)

- Space (32)
- Punctuation: ! " # $ % & ' ( ) * + , - . /
- Digits: 0-9
- More punctuation: : ; < = > ? @
- Uppercase: A-Z
- More punctuation: [ \ ] ^ _ `
- Lowercase: a-z
- More punctuation: { | } ~
- Tilde (126)

## Non-Printable Characters

- Control characters: \n (newline), \t (tab), \r (carriage return)
- ASCII 0-31: Various control codes
- ASCII 127: DEL character