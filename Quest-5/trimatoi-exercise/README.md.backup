# trimatoi

**Status:** Required

## Problem Statement

Write a function that transforms numbers within a `string`, into an `int`.

- Extract all digits from the string and combine them into a number
- If the `-` sign is encountered **before any number**, it determines the sign of the returned `int`
- This function should **only** return an `int`
- In the case of an invalid input (no digits), return `0`
- **Note**: There will never be more than one sign in a `string` in the tests

## Expected Function Signature
```go
func TrimAtoi(s string) int
```

## Expected Output
```console
12345
12345
12345
0
1234
-1234
1234
1234
```

## Requirements

- Create a file named `trimatoi.go`
- Define package as `package piscine`
- Implement the `TrimAtoi` function
- Extract all digits from string
- Combine digits into an integer
- Handle negative sign if it appears before any digit
- Ignore all non-digit characters (except first `-`)
- Return 0 if no digits found

## Submission Structure
```
trimatoi.go
```

## How to Work on This

1. Write your solution in `student/trimatoi.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/trimatoi.go` if you need help

## Hints

- Initialize result = 0 and sign = 1
- Track if you've found a digit yet (for sign handling)
- Loop through each character
- If `-` found before any digit: set sign = -1
- If digit found: extract it and add to result
- Build number: result = result * 10 + digit
- Convert rune to digit: int(rune - '0')
- Return result * sign

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.TrimAtoi("12345"))           // 12345
	fmt.Println(piscine.TrimAtoi("str123ing45"))     // 12345
	fmt.Println(piscine.TrimAtoi("012 345"))         // 12345
	fmt.Println(piscine.TrimAtoi("Hello World!"))    // 0
	fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))    // 1234
	fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))    // -1234
	fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))    // 1234 (- after digit)
	fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))    // 1234 (+ ignored)
}
```

## Key Concepts

- **String parsing**: Extracting digits from mixed content
- **Number building**: Combining digits into integer
- **Sign handling**: Negative numbers
- **Character classification**: Identifying digits
- **State tracking**: Has digit been found yet?

## Important Rules

- Only the first `-` before any digit determines sign
- All digits are extracted and combined
- Non-digit characters (except first `-`) are ignored
- Spaces, letters, symbols between digits are ignored
- `+` sign is always ignored