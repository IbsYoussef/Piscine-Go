# printcomb

**Status:** Required

## Problem Statement

Write a function `PrintComb` that prints all unique combinations of three different digits in ascending order, separated by ", ". The last combination should not have a trailing comma.

## Expected Function Signature

```go
func PrintComb()
```

## Expected Output Format

```
012, 013, 014, 015, 016, 017, 018, 019, 023, 024, ..., 789
```

(All combinations where a < b < c, separated by ", ", ending with newline)

## Requirements

- Create a file named `printcomb.go`
- Define package as `package piscine`
- Implement the `PrintComb` function
- Use `z01.PrintRune()` to print characters
- No parameters, no return value
- Print all combinations where first < second < third digit

## Submission Structure

```
printcomb.go
```

## How to Work on This

1. Write your solution in `student/printcomb.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printcomb.go` if you need help

## Hints

- Use three nested loops: a from '0' to '7', b from a+1 to '8', c from b+1 to '9'
- This ensures a < b < c automatically
- Print each digit with `z01.PrintRune()`
- Check if not last combo (789) before printing ", "
- Last combo check: `a != '7' || b != '8' || c != '9'`
- Don't forget final newline
