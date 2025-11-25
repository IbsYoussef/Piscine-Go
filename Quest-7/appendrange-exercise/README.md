# appendrange

**Status:** Required

## Problem Statement

Write a function that takes an `int` min and an `int` max as parameters. The function should return a slice of `int`s with all the values between min and max.

- Min is **included**
- Max is **excluded**
- If min >= max, return a `nil` slice
- `make` is **not allowed** for this exercise

## Expected Function Signature

```go
func AppendRange(min, max int) []int
```

## Expected Output

```console
$ go run .
[5 6 7 8 9]
[]
```

## Requirements

- Create a file named `appendrange.go`
- Define package as `package piscine`
- Implement the `AppendRange` function
- Return slice with values from min to max-1
- Return nil slice if min >= max
- Do NOT use `make`

## Submission Structure

```
appendrange.go
```

## How to Work on This

1. Write your solution in `student/appendrange.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/appendrange.go` if you need help

## Key Concepts

- **Slice creation**: Using append with empty slice
- **Range generation**: Loop from min to max
- **Edge cases**: min >= max returns nil
- **Nil slice**: Empty slice literal `[]int(nil)`

## Examples

```go
AppendRange(5, 10) → [5, 6, 7, 8, 9]
AppendRange(10, 5) → []
AppendRange(1, 2) → [1]
AppendRange(5, 5) → []
```
