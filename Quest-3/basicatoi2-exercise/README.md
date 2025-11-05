# basicatoi2

**Status:** Optional

## Problem Statement

Write a function that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number defined as a `string` in a number defined as an `int`.

`Atoi` returns `0` if the `string` is not considered as a valid number. For this exercise **non-valid `string` chains will be tested**. Some will contain non-digits characters.

The handling of the signs `+` or `-` does not have to be taken into account.

## Expected Function Signature

```go
func BasicAtoi2(s string) int
```

## Expected Output

```console
$ go run .
12345
12345
0
0
$
```

## Requirements

- Create a file named `basicatoi2.go`
- Define package as `package piscine`
- Implement the `BasicAtoi2` function
- Convert valid digit strings to integers
- Return `0` for invalid strings (containing non-digits)
- No sign handling needed (`+` or `-`)

## Submission Structure

```
basicatoi2.go
```

## How to Work on This

1. Write your solution in `student/basicatoi2.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/basicatoi2.go` if you need help

## Hints

- Initialize result to 0
- Iterate through each character
- Check if character is valid: `char >= '0' && char <= '9'`
- If invalid, return 0 immediately
- If valid, convert and add: `result = result * 10 + int(char - '0')`
- Return result at the end

## Example Usage

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BasicAtoi2("12345"))
	fmt.Println(piscine.BasicAtoi2("0000000012345"))
	fmt.Println(piscine.BasicAtoi2("012 345"))
	fmt.Println(piscine.BasicAtoi2("Hello World!"))
}
```

## Key Concepts

- **Input validation**: Checking each character is valid
- **Early return**: Return 0 as soon as invalid character found
- **ASCII ranges**: '0' to '9' are valid digits
- **Character comparison**: Using < and > with characters

## Difference from BasicAtoi

- **BasicAtoi**: Assumes all input is valid
- **BasicAtoi2**: Must validate input and handle invalid strings

## Notions

- [strconv/Atoi](https://golang.org/pkg/strconv/#Atoi)