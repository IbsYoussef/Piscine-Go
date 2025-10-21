#!/bin/sh
# Exercise: cl-camp5
#
# Optional:
# Create a file named 'lookagain.sh' that searches from the current directory
# and all its subdirectories for all files ending with '.sh'.
#
# The output must:
# - Display only the file names (without the '.sh' extension)
# - Be sorted in descending (reverse alphabetical) order
#
# Explanation:
# - 'find': searches recursively for files starting from the current directory.
# - '-name "*.sh"': matches all files that end with '.sh'.
# - '-exec basename {} .sh ";"':
#     For each file found, 'basename' removes the directory path and the '.sh' extension.
#     '{}' represents the current file found by 'find'.
#     ';' marks the end of the -exec command.
# - '| sort -r': sorts the results in reverse order (descending).
#
# Notes:
# - Ensure there’s no extra output or trailing whitespace.
# - The 'basename' command is essential to strip the '.sh' suffix cleanly.
# - When tested with 'cat -e', each filename should appear on its own line,
#   followed by a '$' to indicate the newline.
#
# Example output:
# $ ./lookagain.sh | cat -e
# file3$
# file2$
# file1$
# $
#
# Solution:
find -name '*.sh' -exec basename {} .sh ";" | sort -r
