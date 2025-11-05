# ultimatepointone

**Status:** Required

## Problem Statement

Write a function that takes a **pointer to a pointer to a pointer to an `int`** as argument and gives to this `int` the value of 1.

## Expected Function Signature

```go
func UltimatePointOne(n ***int)
```

## Expected Output

```console
$ go run .
1
$
```

## Requirements

- Create a file named `ultimatepointone.go`
- Define package as `package piscine`
- Implement the `UltimatePointOne` function
- Take a triple pointer to int (`***int`)
- Set the final int value to 1

## Submission Structure

```
ultimatepointone.go
```

## How to Work on This

1. Write your solution in `student/ultimatepointone.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/ultimatepointone.go` if you need help

## Hints

- Triple pointer means: pointer → pointer → pointer → int
- To access the final int, dereference 3 times: `***n`
- Simply assign: `***n = 1`
- Each `*` "follows" one level of pointer

## Example Usage

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
}
```

## Key Concepts

- **Triple pointers**: Pointer to a pointer to a pointer
- **Dereferencing**: Using `*` to access pointed value
- **Multiple levels**: Each `*` goes one level deeper
- **Memory addresses**: Understanding pointer chains

## Visual Explanation

```
&n     →    n     →    b     →    a
(***int)   (**int)   (*int)    (int)

To reach 'a' from '&n': ***(&n) = a
```

## Notions

- [Pointers](https://golang.org/ref/spec#Pointer_types)