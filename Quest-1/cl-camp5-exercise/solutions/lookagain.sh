#!/bin/sh

# Commands: find, basename, sort
# Searches for .sh files and displays names without extension, sorted in reverse
#
# Breakdown:
# find - searches recursively from current directory
# -name '*.sh' - matches all files ending with .sh
# -exec basename {} .sh ";" - for each file found:
#   basename {} .sh removes the path and .sh extension
#   {} represents the current file
#   ";" marks end of -exec command
# | sort -r - sorts results in reverse (descending) alphabetical order

find -name '*.sh' -exec basename {} .sh ";" | sort -r