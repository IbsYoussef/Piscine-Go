#!/bin/sh
# Exercise: to-git-or-not-to-git
#
# Required:
# Create a file named 'to-git-or-not-to-git.sh' that retrieves and displays
# the name, power, and gender of the superhero whose id is 170.
#
# Explanation:
# - 'curl -s': sends a silent HTTP request (no progress or error messages)
#   to the provided GraphQL API endpoint.
# - The '--data' flag sends a GraphQL query payload to the API.
# - The query retrieves the 'user' field where the 'login' equals the given username.
# - 'jq': a lightweight JSON processor used to parse and extract the relevant fields
#   from the returned JSON structure.
#
# However, based on the exercise description, the correct logic should instead
# query the 'superhero' dataset to fetch:
#   - name
#   - power
#   - gender
# for the entry with 'id' equal to 170.
#
# Example GraphQL query structure (conceptually):
# {
#   superhero(where: {id: {_eq: 170}}) {
#     name
#     power
#     gender
#   }
# }
#
# Notes:
# - The '-s' option in curl prevents unwanted output (like download progress).
# - The correct endpoint to fetch from is:
#     https://learn.01founders.co/api/graphql-engine/v1/graphql
# - Ensure you use proper JSON quoting and escaping for GraphQL syntax.
# - The final output should print the three fields, each on a new line,
#   exactly as shown in the example output below.
#
# Example output:
# $ bash to-git-or-not-to-git.sh
# Chameleon
# 28
# Male
# $
#
# Solution:
curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql \
--data '{"query":"{superhero(where:{id:{_eq:170}}){name power gender}}"}' \
| jq -r '.data.superhero[] | .name, .power, .gender'
