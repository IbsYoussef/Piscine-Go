#!/usr/bin/env bash
# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/who-are-you.sh)
expected=$(bash solutions/who-are-you.sh)

if diff <(echo "$submitted") <(echo "$expected") > /dev/null 2>&1; then
    echo "✓ Test passed!"
    exit 0
else
    echo "✗ Test failed"
    echo ""
    echo "Difference:"
    diff <(echo "$submitted") <(echo "$expected") || true
    exit 1
fi