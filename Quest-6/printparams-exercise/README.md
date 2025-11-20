# printparams

**Status:** Required

## Problem Statement

Write a **program** that prints the arguments received in the command line.

Each argument should be printed on a new line.

## Expected Output

```console
$ go run . choumi is the best cat
choumi
is
the
best
cat
```

## Requirements

- Create a folder named `printparams`
- Create `main.go` inside the folder
- Define `package main`
- Import `"github.com/01-edu/z01"` and `"os"`
- Print each command-line argument on a new line
- Do NOT print the program name (skip `os.Args[0]`)

## Submission Structure

```
printparams/
└── main.go
```

## How to Work on This

1. Write your solution in `student/printparams/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printparams/main.go` if you need help

## Key Concepts

- **Command-line arguments**: `os.Args` slice
- **Skipping first argument**: `os.Args[1:]` (program name is at index 0)
- **Nested loops**: Loop through arguments, then through characters
- **String iteration**: Print each rune

## Hints

- `os.Args[0]` is the program name - skip it!
- `os.Args[1:]` gives you all arguments except the program name
- Loop through each argument
- For each argument, loop through each character
- Print newline after each argument
