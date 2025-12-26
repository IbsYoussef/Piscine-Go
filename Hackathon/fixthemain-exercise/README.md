# fixthemain

## Problem Statement

Fix the following program.

## Program to Fix

```go
package piscine

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	state = CLOSE
	return true
}

func IsDoorOpen(Door Door) {
	PrintStr("is the Door opened ?")
	return Door.state = OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
```

## Expected Output

```console
$ go run .
Door Opening...
is the Door closed ?
is the Door opened ?
Door Closing...
```

## Requirements

- Create a folder named `fixthemain`
- Create `main.go` inside the folder
- Fix all compilation errors
- Fix all logic errors
- Ensure proper output

## Submission Structure

```
fixthemain/
└── main.go
```

## How to Work on This

1. Write your solution in `student/fixthemain/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/fixthemain/main.go` if you need help

## Issues to Fix

The original code has multiple errors:

- Missing constants
- Missing struct definition
- Missing function
- Wrong package name
- Wrong parameter types
- Assignment instead of comparison
- Missing return statements
- Undefined variables
