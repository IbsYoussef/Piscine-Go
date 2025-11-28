# splitwhitespaces

**Status:** Required

## Problem Statement

Write a function that separates the words of a `string` and puts them in a `string slice`.

The separators are spaces, tabs (`\t`), and newlines (`\n`).

## Expected Function Signature

```go
func SplitWhiteSpaces(s string) []string {

}
```

## Expected Output

```console
$ go run .
[]string{"Hello", "how", "are", "you?"}
```

## Requirements

- Create a file named `splitwhitespaces.go`
- Define package as `package piscine`
- Implement the `SplitWhiteSpaces` function
- Split string on whitespace characters (space, tab, newline)
- Return slice of words
- Ignore consecutive whitespace (don't create empty strings)

## Submission Structure

```
splitwhitespaces.go
```

## How to Work on This

1. Write your solution in `student/splitwhitespaces.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/splitwhitespaces.go` if you need help

## Key Concepts

- **String parsing**: Building words character by character
- **Whitespace detection**: Space, tab, newline
- **Word building**: Accumulate characters until whitespace
- **Edge cases**: Multiple consecutive whitespace, trailing whitespace

## Examples

```go
SplitWhiteSpaces("Hello how are you?")
Returns: []string{"Hello", "how", "are", "you?"}

SplitWhiteSpaces("one  two\tthree\nfour")
Returns: []string{"one", "two", "three", "four"}

SplitWhiteSpaces("   trim   ")
Returns: []string{"trim"}
```
