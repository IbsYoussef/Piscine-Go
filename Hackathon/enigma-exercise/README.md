# enigma

## Problem Statement

Write a function called `Enigma` that receives pointers as arguments and move its values around to hide them.

This function will put:

- `a` into `c`.
- `c` into `d`.
- `d` into `b`.
- `b` into `a`.

## Expected Function Signature

```go
func Enigma(a ***int, b *int, c *******int, d ****int) {

}
```

## Expected Output

```console
$ go run .
5
2
7
6
After using Enigma
2
6
5
7
```

## Requirements

- Create a file named `enigma.go`
- Define package as `package piscine`
- Implement the `Enigma` function
- Swap values through multiple levels of pointers
- Handle different pointer indirection levels

## Submission Structure

```
enigma.go
```

## How to Work on This

1. Write your solution in `student/enigma.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/enigma.go` if you need help

## Usage Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	x := 5
	y := &x
	z := &y
	a := &z
	w := 2
	b := &w
	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j
	k := 6
	l := &k
	m := &l
	n := &m
	d := &n
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
	piscine.Enigma(a, b, c, d)
	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}
```

## Key Concepts

- Multiple pointer indirection
- Dereferencing to access values
- Value swapping with different pointer levels
