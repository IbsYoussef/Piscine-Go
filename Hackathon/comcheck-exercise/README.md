# comcheck

## Problem Statement

Write a program `comcheck` that displays on the standard output `Alert!!!` followed by newline (`'\n'`) if at least one of the arguments passed in parameter matches the `string`:

- `01`, `galaxy` or `galaxy 01`.
- If none of the parameters match, the program displays nothing.

## Expected Output

```console
$ go run . "I" "Will" "Enter" "the" "galaxy"
Alert!!!
$ go run . "galaxy 01" "do" "you" "hear" "me"
Alert!!!
```

## Requirements

- Create a folder named `comcheck`
- Create `main.go` inside the folder
- Check if any argument matches: "01", "galaxy", or "galaxy 01"
- Print "Alert!!!" if match found
- Print nothing if no match

## Submission Structure

```
comcheck/
└── main.go
```

## How to Work on This

1. Write your solution in `student/comcheck/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/comcheck/main.go` if you need help
