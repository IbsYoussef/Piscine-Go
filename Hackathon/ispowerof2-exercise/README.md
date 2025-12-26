# ispowerof2

## Problem Statement

Write a program that determines if a given number is a power of 2. A number `n` is a power of 2 if it meets the following condition: **n = 2<sup>m</sup>** where m is another integer number.

For example, **4 = 2<sup>2</sup>** or **16 = 2<sup>4</sup>** are power of 2 numbers while 24 is not.

This program must print `true` if the given number is a power of 2, otherwise it has to print `false`.

- If there is more than one or no argument, the program should print nothing.
- When there is only one argument, it will always be a positive valid `int`.

## Expected Output

```console
$ go run . 2 | cat -e
true$
$ go run . 64
true
$ go run . 513
false
$ go run .
$ go run . 64 1024
```

## Requirements

- Create a folder named `ispowerof2`
- Create `main.go` inside the folder
- Check if exactly one argument is provided
- Convert string argument to integer
- Determine if the number is a power of 2
- Print `true` or `false`

## Submission Structure

```
ispowerof2/
└── main.go
```

## How to Work on This

1. Write your solution in `student/ispowerof2/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/ispowerof2/main.go` if you need help

## Hint

A clever way to check if a number is a power of 2 uses bitwise operations:

- Powers of 2 have exactly ONE bit set in binary
- Use the expression: `x > 0 && (x & (x-1)) == 0`
