# pointone

**Status:** Required

## Problem Statement

Write a function that takes a pointer to an int as argument and gives to this int the value of 1.

## Expected Function Signature

```go
func PointOne(n *int)
```

## Expected Output

```
1
```

## Requirements

- Create a file named `pointone.go`
- Define package as `package piscine`
- Implement the `PointOne` function
- Take a pointer to int as parameter
- Set the value of the int to 1

## Submission Structure

```
pointone.go
```

## How to Work on This

1. Write your solution in `student/pointone.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/pointone.go` if you need help

## Hints

- Use the dereference operator `*` to access the value
- Syntax: `*n = 1` sets the value at the pointer
- Pass address with `&` when calling: `PointOne(&n)`