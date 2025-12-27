# unmatch

## Problem Statement

Write a function, `Unmatch`, that returns the element of the slice that does not have a correspondent pair.

- If all the number have a correspondent pair, it should return `-1`.

## Expected Function Signature

```go
func Unmatch(a []int) int {

}
```

## Expected Output

```console
$ go run .
4
```

## Requirements

- Create a file named `unmatch.go`
- Define package as `package piscine`
- Implement the `Unmatch` function
- Return the unpaired number
- Return -1 if all numbers are paired

## Submission Structure

```
unmatch.go
```

## How to Work on This

1. Write your solution in `student/unmatch.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/unmatch.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := piscine.Unmatch(a)
	fmt.Println(unmatch)
}
```

## Examples

```
[1, 2, 3, 1, 2, 3, 4] → 4 (4 appears once, unpaired)
[5, 5, 7, 7] → -1 (all paired)
[1, 1, 1] → 1 (1 appears 3 times, odd = unpaired)
```
