# swap

**Status:** Required

## Problem Statement

Write a function that takes two pointers to an int (`*int`) and swaps their contents.

## Expected Function Signature

```go
func Swap(a *int, b *int)
```

## Expected Output

```
1
0
```

## Requirements

- Create a file named `swap.go`
- Define package as `package piscine`
- Implement the `Swap` function
- Take two pointers to int as parameters
- Swap the values at those pointers

## Submission Structure

```
swap.go
```

## How to Work on This

1. Write your solution in `student/swap.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/swap.go` if you need help

## Hints

- Use dereference operator `*` to access values
- Go supports multiple assignment: `*a, *b = *b, *a`
- Or use a temporary variable to swap
- The function modifies both original variables