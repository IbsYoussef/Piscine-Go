# printcombn

**Status:** Bonus

## Problem Statement

Write a function `PrintCombN` that prints all combinations of n different digits (0 < n < 10) in ascending order, separated by ", ". The last combination should not have a trailing comma.

## Expected Function Signature

```go
func PrintCombN(n int)
```

## Expected Output Format

For n = 1:

```
0, 1, 2, 3, 4, 5, 6, 7, 8, 9
```

For n = 3:

```
012, 013, 014, ..., 789
```

For n = 9:

```
012345678, 012345679, ..., 123456789
```

## Requirements

- Create a file named `printcombn.go`
- Define package as `package piscine`
- Implement the `PrintCombN` function
- Can use `fmt` package (not just z01)
- Handle n from 1 to 9
- Return immediately if n <= 0 or n >= 10

## Submission Structure

```
printcombn.go
```

## How to Work on This

1. Write your solution in `student/printcombn.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printcombn.go` if you need help

## Hints

- Use recursion to generate combinations
- Helper function: `helper(current string, start int, n int)`
- Base case: when `len(current) == n`, print it
- Recursive case: try adding each digit from start to 9
- Calculate max combination to know when to stop printing commas
- Max for n=3 is "789", for n=5 is "56789", etc.
- Use `fmt.Print()` and `fmt.Println()` for output
