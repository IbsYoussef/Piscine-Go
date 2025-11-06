# eightqueens

**Status:** Bonus

## Problem Statement

Write a function that prints the solutions to the [eight queens puzzle](https://en.wikipedia.org/wiki/Eight_queens_puzzle).

Recursivity must be used to solve this problem.

It should print something like this:

```console
$ go run .
15863724
16837425
17468253
17648253
...
(92 solutions total)
```

Each solution will be on a single line.
The index of the placement of a queen starts at 1.
It reads from left to right and each digit is the position for each column.
The solutions will be printed in ascending order.

## Expected Function Signature

```go
func EightQueens()
```

## Requirements

- Create a file named `eightqueens.go`
- Define package as `package piscine`
- Implement the `EightQueens` function
- Use recursion to solve the problem
- Print all 92 solutions
- Each solution on a single line
- Solutions in ascending order
- Format: 8 digits representing queen positions per column

## Submission Structure

```
eightqueens.go
```

## How to Work on This

1. Write your solution in `student/eightqueens.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/eightqueens.go` if you need help

## Hints

- Use backtracking algorithm
- For each column, try placing a queen in each row
- Check if placement is valid (no conflicts)
- If valid, move to next column recursively
- If all 8 queens placed successfully, print solution
- Backtrack and try next position
- A queen attacks horizontally, vertically, and diagonally
- Need to check: same row, same diagonal (both directions)

## Understanding the Output Format

Each line represents one solution:

- 8 digits (one per column)
- Each digit is the row position (1-8) of the queen in that column
- Example: `15863724` means:
  - Column 1: Queen at row 1
  - Column 2: Queen at row 5
  - Column 3: Queen at row 8
  - Column 4: Queen at row 6
  - Column 5: Queen at row 3
  - Column 6: Queen at row 7
  - Column 7: Queen at row 2
  - Column 8: Queen at row 4

## Key Concepts

- **Backtracking**: Try, validate, recurse, or backtrack
- **Recursion**: Solving smaller subproblems
- **Constraint satisfaction**: Checking valid placements
- **Eight queens puzzle**: Classic CS problem
- **Combinatorial search**: Exploring all possibilities

## Algorithm Approach

1. Start with column 0
2. For each row in current column:
   - Check if safe to place queen
   - If safe, place queen and recurse to next column
   - If recursion returns (backtrack), remove queen and try next row
3. When all 8 columns filled, print solution
4. Continue until all solutions found

## Validation Rules

A queen placement is safe if no other queen can attack it:

- No queen in same row
- No queen in diagonal (top-left to bottom-right)
- No queen in diagonal (top-right to bottom-left)

## Expected Output

Total of 92 solutions, starting with:

```
15863724
16837425
17468253
17648253
18426735
...
```

## Resources

- [Eight Queens Puzzle - Wikipedia](https://en.wikipedia.org/wiki/Eight_queens_puzzle)
- Study backtracking algorithms
- Understand recursion patterns
