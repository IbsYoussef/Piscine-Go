#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

trap 'rm -f student/fixthemain/fixthemain solutions/fixthemain/fixthemain' EXIT

echo "Testing fixthemain..."
echo ""

{
# Build programs
cd solutions/fixthemain
go build 2>/dev/null
cd ../..

cd student/fixthemain
go build 2>/dev/null
cd ../..

student_output=$(student/fixthemain/fixthemain)
solution_output=$(solutions/fixthemain/fixthemain)

} 2>/dev/null

if [ "$student_output" = "$solution_output" ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    echo ""
    echo "Output:"
    echo "$student_output" | sed 's/^/  /'
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "$solution_output" | sed 's/^/  /'
    echo ""
    echo "Got:"
    echo "$student_output" | sed 's/^/  /'
    exit 1
fi