# isprime

**Status:** Optional

## Problem Statement

Write a function that returns `true` if the `int` passed as parameter is a prime number. Otherwise it returns `false`.

The function must be optimized in order to avoid time-outs with the tester.

(We consider that only positive numbers can be prime numbers)

(We also consider that 1 is **not** a prime number)

## Expected Function Signature

```go
func IsPrime(nb int) bool
```

## Expected Output

```console
true
false
```

## Requirements

- Create a file named `isprime.go`
- Define package as `package piscine`
- Implement the `IsPrime` function
- Return true if number is prime
- Return false if not prime
- Must be optimized (don't check all numbers up to nb)
- Only positive numbers can be prime
- 1 is NOT a prime number

## Submission Structure

```
isprime.go
```

## How to Work on This

1. Write your solution in `student/isprime.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/isprime.go` if you need help

## Hints

- Numbers ≤ 1 are not prime
- Only check divisors up to √nb (optimization)
- Calculate sqrt: loop while sqrt\*sqrt <= nb
- Check if divisible by any number from 2 to sqrt-1
- If divisible, return false
- If loop completes, return true

## Example

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrime(5))
	fmt.Println(piscine.IsPrime(4))
}
```

Output:

```
true
false
```

## Key Concepts

- **Prime number**: Only divisible by 1 and itself
- **Optimization**: Only check up to √n
- **Divisibility**: Check if nb % i == 0
- **Early return**: Return false on first divisor found

## Prime Numbers

First few primes: 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37...

Not prime:

- 1 (by definition)
- 4 (divisible by 2)
- 6 (divisible by 2, 3)
- 8 (divisible by 2, 4)
- 9 (divisible by 3)
