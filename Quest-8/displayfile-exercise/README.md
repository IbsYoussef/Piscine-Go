# displayfile

**Status:** Required

## Problem Statement

Write a program that displays, on the standard output, the content of a file given as argument.

## Expected Output

```console
$ go run .
File name missing
$ echo 'Almost there!!' > quest8.txt
$ go run . quest8.txt main.go
Too many arguments
$ go run . quest8.txt
Almost there!!
```

## Requirements

- Create a folder named `displayfile`
- Create `main.go` inside the folder
- Handle command-line arguments:
  - 0 arguments: print "File name missing"
  - 1 argument: read and display file content
  - 2+ arguments: print "Too many arguments"
- Handle file errors gracefully
- Use fmt, io, and os packages (allowed for this exercise!)

## Submission Structure

```
displayfile/
└── main.go
```

## How to Work on This

1. Write your solution in `student/displayfile/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/displayfile/main.go` if you need help

## Key Concepts

- **Command-line arguments**: os.Args
- **File operations**: os.Open(), io.ReadAll()
- **Error handling**: Check err after file operations
- **defer**: Ensure file.Close() is called
- **Argument validation**: Check len(os.Args)

## Steps

1. Check number of arguments
2. Validate: 0 args → "File name missing"
3. Validate: 2+ args → "Too many arguments"
4. Open file using os.Open()
5. Read content using io.ReadAll()
6. Print content to stdout
7. Handle errors appropriately
