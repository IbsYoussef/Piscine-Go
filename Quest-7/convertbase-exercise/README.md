# convertbase

**Status:** Bonus

## Problem Statement

Write a function that receives three arguments:

- `nbr`: A string representing a numeric value in a base
- `baseFrom`: A string representing the base `nbr` is using
- `baseTo`: A string representing the base `nbr` should be represented in

Convert a number from one base to another base.

Only valid bases will be tested.
Negative numbers will not be tested.

## Expected Function Signature

```go
func ConvertBase(nbr, baseFrom, baseTo string) string {

}
```

## Expected Output

```console
$ go run .
43
```

## Requirements

- Create a file named `convertbase.go`
- Define package as `package piscine`
- Implement the `ConvertBase` function
- Convert from any base to any base
- Return string representation in target base

## Submission Structure

```
convertbase.go
```

## How to Work on This

1. Write your solution in `student/convertbase.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/convertbase.go` if you need help

## Key Concepts

- **Base conversion**: Two-step process
- **Step 1**: Convert from baseFrom to decimal (like AtoiBase)
- **Step 2**: Convert from decimal to baseTo (like PrintNbrBase)
- **Base validation**: Check for valid bases

## Examples

```go
ConvertBase("101011", "01", "0123456789")
Returns: "43" (binary to decimal)

ConvertBase("2A", "0123456789ABCDEF", "01")
Returns: "101010" (hex to binary)

ConvertBase("52", "0123456789", "012345678")
Returns: "57" (decimal to octal)
```

## Algorithm

1. Convert `nbr` from `baseFrom` to decimal integer
2. Convert decimal integer to `baseTo`
3. Return result as string
