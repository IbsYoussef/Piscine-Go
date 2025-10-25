#!/usr/bin/env bash
# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

# Run tests in the test-data directory
submitted=$(cd test-data && bash ../student/countfiles.sh)
expected=$(cd test-data && bash ../solutions/countfiles.sh)

# Trim whitespace for comparison
submitted=$(echo "$submitted" | xargs)
expected=$(echo "$expected" | xargs)

if [ "$submitted" = "$expected" ]; then
	echo "✓ Test passed!"
	echo "Count: $submitted"
	exit 0
else
	echo "✗ Test failed"
	echo ""
	echo "Your output: $submitted"
	echo "Expected: $expected"
	exit 1
fi