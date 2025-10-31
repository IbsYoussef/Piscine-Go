#!/bin/sh

# TODO: Write your code here
# Fetch superhero data for id 170 from JSON API
# Display: name, power, gender (each on separate line)
#
# API: https://learn.01founders.co/assets/superhero/all.json
#
# Hints:
# - Use curl -s to fetch the JSON data silently
# - The API returns a JSON array of superhero objects
# - Use jq to parse and filter the JSON
# - Filter for the superhero with id 170 using select(.id == 170)
# - Extract these nested fields in order:
#   1. .name (from root)
#   2. .powerstats.power (nested in powerstats object)
#   3. .appearance.gender (nested in appearance object)
# - Use jq -r for raw output (removes quotes)
# - Separate multiple field extractions with commas