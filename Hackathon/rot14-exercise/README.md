# rot14

## Problem Statement

Write a function rot14 that returns the string within the parameter transformed into a rot14 string. Each letter will be replaced by the letter 14 spots ahead in the alphabetical order.

'z' becomes 'n' and 'Z' becomes 'N'. The case of the letter stays the same.

## Expected Function Signature

```go
func Rot14(s string) string {

}
```

## Expected Output

```console
$ go run .
Vszzc! Vck ofs Mci?
```

## Requirements

- Create a file named `rot14.go`
- Define package as `package piscine`
- Implement the `Rot14` function
- Rotate letters by 14 positions
- Preserve case (uppercase stays uppercase)
- Non-letters remain unchanged

## Submission Structure

```
rot14.go
```

## How to Work on This

1. Write your solution in `student/rot14.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/rot14.go` if you need help

## Usage Example

```go
package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	result := piscine.Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
```
