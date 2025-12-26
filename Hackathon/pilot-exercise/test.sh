#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

trap 'rm -f student/pilot/pilot solutions/pilot/pilot' EXIT

echo "Testing pilot..."
echo ""

{
# Build programs
cd solutions/pilot
go build 2>/dev/null
cd ../..

cd student/pilot
go build 2>/dev/null
cd ../..

student_output=$(student/pilot/pilot)
solution_output=$(solutions/pilot/pilot)

} 2>/dev/null

if [ "$student_output" = "$solution_output" ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    echo ""
    echo "Output:"
    echo "  $student_output"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "  $solution_output"
    echo ""
    echo "Got:"
    echo "  $student_output"
    exit 1
fi