#!/bin/sh

# TODO: Write your code here
# Search for all .sh files recursively
# Output just the filenames (without .sh extension)
# Sort in reverse alphabetical order
#
# Hints:
# - Use find with -name pattern
# - Use basename to strip extension
# - Use sort -r for reverse order

find -name '*.sh' -exec basename {} .sh ";" | sort -r