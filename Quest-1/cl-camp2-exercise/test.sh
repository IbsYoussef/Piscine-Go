#!/usr/bin/env bash
# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

# Filter out comments and empty lines from both files
student_content=$(grep -v '^\s*#' student/r | grep -v '^\s*$')
solution_content=$(grep -v '^\s*#' solutions/r | grep -v '^\s*$')

if [ "$student_content" = "$solution_content" ]; then
	echo "✓ Test passed!"
	exit 0
else
	echo "✗ Test failed"
	echo ""
	echo "Your file content:"
	echo "$student_content" | cat -e
	echo ""
	echo "Expected content:"
	echo "$solution_content" | cat -e
	exit 1
fi