# atoibase

**Status:** Bonus

## Problem Statement

Write a function that takes two arguments:
- `s`: a numeric `string` in a given base
- `base`: a `string` representing all the different digits that can represent a numeric value

And return the integer value of `s` in the given `base`.

If the base is not valid it returns `0`.

## Validity Rules for a Base

- A base must contain at least 2 characters
- Each character of a base must be unique
- A base should not contain `+` or `-` characters
- String number must contain only elements that are in base

Only valid `string` numbers will be tested.
The function **does not have** to manage negative numbers.

## Expected Function Signature
```go
func AtoiBase(s string, base string) int
```

## Expected Output
```console
125
125
125
125
0
```

## Requirements

- Create a file named `atoibase.go`
- Define package as `package piscine`
- Implement the `AtoiBase` function
- Validate the base
- Check if all characters in s exist in base
- Convert string to integer using base conversion
- Return 0 if invalid

## Submission Structure
```
atoibase.go
```

## How to Work on This

1. Write your solution in `student/atoibase.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/atoibase.go` if you need help

## Hints

- This is the reverse of PrintNbrBase
- Validate base (same rules)
- Check all characters in s are in base
- Use Horner's method: result = result * baseLen + digitValue
- Find digit value by finding index of character in base string
- Process characters left to right

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AtoiBase("125", "0123456789"))           // 125
	fmt.Println(piscine.AtoiBase("1111101", "01"))               // 125 (binary)
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))      // 125 (hex)
	fmt.Println(piscine.AtoiBase("uoi", "choumi"))               // 125 (custom)
	fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))              // 0 (invalid)
}
```

## Conversion Formula

For each character from left to right:
```
result = result * base_length + digit_value
```

Where digit_value is the index of the character in the base string.

## Key Concepts

- **Horner's method**: Efficient polynomial evaluation
- **Base conversion**: String to integer
- **Character lookup**: Finding position in base
- **Base validation**: Same as PrintNbrBase