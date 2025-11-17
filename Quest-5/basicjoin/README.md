# basicjoin

**Status:** Optional

## Problem Statement

Write a function that returns a concatenated `string` from the 'strings' passed as arguments.

The function takes a slice of strings and joins them together without any separator.

## Expected Function Signature
```go
func BasicJoin(elems []string) string
```

## Expected Output
```console
Hello! How are you?
```

## Requirements

- Create a file named `basicjoin.go`
- Define package as `package piscine`
- Implement the `BasicJoin` function
- Join all strings in the slice together
- No separator between strings
- Return the concatenated result

## Submission Structure
```
basicjoin.go
```

## How to Work on This

1. Write your solution in `student/basicjoin.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/basicjoin.go` if you need help

## Hints

- Initialize empty result string
- Loop through each element in the slice
- Concatenate each element to result
- Use += operator or string concatenation
- Return the final result

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.BasicJoin(elems))
}
```

Output:
```
Hello! How are you?
```

## Key Concepts

- **String concatenation**: Joining strings together
- **Slice iteration**: Looping through elements
- **Accumulator pattern**: Building result progressively
- **Range loop**: Using for range

## Notions

- [string concatenation](https://golang.org/ref/spec#Arithmetic_operators)