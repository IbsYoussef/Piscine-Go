# printreversealphabet

**Status:** Required

## Problem Statement

Write a program that prints the lowercase Latin alphabet in reverse order (z-a) on a single line, followed by a newline character.

## Expected Output Format

```
zyxwvutsrqponmlkjihgfedcba
```

(One line: all lowercase letters z-a in reverse)

## Requirements

- Create a folder named `printreversealphabet/`
- Inside, create `main.go`
- Print the alphabet z-a in reverse order on one line
- End with a newline
- Use `z01.PrintRune()` to print characters

## Submission Structure

```
printreversealphabet/
└── main.go
```

## How to Work on This

1. Write your solution in `student/printreversealphabet/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printreversealphabet/main.go` if you need help

## Hints

- Use a for loop from 'z' to 'a'
- Rune literals: 'z', 'y', 'x', etc.
- Loop: `for i := 'z'; i >= 'a'; i--` (note: decrement)
- Print each character with `z01.PrintRune(i)`
- Don't forget the newline at the end: `z01.PrintRune('\n')`
