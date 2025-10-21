#!/bin/sh
# Exercise: cl-camp8
#
# Optional:
# Create a file named 'skip.sh' that prints the result of 'ls -l'
# while skipping every other line, starting with the first one.
#
# Explanation:
# - 'ls -l': lists directory contents in long format, one entry per line.
# - '|': pipes the output to another command.
# - 'sed '1~2d'': uses sed to delete every second line starting from line 1.
#     - '1~2' means "start at line 1, then every 2nd line after that".
#     - 'd' deletes those matching lines.
#   As a result, only the even-numbered lines of 'ls -l' are displayed.
#
# Notes:
# - This approach uses 'sed' for simplicity, but 'awk' could achieve
#   the same effect (e.g. 'awk "NR % 2 == 0"').
# - Ensure there’s no extra whitespace or unintended blank lines.
#
# Example output:
# $ ./skip.sh
# total 8
# -rw-r--r--  1 user  staff   234 Jan 21 10:00 file2.txt
# -rw-r--r--  1 user  staff   512 Jan 21 10:02 file4.txt
#
# Solution:
ls -l | sed '1~2d'
