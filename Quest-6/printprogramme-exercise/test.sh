#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing printprogramname..."
echo ""

{
# Build and run student program
cd student/printprogramname
go build
student_output=$(./printprogramname)
rm -f printprogramname
cd ../..

# Build and run solution program
cd solutions/printprogramname
go build
solution_output=$(./printprogramname)
rm -f printprogramname
cd ../..

} 2>/dev/null

if [ "$student_output" = "$solution_output" ]; then
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