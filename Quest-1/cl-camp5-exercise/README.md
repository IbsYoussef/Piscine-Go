# cl-camp5

**Status:** Optional

## Problem Statement

Create a file named `lookagain.sh` that searches from the current directory and all its subdirectories for all files ending with `.sh`.

The output must:

- Display only the file names (without the `.sh` extension)
- Be sorted in descending (reverse alphabetical) order

## Expected Output Format

```
test
script
hello
example
```

(Sorted in reverse alphabetical order, no `.sh` extensions)

## Requirements

- Search recursively through all subdirectories
- Find all files ending with `.sh`
- Remove the `.sh` extension from output
- Sort results in reverse alphabetical order

## How to Work on This

1. Write your solution in `student/lookagain.sh`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/lookagain.sh` if you need help

## Hints

- Use `find` command with `-name` pattern
- Use `basename` to remove path and extension
- Use `sort -r` for reverse sorting
- Use `-exec` with find to process each file
