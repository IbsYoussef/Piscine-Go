#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing point..."
echo ""

{
# Build student program
cd student/point
go build

student_output=$(./point)

rm -f point
cd ../..

# Build solution program
cd solutions/point
go build

solution_output=$(./point)

rm -f point
cd ../..

} 2>/dev/null

if [ "$student_output" = "$solution_output" ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    echo ""
    echo "Output: $student_output"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected: $solution_output"
    echo "Got: $student_output"
    exit 1
fi