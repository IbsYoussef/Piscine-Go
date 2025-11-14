# index

**Status:** Required

## Problem Statement

Write a function that behaves like the `Index` function from the strings package.

The function returns the index of the first occurrence of a substring in a string. If the substring is not found, it returns -1.

## Expected Function Signature

```go
func Index(s string, toFind string) int
```

## Expected Output

```console
2
1
-1
```

## Requirements

- Create a file named `index.go`
- Define package as `package piscine`
- Implement the `Index` function
- Return the index of first occurrence of toFind in s
- Return -1 if toFind is not found
- Return 0 if toFind is empty string

## Submission Structure

```
index.go
```

## How to Work on This

1. Write your solution in `student/index.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/index.go` if you need help

## Hints

- Handle edge cases: empty toFind, toFind longer than s
- Loop through s, checking substrings
- Use string slicing: s[i:i+len(toFind)]
- Compare substring with toFind
- Return index on first match
- Return -1 if no match found

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))     // 2 (first 'l')
	fmt.Println(piscine.Index("Salut!", "alu"))   // 1 (substring at index 1)
	fmt.Println(piscine.Index("Ola!", "hOl"))     // -1 (not found)
}
```

Output:

```
2
1
-1
```

## Key Concepts

- **Substring search**: Finding pattern in text
- **String slicing**: s[start:end]
- **Index positioning**: 0-based indexing
- **Edge cases**: Empty strings, not found

## Edge Cases

- Empty toFind → return 0
- toFind longer than s → return -1
- Empty s → return -1
- toFind not in s → return -1

## Notions

- [strings/Index](https://golang.org/pkg/strings/#Index)
