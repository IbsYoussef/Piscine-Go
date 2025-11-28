# split

**Status:** Optional

## Problem Statement

Write a function that receives a string and a separator and returns a `slice of strings` that results from splitting the string `s` by the separator `sep`.

## Expected Function Signature

```go
func Split(s, sep string) []string {

}
```

## Expected Output

```console
$ go run .
[]string{"Hello", "how", "are", "you?"}
```

## Requirements

- Create a file named `split.go`
- Define package as `package piscine`
- Implement the `Split` function
- Split string by the separator
- Return slice of strings
- Handle edge cases (empty separator, separator not found, etc.)

## Submission Structure

```
split.go
```

## How to Work on This

1. Write your solution in `student/split.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/split.go` if you need help

## Key Concepts

- **String searching**: Finding separator in string
- **Substring extraction**: Getting parts between separators
- **Multi-character separator**: Unlike SplitWhiteSpaces (single char)
- **Edge cases**: Empty separator, no separator found

## Examples

```go
Split("HelloHAhowHAareHAyou?", "HA")
Returns: []string{"Hello", "how", "are", "you?"}

Split("a,b,c", ",")
Returns: []string{"a", "b", "c"}

Split("hello", "X")
Returns: []string{"hello"} (separator not found)
```

## Hints

Think about:

- How to find where the separator appears in the string
- How to extract the part before the separator
- How to continue searching in the remaining string
- When to stop splitting
