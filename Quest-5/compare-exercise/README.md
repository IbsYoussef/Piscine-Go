# compare

**Status:** Required

## Problem Statement

Write a function that behaves like the `Compare` function from the strings package.

The function compares two strings lexicographically (alphabetically) and returns:

- `0` if a == b
- `-1` if a < b
- `1` if a > b

## Expected Function Signature

```go
func Compare(a, b string) int
```

## Expected Output

```console
0
-1
1
```

## Requirements

- Create a file named `compare.go`
- Define package as `package piscine`
- Implement the `Compare` function
- Return 0 if strings are equal
- Return -1 if a comes before b (alphabetically)
- Return 1 if a comes after b (alphabetically)

## Submission Structure

```
compare.go
```

## How to Work on This

1. Write your solution in `student/compare.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/compare.go` if you need help

## Hints

- Compare strings character by character
- Convert to rune slices for proper Unicode handling
- If characters differ, return based on comparison
- If one string is prefix of another, shorter one is "less"
- Handle empty strings as edge cases

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "Hello!"))  // 0 (equal)
	fmt.Println(piscine.Compare("Salut!", "lut!"))    // -1 (S < l)
	fmt.Println(piscine.Compare("Ola!", "Ol"))        // 1 (longer)
}
```

Output:

```
0
-1
1
```

## Key Concepts

- **Lexicographic comparison**: Dictionary order
- **Character comparison**: Using Unicode values
- **String length**: When one is prefix of another
- **Edge cases**: Empty strings

## Comparison Logic

1. Compare character by character
2. First difference determines result
3. If all match but different lengths, longer > shorter
4. If completely equal, return 0

## Notions

- [strings/Compare](https://golang.org/pkg/strings/#Compare)
