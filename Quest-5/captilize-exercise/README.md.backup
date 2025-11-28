# capitalize

**Status:** Required

## Problem Statement

Write a function that capitalizes the first letter of each word **and** lowercases the rest.

A word is a sequence of **alphanumeric** characters (letters and digits).

## Expected Function Signature
```go
func Capitalize(s string) string
```

## Expected Output
```console
Hello! How Are You? How+Are+Things+4you?
```

## Requirements

- Create a file named `capitalize.go`
- Define package as `package piscine`
- Implement the `Capitalize` function
- Capitalize first letter of each word
- Lowercase all other letters in each word
- A word is alphanumeric characters (letters + digits)
- Non-alphanumeric characters separate words

## Submission Structure
```
capitalize.go
```

## How to Work on This

1. Write your solution in `student/capitalize.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/capitalize.go` if you need help

## Hints

- Track whether you're inside a word with a boolean flag
- When you find first letter of a word: capitalize it
- When you find subsequent letters: lowercase them
- Digits are part of words (but don't capitalize)
- Non-alphanumeric characters end words
- Use ASCII arithmetic: -32 for uppercase, +32 for lowercase

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
}
```

Output:
```
Hello! How Are You? How+Are+Things+4you?
```

Analysis:
- "Hello" → "Hello" (first letter capitalized)
- "How" → "How" (first letter capitalized)
- "are" → "Are" (first letter capitalized)
- "you" → "You" (first letter capitalized)
- "How+are+things+4you" → "How+Are+Things+4you" (+ separates words)

## Key Concepts

- **Word detection**: Alphanumeric sequences
- **State tracking**: Are we inside a word?
- **Case conversion**: First letter uppercase, rest lowercase
- **Character classification**: Letters, digits, separators

## Word Definition

- Alphanumeric: Letters (A-Z, a-z) and digits (0-9)
- Word boundaries: Spaces, punctuation, symbols