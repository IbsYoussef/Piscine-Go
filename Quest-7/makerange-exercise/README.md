# makerange

**Status:** Required

## Problem Statement

Write a function that takes an `int` min and an `int` max as parameters. The function must return a slice of `int`s with all the values between min and max.

- Min is **included**
- Max is **excluded**
- If min >= max, return a `nil` slice
- `append` is **not allowed** for this exercise (use `make`)

## Expected Function Signature

```go
func MakeRange(min, max int) []int
```

## Expected Output

```console
$ go run .
[5 6 7 8 9]
[]
```

## Requirements

- Create a file named `makerange.go`
- Define package as `package piscine`
- Implement the `MakeRange` function
- Return slice with values from min to max-1
- Return nil if min >= max
- Use `make` (do NOT use `append`)

## Submission Structure

```
makerange.go
```

## How to Work on This

1. Write your solution in `student/makerange.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/makerange.go` if you need help

## Key Concepts

- **Slice creation with make**: Pre-allocate size
- **Direct indexing**: Set values by index
- **Range calculation**: length = max - min
- **Nil return**: When min >= max

## Examples

```go
MakeRange(5, 10) → [5, 6, 7, 8, 9]
MakeRange(10, 5) → nil
MakeRange(1, 2) → [1]
MakeRange(5, 5) → nil
```
