# boolean

**Status:** Required

## Problem Statement

Create a new directory called `boolean`.

- The code below must be copied into a file called `main.go` inside of the `boolean` directory.
- The necessary changes must be applied for the program to work.

## Code to be copied

```go
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
```

## Expected Output

```console
$ go run . "not" "odd"
I have an even number of arguments
$ go run . "not even"
I have an odd number of arguments
```

## Requirements

- Create a folder named `boolean`
- Create `main.go` inside the folder
- Fix the broken code to make it work
- Program checks if number of arguments is even or odd
- Print appropriate message

## Submission Structure

```
boolean/
└── main.go
```

## How to Work on This

1. Copy the broken code into `student/boolean/main.go`
2. Fix the code to make it work
3. Run `./run.sh` to see your output
4. Run `./test.sh` to check if it works
5. Compare with `solutions/boolean/main.go` if you need help

## Key Concepts

- **Bool type in Go**: Use `bool` not `boolean`
- **Boolean values**: Use `true` and `false`
- **Command-line arguments**: Use `os.Args`
- **Argument counting**: `len(os.Args) - 1` (exclude program name)
- **String constants**: Define message strings

## What's Wrong with the Code

The provided code has several issues:

1. `boolean` is not a valid Go type (should be `bool`)
2. `yes` and `no` are not defined (should be `true` and `false`)
3. `even()` function is not defined
4. `lengthOfArg` is not defined (should calculate from `os.Args`)
5. `EvenMsg` and `OddMsg` constants are not defined
6. Comparing booleans to `1` is incorrect
