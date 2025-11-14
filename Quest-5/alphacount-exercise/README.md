# alphacount

**Status:** Required

## Problem Statement

Write a function that counts the letters of a string and returns the count.

The letters are from the Latin alphabet list only, any other characters, symbols or empty spaces shall not be counted.

## Expected Function Signature

```go
func AlphaCount(s string) int
```

## Expected Output

```console
10
```

## Requirements

- Create a file named `alphacount.go`
- Define package as `package piscine`
- Implement the `AlphaCount` function
- Count only alphabetic characters (A-Z, a-z)
- Ignore numbers, spaces, symbols, and special characters
- Return the total count

## Submission Structure

```
alphacount.go
```

## How to Work on This

1. Write your solution in `student/alphacount.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/alphacount.go` if you need help

## Hints

- Convert string to rune slice
- Loop through each character
- Check if character is between 'A'-'Z' or 'a'-'z'
- Increment counter for valid letters
- Return the count

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}
```

Output:

```
10
```

Breakdown: "Hello" (5) + "World" (5) = 10 letters

## Key Concepts

- **Character validation**: Checking if character is alphabetic
- **ASCII ranges**: 'A'-'Z' (65-90), 'a'-'z' (97-122)
- **Filtering**: Ignoring non-alphabetic characters
- **Counting**: Accumulator pattern

## Valid Characters

- Uppercase: A, B, C, ..., Z
- Lowercase: a, b, c, ..., z

## Invalid Characters (not counted)

- Numbers: 0-9
- Spaces
- Symbols: !, @, #, $, %, etc.
- Punctuation: . , ; : ? etc.
