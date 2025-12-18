# ztail

**Status:** Bonus

## Problem Statement

Write a program that behaves like a simplified `tail` command that takes at least one file as an argument.

The only option to be handled is `-c` and will be used in all the tests as the first argument, with positive values.

For this program the `os` package can be used.

Handle the errors by returning a non-zero exit status but process all the files.

If several files are given, print a newline and the file name between each one of them.

## Expected Output

If `file1.txt` contains:

```
abcdefghijklmnopqrstuvwxyz
```

**Note** that the file ends with a newline (27 bytes total).

### Normal cases:

```console
$ go run . -c 4 file1.txt
xyz
$ go run . -c 4 file1.txt file2.txt
==> file1.txt <==
xyz

==> file2.txt <==
xyz
```

### Error cases:

```console
$ go run . -c 4 file1.txt nonexisting.txt file2.txt
==> file1.txt <==
xyz
open nonexisting.txt: no such file or directory

==> file2.txt <==
xyz
$ echo $?
1
```

## Requirements

- Create a folder named `ztail`
- Create `main.go` inside the folder
- Handle `-c` flag with positive integer
- Read last N bytes from file(s)
- Handle multiple files with headers
- Continue processing despite errors
- Exit with status 1 if any error occurs

## Submission Structure

```
ztail/
└── main.go
```

## Key Concepts

- **Command-line parsing**: `-c <number> <files...>`
- **File seeking**: file.Seek() to jump to position
- **Byte reading**: Read last N bytes
- **Error handling**: Continue processing, track errors
- **Exit status**: os.Exit(1) on error

## Algorithm

1. Parse arguments
2. Convert number to int (manual)
3. For each file:
   - Get file size
   - Seek to (size - N) position
   - Read N bytes
   - Print content
4. Handle errors, continue processing
5. Exit with 1 if any errors
