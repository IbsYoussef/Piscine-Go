# join

**Status:** Optional

## Problem Statement

Write a function that returns the concatenation of all the `string`s of a slice of `string`s **separated** by the separator passed as the argument `sep`.

## Expected Function Signature
```go
func Join(strs []string, sep string) string
```

## Expected Output
```console
Hello!: How: are: you?
```

## Requirements

- Create a file named `join.go`
- Define package as `package piscine`
- Implement the `Join` function
- Join all strings in the slice
- Use the separator between elements
- No separator after the last element
- Return empty string if slice is empty

## Submission Structure
```
join.go
```

## How to Work on This

1. Write your solution in `student/join.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/join.go` if you need help

## Hints

- Handle edge cases: empty slice, single element
- Loop through each element with index
- Add separator before each element (except first)
- Use index to check if it's the first element
- Concatenate: element + separator + element

## Example
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
}
```

Output:
```
Hello!: How: are: you?
```

## Key Concepts

- **String joining**: Combining with separator
- **Edge cases**: Empty slice, single element
- **Index checking**: Skip separator for first element
- **Accumulator pattern**: Building result progressively

## Logic Breakdown

- Empty slice → return ""
- Single element → return that element (no separator needed)
- Multiple elements → element + sep + element + sep + ...
- No separator after last element

## Notions

- [strings/Join](https://golang.org/pkg/strings/#Join)