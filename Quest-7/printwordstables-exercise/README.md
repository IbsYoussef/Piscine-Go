# printwordstables

**Status:** Optional

## Problem Statement

Write a function that receives a `string slice` and prints each element of the slice in a separate line.

## Expected Function Signature

```go
func PrintWordsTables(a []string)
```

## Expected Output

```console
$ go run .
Hello
how
are
you?
```

## Requirements

- Create a file named `printwordstables.go`
- Define package as `package piscine`
- Import `"github.com/01-edu/z01"`
- Implement the `PrintWordsTables` function
- Print each string on a new line
- Use z01.PrintRune for printing

## Submission Structure

```
printwordstables.go
```

## How to Work on This

1. Write your solution in `student/printwordstables.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printwordstables.go` if you need help

## Key Concepts

- **Slice iteration**: Loop through string slice
- **Nested loops**: Loop through each character in each string
- **Printing**: Using z01.PrintRune

## Examples

```go
PrintWordsTables([]string{"Hello", "how", "are", "you?"})
Prints:
Hello
how
are
you?

PrintWordsTables([]string{"one"})
Prints:
one

PrintWordsTables([]string{})
Prints: (nothing)
```
