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

- Use GraphQL API to fetch superhero data
- Query for superhero with id = 170
- Extract and display: name, power, gender
- Each field on a separate line
- **No `echo` commands allowed** (must use actual API calls)

## API Details

- **Endpoint:** `https://learn.01founders.co/api/graphql-engine/v1/graphql`
- **Method:** GraphQL query
- **Fields needed:** name, power, gender
- **Filter:** `id: 170`

## How to Work on This

1. Write your solution in `student/to-git-or-not-to-git.sh`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/to-git-or-not-to-git.sh` if you need help

## Hints

- Use `curl -s` for silent HTTP requests
- Use `--data` to send GraphQL query
- Use `jq` to parse JSON response
- GraphQL query format: `{superhero(where:{id:{_eq:170}}){name power gender}}`
