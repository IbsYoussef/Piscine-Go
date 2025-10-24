# cl-camp1

**Status:** Required

## Problem Statement

Create a file named `mastertheLS` that lists files and directories in the current directory while following these rules:

- Ignore hidden files ('.' and '..')
- Separate results with commas
- Sort by access time, newest first
- Append '/' to directory names

## Expected Output Format

```
file1.txt,dir1/,file2.go,script.sh
```

(Actual output will vary based on directory contents)

## Requirements

- No hidden files in output
- Comma-separated list
- Sorted by time (newest first)
- Directories marked with `/`

## How to Work on This

1. Write your solution in `student/mastertheLS`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/mastertheLS` if you need help

## Notes

The test will run your script in different directories (`cl-camp1/folder1` and `cl-camp1/folder2`) to verify it works correctly.
