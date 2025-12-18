# cat

**Status:** Optional

## Problem Statement

Write a program that behaves like a simplified `cat` command.

- The options do not have to be handled.
- If the program is called without arguments, it should take the standard input (stdin) and print it back on the standard output (stdout).
- If files are provided, print their contents.
- Handle errors for missing files but continue processing other files.

## Expected Output

```console
$ echo '"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing' > quest8.txt
$ go run . abc
ERROR: open abc: no such file or directory
exit status 1
$ go run . quest8.txt
"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing
$ go run . quest8.txt abc
"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing
ERROR: open abc: no such file or directory
$ cat quest8.txt | go run .
"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing
$ go run .
Hello
Hello
^C
$ go run . quest8.txt quest8T.txt
"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing
"Alan Mathison Turing was an English mathematician, computer scientist..."
```

## Requirements

- Create a folder named `cat`
- Create `main.go` inside the folder
- Handle two modes:
  - **No arguments:** Read from stdin, echo to stdout
  - **With files:** Read and print each file
- Handle file errors gracefully
- Exit with status 1 if any errors occur

## Submission Structure

```
cat/
└── main.go
```
