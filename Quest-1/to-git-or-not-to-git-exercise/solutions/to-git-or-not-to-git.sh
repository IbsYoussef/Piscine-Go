#!/bin/sh

# Commands: curl, jq
# Fetches superhero data from GraphQL API and extracts specific fields
#
# Breakdown:
# curl -s - sends silent HTTP request (no progress/error messages)
# https://... - GraphQL API endpoint
# --data '...' - sends GraphQL query as POST data
# 
# GraphQL query structure:
# {superhero(where:{id:{_eq:170}}){name power gender}}
#   - superhero: the table/type to query
#   - where:{id:{_eq:170}}: filter condition (id equals 170)
#   - {name power gender}: fields to retrieve
#
# | jq -r - pipes JSON response to jq for parsing
#   -r: raw output (no quotes)
#   '.data.superhero[]': navigates to superhero array in response
#   '| .name, .power, .gender': extracts these three fields
#   Each field is printed on a separate line

curl -s https://learn.01founders.co/assets/superhero/all.json \
| jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'