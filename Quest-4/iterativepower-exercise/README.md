# iterativepower

**Status:** Required

## Problem Statement

Write an **iterative** function that returns the value of `nb` to the power of `power`.

Negative powers will return `0`. Overflows do **not** have to be dealt with.

## Expected Function Signature

```go
func IterativePower(nb int, power int) int
```

## Expected Output

```console
64
```

## Requirements

- Create a file named `iterativepower.go`
- Define package as `package piscine`
- Implement the `IterativePower` function
- Use iteration (loops), **not recursion**
- Negative power returns 0
- Any number to power 0 is 1
- Calculate: nb^power = nb × nb × ... (power times)

## Hints

- Handle negative power: return 0
- Handle power == 0: return 1 (any number^0 = 1)
- Initialize result to 1
- Loop power times, multiplying by nb each time
