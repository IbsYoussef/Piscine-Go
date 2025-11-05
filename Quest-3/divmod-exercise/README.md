# divmod

**Status:** Required

## Problem Statement

Write a function that divides two integers and stores the result and remainder using pointers.

## Expected Function Signature

```go
func DivMod(a int, b int, div *int, mod *int)
```

## Expected Output

```
6
1
```

## Requirements

- Create a file named `divmod.go`
- Define package as `package piscine`
- Implement the `DivMod` function
- Divide `a` by `b`
- Store the division result in the int pointed to by `div`
- Store the remainder in the int pointed to by `mod`

## Submission Structure

```
divmod.go
```

## How to Work on This

1. Write your solution in `student/divmod.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/divmod.go` if you need help

## Hints

- Use `/` operator for division
- Use `%` operator for modulo (remainder)
- Store results using dereference: `*div = result`
- Example: 13 / 2 = 6 remainder 1