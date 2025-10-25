#!/bin/sh

# TODO: Write your code here
# Run ls -l and skip every other line, starting with the first
# Output only even-numbered lines (2nd, 4th, 6th, etc.)
#
# Hints:
# - Use ls -l for long format listing
# - Use sed with pattern '1~2d' to delete odd lines
# - Pipe the commands together