# printdigits

**Status:** Required

## Problem Statement

Write a program that prints all digits (0-9) in ascending order on a single line, followed by a newline character.

## Expected Output Format

```
0123456789
```

(One line: all digits 0-9)

## Requirements

- Create a folder named `printdigits/`
- Inside, create `main.go`
- Print digits 0-9 on one line
- End with a newline
- Use `z01.PrintRune()` to print characters

## Submission Structure

```
printdigits/
└── main.go
```

## How to Work on This

1. Write your solution in `student/printdigits/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printdigits/main.go` if you need help

## Hints

- Use a for loop from '0' to '9'
- Rune literals: '0', '1', '2', etc.
- Loop: `for i := '0'; i <= '9'; i++`
- Print each digit with `z01.PrintRune(i)`
- Don't forget the newline at the end: `z01.PrintRune('\n')`
