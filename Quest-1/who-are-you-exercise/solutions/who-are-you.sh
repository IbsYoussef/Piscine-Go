#!/bin/sh

# Command: curl + jq
# Downloads JSON data and extracts specific field
#
# curl -s: Downloads silently (no progress bar)
# jq: JSON processor for parsing and filtering
#
# Breakdown:
# '.[]' - iterate through array
# 'select(.id==70)' - filter for id equal to 70
# '.name' - extract the name field

curl -s 'https://learn.01founders.co/assets/superhero/all.json' | jq '.[] | select(.id==70) | .name'