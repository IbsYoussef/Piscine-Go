# concat

**Status:** Required

## Problem Statement

Write a function that returns the concatenation of two `string` passed in arguments.

Concatenation means joining two strings together to form one string.

## Expected Function Signature
```go
func Concat(str1 string, str2 string) string
```

## Expected Output
```console
Hello! How are you?
```

## Requirements

- Create a file named `concat.go`
- Define package as `package piscine`
- Implement the `Concat` function
- Join str1 and str2 together
- Return the combined string

## Submission Structure
```
concat.go
```

## How to Work on This

1. Write your solution in `student/concat.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/concat.go` if you need help

## Hints

- Use the + operator to concatenate strings
- Simply: `return str1 + str2`
- This is a very simple exercise!

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Concat("Hello!", " How are you?"))
}
```

Output:
```
Hello! How are you?
```

## Key Concepts

- **String concatenation**: Joining strings together
- **+ operator**: Used for string concatenation in Go
- **String immutability**: Creates new string (doesn't modify originals)

## Examples

- Concat("Hello", "World") → "HelloWorld"
- Concat("Go", "Lang") → "GoLang"
- Concat("", "test") → "test"
- Concat("test", "") → "test"