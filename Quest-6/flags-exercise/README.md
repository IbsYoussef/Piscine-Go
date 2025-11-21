# flags

**Status:** Optional

## Problem Statement

Write a **program** that can take `--insert` (or `-i`), `--order` (or `-o`) and a `string` as arguments.

This program should:

- Insert the string given to `--insert` (or `-i`) into the argument string
- If flag `--order` (or `-o`) is given, order the string argument (ASCII order)
- If no arguments or `--help` (or `-h`) is given, print the help message

## Expected Output

```console
$ go run . --insert=4321 --order asdad
1234aadds

$ go run . --insert=4321 asdad
asdad4321

$ go run . asdad
asdad

$ go run . --order 43a21
1234a

$ go run .
--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.

$ go run . -h
--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.
```

## Requirements

- Create a folder named `flags`
- Create `main.go` inside the folder
- Handle multiple flag formats: `--insert=value`, `-i=value`
- Handle boolean flag: `--order`, `-o`
- Print help when no args or `--help`/`-h`
- Combine operations when both flags present

## Submission Structure

```
flags/
└── main.go
```

## How to Work on This

1. Write your solution in `student/flags/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/flags/main.go` when ready

## Key Concepts

- **Flag parsing**: Detecting and extracting flag values
- **String manipulation**: Parsing `--insert=value`
- **Boolean flags**: `--order` doesn't need a value
- **Combining operations**: Insert then order
- **Help message**: Formatted output

## Hints

- Parse arguments to detect flags
- Extract value from `--insert=value` format
- Track which flags are set
- Process the final string argument
- Order of operations: insert first, then order if needed
