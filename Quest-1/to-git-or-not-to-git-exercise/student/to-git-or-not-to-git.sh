#!/bin/sh

# TODO: Write your code here
# Fetch superhero data for id 170 from JSON API
# Display: name, power, gender (each on separate line)
#
# API: https://learn.01founders.co/assets/superhero/all.json
#
# Hints:
# - Use curl -s to fetch the JSON data
# - Use jq with select() to filter for id 170
# - Extract nested fields: .name, .powerstats.power, .appearance.gender
# - Use jq -r for raw output (no quotes)