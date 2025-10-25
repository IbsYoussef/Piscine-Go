# cl-camp2

**Status:** Required

## Problem Statement

Create a file named `r` that shows the letter "R" on a line when viewed using the `cat` command.

## Expected Output

```
R
```

When using `cat -e` (which shows line endings with `$`):

```
R$
```

## Requirements

- File must contain the letter "R" followed by a newline
- When you run `cat student/r`, it should display: `R`
- The file should have exactly one line with one character

## How to Work on This

1. Create your file in `student/r`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/r` if you need help

## Hint

This exercise is about understanding how text files work with newlines. You can create the file using `echo "R" > student/r` or by editing it directly.
