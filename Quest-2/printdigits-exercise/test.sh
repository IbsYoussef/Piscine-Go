#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing printdigits..."
echo ""

# Run student solution
cd student/printdigits
student_output=$(go run . 2>&1)
student_exit=$?
cd ../..

# Run expected solution
cd solutions/printdigits
solution_output=$(go run . 2>&1)
solution_exit=$?
cd ../..

# Compare outputs and exit codes
if [ "$student_output" = "$solution_output" ] && [ $student_exit -eq $solution_exit ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    echo ""
    echo "Output:"
    echo "$student_output"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "$solution_output"
    echo ""
    echo "Got:"
    echo "$student_output"
    exit 1
fi
```