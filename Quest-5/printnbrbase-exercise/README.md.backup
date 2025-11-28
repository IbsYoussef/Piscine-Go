# printnbrbase

**Status:** Bonus

## Problem Statement

Write a function that prints an `int` in a `string` base passed as parameters.

If the base is not valid, the function prints `NV` (Not Valid).

## Validity Rules for a Base

- A base must contain at least 2 characters
- Each character of a base must be unique
- A base should not contain `+` or `-` characters

The function has to manage negative numbers.

## Expected Function Signature
```go
func PrintNbrBase(nbr int, base string)
```

## Expected Output
```console
125
-1111101
7D
-uoi
NV
```

## Requirements

- Create a file named `printnbrbase.go`
- Define package as `package piscine`
- Import `"github.com/01-edu/z01"`
- Implement the `PrintNbrBase` function
- Validate the base
- Convert number to the given base
- Handle negative numbers
- Print result or "NV"

## Submission Structure
```
printnbrbase.go
```

## How to Work on This

1. Write your solution in `student/printnbrbase.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/printnbrbase.go` if you need help

## Hints

- Create helper function to validate base
- Check base length >= 2
- Check for duplicate characters
- Check for '+' or '-' in base
- For conversion: use modulo and division
- Handle negative numbers: print '-', then convert absolute value
- Build result in reverse, then print forwards

## Example
```go
package main

import (
	"piscine"
	"github.com/01-edu/z01"
)

func main() {
	piscine.PrintNbrBase(125, "0123456789")        // 125 (decimal)
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")               // -1111101 (binary)
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")  // 7D (hexadecimal)
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")           // -uoi (custom base)
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa")                // NV (duplicate 'a')
	z01.PrintRune('\n')
}
```

## Base Conversion Logic

Binary (base 2): "01"
- 125 ÷ 2 = 62 remainder 1
- 62 ÷ 2 = 31 remainder 0
- 31 ÷ 2 = 15 remainder 1
- Continue... Result: 1111101

Hexadecimal (base 16): "0123456789ABCDEF"
- 125 ÷ 16 = 7 remainder 13
- 13 maps to 'D', 7 maps to '7'
- Result: 7D

## Key Concepts

- **Base conversion**: Converting numbers to different bases
- **Base validation**: Checking validity rules
- **Modulo arithmetic**: Finding remainders
- **Digit extraction**: Building number representation

## Notions

- [01-edu/z01](https://github.com/01-edu/z01)