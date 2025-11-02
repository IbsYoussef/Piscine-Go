# printcomb2

**Status:** Optional

## Problem Statement

Write a function `PrintComb2` that prints all possible combinations of two different two-digit numbers in ascending order, separated by ", ". The last combination should not have a trailing comma.

## Expected Function Signature

```go
func PrintComb2()
```

## Expected Output Format

```
00 01, 00 02, 00 03, ..., 98 99
```

(All pairs where first number < second number, separated by ", ", ending with newline)

## Requirements

- Create a file named `printcomb2.go`
- Define package as `package piscine`
- Implement the `PrintComb2` function
- Use `z01.PrintRune()` to print characters
- No parameters, no return value
- Print pairs where first two-digit number < second two-digit number

## Submission Structure

```
printcomb2.go
```

## How to Work on This

1. Write your solution in `student/printcomb2.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printcomb2.go` if you need help

## Hints

- Use four nested loops for digits: a, b (first number), c, d (second number)
- First number: ab, Second number: cd
- Ensure ab < cd
- Start c at a, and d at b+1 (or 0 if c > a)
- Convert digit to rune: `rune('0' + digit)`
- Last combo is 98 99, no comma after it
- Format: "ab cd, " (space between numbers, comma-space between pairs)
