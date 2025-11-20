# sortparams

**Status:** Optional

## Problem Statement

Write a **program** that prints the arguments received in the command line in **ASCII order**.

ASCII order means sorting by the numeric values of characters (numbers before uppercase before lowercase).

## Expected Output

```console
$ go run . 1 a 2 A 3 b 4 C
1
2
3
4
A
C
a
b
```

## Requirements

- Create a folder named `sortparams`
- Create `main.go` inside the folder
- Define `package main`
- Import `"github.com/01-edu/z01"` and `"os"`
- Sort arguments in ASCII order
- Print each argument on a new line
- Do NOT print the program name

## Submission Structure

```
sortparams/
└── main.go
```

## How to Work on This

1. Write your solution in `student/sortparams/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/sortparams/main.go` if you need help

## Key Concepts

- **Sorting algorithms**: Bubble sort, selection sort, etc.
- **ASCII comparison**: Comparing strings directly
- **String comparison**: Go can compare strings with `>`, `<`, `==`

## Hints

- Get arguments with `os.Args[1:]`
- Sort the slice using bubble sort or any sorting algorithm
- In Go, you can compare strings directly: `"a" > "b"` is false
- ASCII order: digits (0-9) < uppercase (A-Z) < lowercase (a-z)
- After sorting, print each argument

## ASCII Order

- Numbers: 0-9 (ASCII 48-57)
- Uppercase: A-Z (ASCII 65-90)
- Lowercase: a-z (ASCII 97-122)
