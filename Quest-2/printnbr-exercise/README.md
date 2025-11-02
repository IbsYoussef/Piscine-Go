# printnbr

**Status:** Optional

## Problem Statement

Write a function `PrintNbr` that prints an integer passed as parameter using only `z01.PrintRune`. Must handle all int values including negative numbers and zero.

## Expected Function Signature

```go
func PrintNbr(n int)
```

## Expected Output Format

For n = -123:

```
-123
```

For n = 0:

```
0
```

For n = 456:

```
456
```

## Requirements

- Create a file named `printnbr.go`
- Define package as `package piscine`
- Implement the `PrintNbr` function
- Use only `z01.PrintRune()` for output
- Handle negative numbers, zero, and positive numbers
- No conversion to int64 or other types

## Submission Structure

```
printnbr.go
```

## How to Work on This

1. Write your solution in `student/printnbr.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printnbr.go` if you need help

## Hints

- Handle zero as special case: print '0' and return
- For negative: print '-', then work with digits
- Extract digits using modulo: `digit := n % 10`
- Store digits in array (up to 19 digits for int range)
- Print digits in reverse order
- Convert digit to rune: `rune('0' + digit)`
- Handle negative digits: `if d < 0 { d = -d }`
