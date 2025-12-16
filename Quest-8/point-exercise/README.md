# point

**Status:** Required

## Problem Statement

Create a new directory called `point`.

- The code below must be copied into a file called `main.go` inside the `point` directory.
- The necessary changes must be applied so that the program works.
- The function `setPoint()` must work with `int`.

## Code to be copied

```go
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}
```

## Expected Output

```console
$ go run .
x = 42, y = 21
```

## Requirements

- Create a folder named `point`
- Create `main.go` inside the folder
- Define a `point` struct with `x` and `y` fields (both `int`)
- Implement `setPoint()` to set x=42, y=21
- Print the values using z01.PrintRune (no fmt package allowed)
- Must use pointer to modify the struct

## Submission Structure

```
point/
└── main.go
```

## How to Work on This

1. Write your solution in `student/point/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/point/main.go` if you need help

## Key Concepts

- **Structs**: Defining custom types
- **Pointers**: Modifying struct through pointer
- **z01.PrintRune**: Printing without fmt package
- **Number printing**: Converting int to digits manually

## What You Need to Add

1. Define `point` struct with x and y fields (both int)
2. Import necessary packages (github.com/01-edu/z01)
3. Create helper functions to print strings and numbers
4. Format output to match: "x = 42, y = 21\n"
