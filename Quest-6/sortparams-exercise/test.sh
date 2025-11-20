#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing sortparams..."
echo ""

{
# Build and run student program
cd student/sortparams
go build
student_output=$(./sortparams 1 a 2 A 3 b 4 C)
rm -f sortparams
cd ../..

# Build and run solution program
cd solutions/sortparams
go build
solution_output=$(./sortparams 1 a 2 A 3 b 4 C)
rm -f sortparams
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