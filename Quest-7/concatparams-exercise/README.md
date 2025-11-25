# concatparams

**Status:** Required

## Problem Statement

Write a function that takes arguments received in parameters and returns them as a `string`. The `string` is the result of all the arguments concatenated with a newline (`\n`) between.

## Expected Function Signature

```go
func ConcatParams(args []string) string
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

- Create a file named `concatparams.go`
- Define package as `package piscine`
- Implement the `ConcatParams` function
- Concatenate all strings from slice
- Add newline (`\n`) between each string
- No newline after the last string

## Submission Structure

```
concatparams.go
```

## How to Work on This

1. Write your solution in `student/concatparams.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/concatparams.go` if you need help

## Key Concepts

- **String concatenation**: Building result string
- **Newline handling**: Add `\n` between elements
- **Edge case**: No newline after last element

## Examples

```go
ConcatParams([]string{"Hello", "how", "are", "you?"})
Returns: "Hello\nhow\nare\nyou?"

ConcatParams([]string{"one"})
Returns: "one"

ConcatParams([]string{})
Returns: ""
```
