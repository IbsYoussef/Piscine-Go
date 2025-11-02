# printalphabet

**Status:** Required

## Problem Statement

Write a program that prints the lowercase Latin alphabet (a–z) on a single line, followed by a newline character.

## Expected Output Format

```
abcdefghijklmnopqrstuvwxyz
```

(One line: all lowercase letters a-z)

## Requirements

- Create a folder named `printalphabet/`
- Inside, create `main.go`
- Print the alphabet a-z on one line
- End with a newline
- Use `z01.PrintRune()` to print characters

## Submission Structure

```
printalphabet/
└── main.go
```

## How to Work on This

1. Write your solution in `student/printalphabet/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printalphabet/main.go` if you need help

## Hints

- Use a for loop from 'a' to 'z'
- Rune literals: 'a', 'b', 'c', etc.
- Loop: `for i := 'a'; i <= 'z'; i++`
- Print each character with `z01.PrintRune(i)`
- Don't forget the newline at the end: `z01.PrintRune('\n')`
