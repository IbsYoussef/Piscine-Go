#!/usr/bin/env bash
# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

# Run tests in the test-data directory
submitted=$(cd test-data && grep -v '^\s*#' ../student/look | grep -v '^\s*$' | bash | sort)
expected=$(cd test-data && grep -v '^\s*#' ../solutions/look | grep -v '^\s*$' | bash | sort)

if [ "$submitted" = "$expected" ]; then
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