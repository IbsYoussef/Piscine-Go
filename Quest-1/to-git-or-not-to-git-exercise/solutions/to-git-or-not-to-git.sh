#!/bin/sh

# Commands: curl, jq
# Fetches superhero data from JSON API and extracts specific fields
#
# Breakdown:
# curl -s https://learn.01founders.co/assets/superhero/all.json
#   - curl: command-line tool for transferring data with URLs
#   - -s: silent mode (no progress bar or error messages)
#   - Returns: JSON array of all superhero objects
#
# | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
#   - jq: lightweight JSON processor
#   - -r: raw output mode (removes quotes from strings)
#   - .[]: iterate through each element in the JSON array
#   - select(.id == 170): filter to only the superhero with id 170
#   - .name: extract the name field from root object
#   - .powerstats.power: extract power field from nested powerstats object
#   - .appearance.gender: extract gender field from nested appearance object
#   - Each field is printed on a separate line
#
# Data structure (simplified):
# [
#   {
#     "id": 170,
#     "name": "Chameleon",
#     "powerstats": { "power": 28, ... },
#     "appearance": { "gender": "Male", ... },
#     ...
#   },
#   ...
# ]
#
# Output:
# Chameleon
# 28
# Male

curl -s https://learn.01founders.co/assets/superhero/all.json \
| jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'