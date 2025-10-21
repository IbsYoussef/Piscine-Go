#!/bin/sh
# Exercise: who-are-you
#
# Required:
# You just woke up in a dark alley with the tag "subject Id: 70".
# Your task is to fetch the superhero list and print the name corresponding
# to the entry whose id is 70.
#
# This script will:
# - Download the JSON list from the provided URL using curl.
# - Use jq to select the object with id == 70 and print its "name" field.
# - The output will be a JSON string (including quotes) to match the required
#   format when piped to `cat -e` (e.g. "name"$).
#
# Command used:
curl -s 'https://learn.01founders.co/assets/superhero/all.json' | jq '.[] | select(.id==70) | .name'
#
# Explanation:
# - curl -s:
#     Downloads the content from the URL silently (no progress bar or error messages).
#     The '-s' flag ensures that only the content is printed to standard output.
# - jq '.[] | select(.id==70) | .name':
#     Parses the JSON array returned by curl.
#     '.[]' iterates through all objects in the array.
#     'select(.id==70)' filters for the object whose id field equals 70.
#     '.name' extracts only the name field of that object.
#     The output includes quotes (") as it is a JSON string, which is required
#     for the correct output formatting.
#
# Example usage:
# $ ./who-are-you.sh | cat -e
# "name"$
#
# Notes:
# - Make sure 'jq' is installed on your system before running this script.
# - The expected output includes the quotes to match the grader’s requirements.
# - This exercise introduces basic JSON parsing and data extraction using curl and jq.
