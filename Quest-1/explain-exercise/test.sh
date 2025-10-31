#!/usr/bin/env bash
# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

submitted=$(bash student/explain.sh)
expected=$(bash solutions/explain.sh)

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