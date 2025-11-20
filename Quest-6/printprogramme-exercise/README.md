# printprogramname

**Status:** Required

## Problem Statement

Write a **program** that prints the name of the program.

The program should print its own executable name, regardless of what it's compiled as.

## Expected Output
```console
$ go build main.go
$ ./main
main
$ go build
$ ./printprogramname
printprogramname
$ go build -o Nessy
$ ./Nessy
Nessy
```

## Requirements

- Create a folder named `printprogramname`
- Create `main.go` inside the folder
- Define `package main`
- Import `"github.com/01-edu/z01"` and `"os"`
- Print the program name (not the full path)
- Output must end with a newline

## Submission Structure
```
printprogramname/
└── main.go
```

## How to Work on This

1. Write your solution in `student/printprogramname/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printprogramname/main.go` if you need help

## Key Concepts

- **Command-line arguments**: `os.Args` slice
- **Program name**: First element of `os.Args[0]`
- **Path extraction**: Getting filename from full path
- **String manipulation**: Finding last slash, extracting substring

## Hints

- `os.Args[0]` contains the full path to the executable
- You need to extract just the filename (after the last `/`)
- Use z01.PrintRune to print each character
- Don't forget the newline at the end

## Example

If the full path is `/home/user/program`, you should print `program`.
If the full path is just `main`, you should print `main`.