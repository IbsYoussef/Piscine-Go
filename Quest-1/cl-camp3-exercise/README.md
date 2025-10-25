# cl-camp3

**Status:** Required

## Problem Statement

Create a file named `look` that searches within the current directory and all its subdirectories for:

- Files or directories that start with 'a'
- Files that end with 'z'
- Files that start with 'z' **AND** end with 'a!'

## Expected Output Format

```
./apple
./amazing.txt
./buzz
./zebra-a!
```

(Actual output will vary based on directory contents)

## Requirements

- Search recursively through all subdirectories
- Match files/directories starting with 'a'
- Match files ending with 'z'
- Match files that start with 'z' AND end with 'a!'
- Use `find` command with pattern matching

## How to Work on This

1. Write your solution in `student/look`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/look` if you need help

## Hints

- Use `find` command with `-name` for pattern matching
- Use `-o` for OR logic
- Use `-a` for AND logic
- Patterns: `a*` (starts with a), `*z` (ends with z)
