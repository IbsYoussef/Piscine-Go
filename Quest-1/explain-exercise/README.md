# explain

**Status:** Optional

## Problem Statement

Create a file named `explain.sh` that documents how you solved the command-line murder mystery. The script must output specific details about your investigation in order.

## Expected Output Format
```
Annabel Church
699607
Blue Honda
Aldo Nicolas
Andrei Masna
Matt Waite
```

(Six lines: witness name, interview number, car description, three other suspects)

## Requirements

- Output the first and last name of your key witness
- Output the interview number of this witness
- Output the color and make of the suspect's car
- Output the names of three other main suspects (not arrested)
- Names must be sorted alphabetically by last name
- Each piece of information on a separate line

## How to Work on This

1. Write your solution in `student/explain.sh`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/explain.sh` if you need help

## Hints

- The key witness is the person mentioned by the barista
- The interview number can be found by following the witness's address
- The car details come from the witness's interview
- Other suspects are tall males with all four memberships
- Exclude the arrested suspect (Joe Germuska) from the list
- Sort remaining suspects by last name