# pilot

## Problem Statement

- Create a directory called `pilot`.
- Inside the directory `pilot` create a file `main.go`.
- Copy the code below to `main.go` and add the code needed so that the program compiles.

> Note: You can only add code, not delete.

## Starting Code

```go
package main

import "fmt"

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1
	fmt.Println(donnie)
}

const AIRCRAFT1 = 1
```

## Expected Output

```console
$ go run .
{Donnie 100 24 1}
```

## Requirements

- Create a folder named `pilot`
- Create `main.go` inside the folder
- Define the `Pilot` struct with appropriate fields
- Only ADD code, do not delete existing code

## Submission Structure

```
pilot/
└── main.go
```

## How to Work on This

1. Write your solution in `student/pilot/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/pilot/main.go` if you need help

## Hints

You need to define a `Pilot` struct with fields:

- `Name` (string)
- `Life` (float64)
- `Age` (int)
- `Aircraft` (appropriate type for the constant)
