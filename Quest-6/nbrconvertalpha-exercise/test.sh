#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing nbrconvertalpha..."
echo ""

{
# Build student program
cd student/nbrconvertalpha
go build

# Test case 1: hello
student_output1=$(./nbrconvertalpha 8 5 12 12 15)

# Test case 2: uppercase
student_output2=$(./nbrconvertalpha --upper 8 5 25)

# Test case 3: invalid
student_output3=$(./nbrconvertalpha 32 86 h)

rm -f nbrconvertalpha
cd ../..

# Build solution program
cd solutions/nbrconvertalpha
go build

# Test case 1: hello
solution_output1=$(./nbrconvertalpha 8 5 12 12 15)

# Test case 2: uppercase
solution_output2=$(./nbrconvertalpha --upper 8 5 25)

# Test case 3: invalid
solution_output3=$(./nbrconvertalpha 32 86 h)

rm -f nbrconvertalpha
cd ../..

} 2>/dev/null

if [ "$student_output1" = "$solution_output1" ] && \
   [ "$student_output2" = "$solution_output2" ] && \
   [ "$student_output3" = "$solution_output3" ]; then
    echo -e "${GREEN}✓ All tests passed!${NC}"
    echo ""
    echo "Test 1 (hello):"
    echo "$student_output1"
    echo ""
    echo "Test 2 (uppercase):"
    echo "$student_output2"
    echo ""
    echo "Test 3 (invalid):"
    echo "$student_output3"
    exit 0
else
    echo -e "${RED}✗ Tests failed${NC}"
    echo ""
    echo "Test 1 - Expected: $solution_output1"
    echo "Test 1 - Got: $student_output1"
    echo ""
    echo "Test 2 - Expected: $solution_output2"
    echo "Test 2 - Got: $student_output2"
    echo ""
    echo "Test 3 - Expected: $solution_output3"
    echo "Test 3 - Got: $student_output3"
    exit 1
fi