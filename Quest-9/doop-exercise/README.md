# doop

**Status:** Optional

## Problem Statement

Write a program that is called `doop`.

The program has to be used with three arguments:

- A value
- An operator, one of : `+`, `-`, `/`, `*`, `%`
- Another value

In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

The program has to handle the modulo and division operations by 0 as shown on the output examples below.

## Expected Output
```console
$ go run .
$ go run . 1 + 1 | cat -e
2
$
$ go run . hello + 1
$ go run . 1 p 1
$ go run . 1 / 0 | cat -e
No division by 0
$
$ go run . 1 % 0 | cat -e
No modulo by 0
$
$ go run . 9223372036854775807 + 1
$ go run . -9223372036854775809 - 3
$ go run . 9223372036854775807 "*" 3
$ go run . 1 "*" 1
1
$ go run . 1 "*" -1
-1
```

## Requirements

- Create a folder named `doop`
- Create `main.go` inside the folder
- Handle three arguments: value, operator, value
- Valid operators: `+`, `-`, `/`, `*`, `%`
- Handle division by zero
- Handle modulo by zero
- Handle overflow (print nothing)
- Handle invalid inputs (print nothing)

## Submission Structure
```
doop/
└── main.go
```