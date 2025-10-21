#!/bin/sh
# Exercise: cl-camp6
#
# Optional: 
# Create a file named 'countfiles.sh' that prints the number (and only the number)
# of regular files and directories contained in the current directory and
# all its subdirectories. The current directory itself must be included in the count.
#
# Explanation:
# - 'find .': starts the search from the current directory ('.').
# - '\( -type f -or -type d \)': matches both files (-type f) and directories (-type d).
# - '\( -not -path "*/." \)': ensures that special directory entries like '.' are not double-counted.
# - '| wc -l': counts the total number of matching lines (each representing a file or folder),
#   and outputs only the count.
#
# Notes:
# - The parentheses are escaped with backslashes to ensure they’re interpreted correctly by the shell.
# - The result is a single number followed by a newline, as required.
# - The 'cat -e' test confirms that the output ends cleanly with a single '$'.
#
# Example output:
# $ ./countfiles.sh | cat -e
# 12$
# $
#
# Solution:
find . \( -type f -or -type d \) -and \( -not -path "*/." \) | wc -l
