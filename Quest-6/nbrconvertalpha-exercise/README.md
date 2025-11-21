# nbrconvertalpha

**Status:** Optional

## Problem Statement

Write a **program** that prints the corresponding letter in the `n` position of the latin alphabet, where `n` is each argument received.

- `1` → `a`, `2` → `b`, ..., `26` → `z`
- If `n` is not valid (not 1-26) or not an integer → print space (" ")
- Flag `--upper`: prints result in uppercase (always first argument)

## Expected Output

```console
$ go run .

$ go run . 8 5 12 12 15
hello
$ go run . 12 5 7 5 14 56 4 1 18 25
legen dary
$ go run . 32 86 h

$ go run . --upper 8 5 25
HEY
```

## Requirements

- Create a folder named `nbrconvertalpha`
- Create `main.go` inside the folder
- Define `package main`
- Import `"github.com/01-edu/z01"` and `"os"`
- Convert numbers to letters (1=a, 2=b, etc.)
- Print space for invalid inputs
- Handle `--upper` flag
- End with newline

## Submission Structure

```
nbrconvertalpha/
└── main.go
```

## How to Work on This

1. Write your solution in `student/nbrconvertalpha/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/nbrconvertalpha/main.go` if you need help

## Key Concepts

- **Number to letter conversion**: ASCII arithmetic
- **String to int parsing**: Manual conversion
- **Flag handling**: Check first argument
- **Validation**: Check if number is 1-26
- **Case conversion**: Uppercase vs lowercase

## Hints

- Check if first arg is `--upper`, set a flag
- For each argument, convert string to number (manually!)
- Validate: is it a number? Is it 1-26?
- If valid: convert to letter (1 → 'a', 2 → 'b', etc.)
- If invalid: print space
- Apply uppercase if flag is set
