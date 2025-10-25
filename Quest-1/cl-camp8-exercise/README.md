# cl-camp8

**Status:** Optional

## Problem Statement

Create a file named `skip.sh` that prints the result of `ls -l` while skipping every other line, starting with the first one.

## Expected Output Format

```
total 8
-rw-r--r--  1 user  staff   234 Jan 21 10:00 file2.txt
-rw-r--r--  1 user  staff   512 Jan 21 10:02 file4.txt
```

(Shows only even-numbered lines from `ls -l` output)

## Requirements

- Run `ls -l` command
- Skip every other line starting with line 1
- Output only the even-numbered lines (2nd, 4th, 6th, etc.)

## How to Work on This

1. Write your solution in `student/skip.sh`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/skip.sh` if you need help

## Hints

- Use `ls -l` to list files in long format
- Use `sed` with line deletion patterns
- Pattern `1~2d` means "start at line 1, every 2nd line after"
- Alternative: `awk` with modulo operator
