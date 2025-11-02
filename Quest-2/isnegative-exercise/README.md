# isnegative

**Status:** Required

## Problem Statement

Write a function `IsNegative` that prints 'T' (true) if the integer passed as parameter is negative, otherwise prints 'F' (false). Always print a newline after the character.

## Expected Function Signature

```go
func IsNegative(nb int)
```

## Expected Output Format

For each call:

- If negative: `T\n`
- If zero or positive: `F\n`

## Requirements

- Create a file named `isnegative.go`
- Define package as `package piscine`
- Implement the `IsNegative` function
- Use `z01.PrintRune()` to print characters

## Submission Structure

```
isnegative.go
```

## How to Work on This

1. Write your solution in `student/isnegative.go`
2. Run `./run.sh` to see your output (tests with multiple values)
3. Run `./test.sh` to check if it works
4. Compare with `solutions/isnegative.go` if you need help

## Hints

- Check if `nb < 0` for negative numbers
- Use if/else to decide between 'T' and 'F'
- Print the character with `z01.PrintRune('T')` or `z01.PrintRune('F')`
- Always print newline: `z01.PrintRune('\n')`
- Function doesn't return anything, just prints
