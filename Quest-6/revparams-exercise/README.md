# revparams

**Status:** Required

## Problem Statement

Write a **program** that prints the arguments received in the command line in **reverse order**.

Each argument should be printed on a new line, starting from the last argument.

## Expected Output

```console
$ go run . choumi is the best cat
cat
best
the
is
choumi
```

## Requirements

- Create a folder named `revparams`
- Create `main.go` inside the folder
- Define `package main`
- Import `"github.com/01-edu/z01"` and `"os"`
- Print each command-line argument on a new line
- Print in REVERSE order (last argument first)
- Do NOT print the program name

## Submission Structure

```
revparams/
└── main.go
```

## How to Work on This

1. Write your solution in `student/revparams/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/revparams/main.go` if you need help

## Key Concepts

- **Reverse iteration**: Loop backwards through slice
- **Index-based loop**: Using `for i := len(args)-1; i >= 0; i--`
- **Command-line arguments**: `os.Args[1:]`

## Hints

- Get arguments with `os.Args[1:]`
- Loop backwards: start at `len(args)-1`, go down to `0`
- Use index-based loop instead of range
- Print each argument's characters, then newline
