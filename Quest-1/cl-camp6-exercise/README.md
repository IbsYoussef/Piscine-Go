# cl-camp6

**Status:** Optional

## Problem Statement

Create a file named `countfiles.sh` that prints the number (and only the number) of regular files and directories contained in the current directory and all its subdirectories. The current directory itself must be included in the count.

## Expected Output Format

```
12
```

(Just a single number - the total count of files and directories)

## Requirements

- Count both regular files and directories
- Include all subdirectories recursively
- Include the current directory in the count
- Output only the number (no extra text)

## How to Work on This

1. Write your solution in `student/countfiles.sh`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/countfiles.sh` if you need help

## Hints

- Use `find` with `-type f` for files and `-type d` for directories
- Use `-or` to combine conditions
- Use `wc -l` to count lines
- Be careful not to double-count special directory entries
