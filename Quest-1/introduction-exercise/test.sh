#!/usr/bin/env bash
# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/hello.sh)
expected=$(bash solutions/hello.sh)

if diff <(echo "$submitted") <(echo "$expected") > /dev/null 2>&1; then
    echo "✓ Test passed!"
    exit 0
else
    echo "✗ Test failed"
    echo ""
    echo "Difference:"
    diff <(echo "$submitted") <(echo "$expected") | cat -t || true
    exit 1
fi