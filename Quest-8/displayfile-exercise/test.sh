#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing displayfile..."
echo ""

# Create test file
echo "Almost there!!" > quest8.txt

{
# Build student program
cd student/displayfile
go build

student_test1=$(./displayfile 2>&1)
student_test2=$(./displayfile quest8.txt main.go 2>&1)
student_test3=$(./displayfile ../../quest8.txt 2>&1)

rm -f displayfile
cd ../..

# Build solution program
cd solutions/displayfile
go build

solution_test1=$(./displayfile 2>&1)
solution_test2=$(./displayfile quest8.txt main.go 2>&1)
solution_test3=$(./displayfile ../../quest8.txt 2>&1)

rm -f displayfile
cd ../..

} 2>/dev/null

# Clean up test file
rm -f quest8.txt

if [ "$student_test1" = "$solution_test1" ] && \
   [ "$student_test2" = "$solution_test2" ] && \
   [ "$student_test3" = "$solution_test3" ]; then
    echo -e "${GREEN}✓ All tests passed!${NC}"
    echo ""
    echo "Test 1 (no args): $student_test1"
    echo "Test 2 (too many): $student_test2"
    echo "Test 3 (file content): $student_test3"
    exit 0
else
    echo -e "${RED}✗ Tests failed${NC}"
    echo ""
    if [ "$student_test1" != "$solution_test1" ]; then
        echo "Test 1 - Expected: '$solution_test1', Got: '$student_test1'"
    fi
    if [ "$student_test2" != "$solution_test2" ]; then
        echo "Test 2 - Expected: '$solution_test2', Got: '$student_test2'"
    fi
    if [ "$student_test3" != "$solution_test3" ]; then
        echo "Test 3 - Expected: '$solution_test3', Got: '$student_test3'"
    fi
    exit 1
fi