#!/bin/sh

# Commands: find, wc
# Counts total number of files and directories
#
# Breakdown:
# find . - starts search from current directory
# \( -type f -or -type d \) - matches both files AND directories
#   -type f: regular files
#   -type d: directories
#   -or: logical OR operator
# -and \( -not -path "*/." \) - excludes special '.' entries to avoid double-counting
# | wc -l - counts the number of lines (each line is one file/directory)
#
# Note: The parentheses are escaped with backslashes for proper shell interpretation
# The current directory '.' is included in the count

find . \( -type f -or -type d \) -and \( -not -path "*/." \) | wc -l