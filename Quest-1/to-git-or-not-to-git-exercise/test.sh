#!/usr/bin/env bash
# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

FILENAME="student/to-git-or-not-to-git.sh"

# Check if file exists
if [ ! -f ${FILENAME} ]; then
    echo "File does not exist"
    exit 1
fi

# Check if file is not empty
if [ ! -s ${FILENAME} ]; then
    echo "The file exist but is empty"
    exit 1
fi

# Check for echo command (not allowed)
if grep -q echo "$FILENAME"; then
    echo "echo is not allowed in this exercise!"
    exit 1
fi

# Run the tests
submitted=$(bash $FILENAME)
expected=$(bash solutions/to-git-or-not-to-git.sh)

if diff <(echo "$submitted") <(echo "$expected") > /dev/null 2>&1; then
    echo "✓ Test passed!"
    exit 0
else
    echo "✗ Test failed"
    echo ""
    echo "Your output:"
    echo "$submitted"
    echo ""
    echo "Expected output:"
    echo "$expected"
    exit 1
fi