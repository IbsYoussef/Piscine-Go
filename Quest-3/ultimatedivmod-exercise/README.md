# ultimatedivmod

**Status:** Required

## Problem Statement

Write a function that divides two integers and stores both results back into the same pointers.

## Expected Function Signature

```go
func UltimateDivMod(a *int, b *int)
```

## Expected Output

```
6
1
```

## Requirements

- Create a file named `ultimatedivmod.go`
- Define package as `package piscine`
- Implement the `UltimateDivMod` function
- Divide the value at `a` by the value at `b`
- Store the division result back in `a`
- Store the remainder back in `b`

## Submission Structure

```
ultimatedivmod.go
```

## How to Work on This

1. Write your solution in `student/ultimatedivmod.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/ultimatedivmod.go` if you need help

## Hints

- Read values first: `*a` and `*b`
- Calculate division: `*a / *b`
- Calculate modulo: `*a % *b`
- Store back to same pointers
- Be careful: don't lose values before calculating!
- Use temporary variables to store results before writing back

## Example

```go
a := 13
b := 2
UltimateDivMod(&a, &b)
fmt.Println(a)  // Output: 6
fmt.Println(b)  // Output: 1
```

## Key Concepts

- **Pointers**: Working with memory addresses
- **Dereferencing**: Using `*` to read/write pointer values
- **Order of operations**: Reading before writing to avoid data loss
- **Division operator**: `/` for quotient
- **Modulo operator**: `%` for remainder