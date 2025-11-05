# printstr

**Status:** Required

## Problem Statement

Write a function that prints one by one the characters of a `string` on the screen.

## Expected Function Signature

```go
func PrintStr(s string)
```

## Expected Output

```console
go run . | cat -e
Hello World!
```

## Requirements

- Create a file named `printstr.go`
- Define package as `package piscine`
- Implement the `PrintStr` function
- Print each character of the string one by one
- Use `z01.PrintRune()` to print each character

## Submission Structure

```
printstr.go
```

## How to Work on This

1. Write your solution in `student/printstr.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printstr.go` if you need help

## Hints

- Iterate through the string using a `for range` loop
- Each character in the string is a `rune`
- Use `z01.PrintRune(r)` to print each rune
- The z01 package is from `github.com/01-edu/z01`

## Example Usage

```go
package main

import "piscine"

func main() {
	piscine.PrintStr("Hello World!")
}
```

## Key Concepts

- **String iteration**: Using `for range` to iterate over string characters
- **Runes**: Characters in Go are represented as runes
- **z01 package**: Custom printing utility used in 01-edu exercises
- **Function without return**: This function prints but doesn't return anything

## Notions

- [01-edu/z01](https://github.com/01-edu/z01)