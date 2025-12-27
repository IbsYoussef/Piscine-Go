# maxwordcountn

## Problem Statement

Write a function `MaxWordCountN` that will return a `map` of the `n` words that occurs the most in a string `text`. A word is defined as separated by spaces. The `map` you should return will have the word as key and the number of occurences of this word as value.

If two words have the same number of occurences, the one with the lowest ASCII value should be prioritized.

## Expected Function Signature

```go
func MaxWordCountN(text string, n int) map[string]int {

}
```

## Expected Output

```console
$ go run .
map[Orange:6 of:7 the:8]
```

## Requirements

- Create a file named `maxwordcountn.go`
- Define package as `package piscine`
- Implement the `MaxWordCountN` function
- Return top N most frequent words
- If counts are equal, prioritize lower ASCII value
- Words are separated by spaces, newlines, tabs

## Submission Structure

```
maxwordcountn.go
```

## How to Work on This

1. Write your solution in `student/maxwordcountn.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/maxwordcountn.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.MaxWordCountN(`
	Orange is the sun sliding to the horizon after a summer day. Orange is the sound of dribbling basketball. Orange is the smell of a tiger lily petal. Orange is the taste of thirst-quenching Nehi Soda. Orange is the color of peach marmalade on a side of toast. Orange is the sound of a carrot popping out of the ground.
	`, 3))
}
```
