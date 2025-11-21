# rotatevowels

**Status:** Bonus

## Problem Statement

Write a **program** that checks the arguments for vowels and **"mirrors"** their positions.

- Find all vowels across all arguments
- Reverse the order of vowels
- Put reversed vowels back in their original positions
- For this exercise, `y` is NOT considered a vowel
- If less than 1 argument: print newline
- If no vowels: print arguments unchanged

## Expected Output

```console
$ go run . "Hello World"
Hollo Werld

$ go run . "HEllO World" "problem solved"
Hello Werld problom sOlvEd

$ go run . "str" "shh" "psst"
str shh psst

$ go run . "happy thoughts" "good luck"
huppy thooghts guod lack

$ go run . "aEi" "Ou"
uOi Ea

$ go run .
```

## Requirements

- Create a folder named `rotatevowels`
- Create `main.go` inside the folder
- Vowels: `a, e, i, o, u` (uppercase and lowercase)
- `y` is NOT a vowel
- Mirror vowels across ALL arguments
- Print arguments separated by spaces

## Submission Structure

```
rotatevowels/
└── main.go
```

## How to Work on This

1. Write your solution in `student/rotatevowels/main.go`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/rotatevowels/main.go` when ready

## Key Concepts

- **Vowel detection**: Checking if character is a vowel
- **Collection**: Gathering all vowels from all arguments
- **Reversal**: Reversing the vowel list
- **Replacement**: Putting reversed vowels back

## Hints

- Collect all vowels from all arguments first
- Reverse the vowel list
- Loop through arguments again, replacing vowels
- Track which reversed vowel to use next
