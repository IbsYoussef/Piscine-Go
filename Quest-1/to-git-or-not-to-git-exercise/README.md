# to-git-or-not-to-git

**Status:** Required

## Problem Statement

Create a file named `to-git-or-not-to-git.sh` that retrieves and displays the name, power, and gender of the superhero whose id is 170.

## Expected Output Format
```
Chameleon
28
Male
```

(Three lines: name, power level, gender)

## Requirements

- Fetch superhero data from JSON API
- Filter for superhero with id = 170
- Extract and display: name, power, gender
- Each field on a separate line
- **No `echo` commands allowed** (must use actual API calls)

## API Details

- **Endpoint:** `https://learn.01founders.co/assets/superhero/all.json`
- **Method:** HTTP GET (returns JSON array)
- **Fields needed:** name (from root), power (from powerstats object), gender (from appearance object)
- **Filter:** Select where `id == 170`

## How to Work on This

1. Write your solution in `student/to-git-or-not-to-git.sh`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/to-git-or-not-to-git.sh` if you need help

## Hints

- Use `curl -s` for silent HTTP requests
- Use `jq` to parse JSON response
- Filter with `select(.id == 170)`
- Access nested fields: `.powerstats.power` and `.appearance.gender`
- Use `jq -r` for raw output (no quotes)