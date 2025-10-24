# who-are-you

**Status:** Required

## Problem Statement

You just woke up in a dark alley with the tag "subject Id: 70".

Your task is to fetch the superhero list and print the name corresponding to the entry whose id is 70.

**URL:** `https://learn.01founders.co/assets/superhero/all.json`

## Expected Output

```
"Batman"
```

(The output includes the quotes as it's a JSON string)

## Requirements

- Use `curl` to download the JSON data
- Use `jq` to parse and extract the name
- Make sure `jq` is installed on your system

## How to Work on This

1. Write your solution in `student/who-are-you.sh`
2. Run `./run.sh` to see your output
3. Run `./test.sh` to check if it works
4. Compare with `solutions/who-are-you.sh` if you need help
