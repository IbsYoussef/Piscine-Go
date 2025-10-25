#!/bin/sh

# Commands: ls, sed
# Lists files in long format, skipping every other line starting with the first
#
# Breakdown:
# ls -l - lists directory contents in long format (one entry per line)
# | - pipes output to next command
# sed '1~2d' - deletes every second line starting from line 1
#   1~2 means: start at line 1, then every 2nd line
#   d means: delete those lines
#
# Result: Only even-numbered lines (2, 4, 6, etc.) are displayed
#
# Alternative using awk:
# ls -l | awk 'NR % 2 == 0'

ls -l | sed '1~2d'