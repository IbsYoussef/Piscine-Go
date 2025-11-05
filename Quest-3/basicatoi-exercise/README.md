# basicatoi

**Status:** Optional

## Problem Statement

Write a function that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number defined as a `string` into a number defined as an `int`.

For this exercise **only valid** `string` will be tested. They will only contain one or several digits as characters. The handling of the signs `+` or `-` does not have to be taken into account.

## Expected Function Signature

```go
func BasicAtoi(s string) int
```

## Expected Output

```console
$ go run .
12345
12345
0
$
```

## Requirements

- Create a file named `basicatoi.go`
- Define package as `package piscine`
- Implement the `BasicAtoi` function
- Convert string of digits to integer
- Only valid digit strings will be tested
- No sign handling needed (`+` or `-`)
- Return `0` for empty strings

## Submission Structure

```
basicatoi.go
```

## How to Work on This

1. Write your solution in `student/basicatoi.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/basicatoi.go` if you need help

## Hints

- Initialize result to 0
- Iterate through each character
- Convert character to digit: `int(char - '0')`
- Build number: `result = result * 10 + digit`
- ASCII: '0' is 48, '1' is 49, etc.
- Subtracting '0' (or 48) gives the digit value

## Example Usage

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BasicAtoi("12345"))
	fmt.Println(piscine.BasicAtoi("0000000012345"))
	fmt.Println(piscine.BasicAtoi("000000"))
}
```

## Key Concepts

- **ASCII values**: Characters are stored as numbers
- **String to int conversion**: Building a number digit by digit
- **Base 10 arithmetic**: Multiplying by 10 to shift digits left
- **Character arithmetic**: Subtracting '0' to get digit value

## How It Works

For "123":
1. Start: result = 0
2. '1': result = 0 * 10 + 1 = 1
3. '2': result = 1 * 10 + 2 = 12
4. '3': result = 12 * 10 + 3 = 123

## Notions

- [strconv/Atoi](https://golang.org/pkg/strconv/#Atoi)